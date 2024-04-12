package application

import (
	"app/internal/application/loader"
	"app/internal/repository"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationLoader struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Addr is the server address.
	Addr string
}

type ApplicationLoader struct {
	// cfgDb is the database configuration.
	cfgDb *mysql.Config
	// db is the database connection.
	db *sql.DB
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationLoader(config *ConfigApplicationLoader) *ApplicationLoader {
	// default values
	defaultCfg := &ConfigApplicationLoader{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationLoader{
		cfgDb: defaultCfg.Db,
	}
}

// Run runs the application.
func (a *ApplicationLoader) Run() (err error) {
	// dependencies
	// - db: init
	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(a.db)
	rpProduct := repository.NewProductsMySQL(a.db)
	rpInvoice := repository.NewInvoicesMySQL(a.db)
	rpSale := repository.NewSalesMySQL(a.db)

	customerLoader := loader.NewCustomerLoaderJSON("./docs/db/json/customers.json", rpCustomer)
	productLoader := loader.NewProductLoaderJSON("./docs/db/json/products.json", rpProduct)
	invoiceLoader := loader.NewInvoiceLoaderJSON("./docs/db/json/invoices.json", rpInvoice)
	saleLoader := loader.NewSaleLoaderJSON("./docs/db/json/sales.json", rpSale)

	err = customerLoader.Load()
	if err != nil {
		return
	}

	err = productLoader.Load()
	if err != nil {
		return
	}

	err = invoiceLoader.Load()
	if err != nil {
		return
	}

	err = saleLoader.Load()
	if err != nil {
		return
	}

	return
}
