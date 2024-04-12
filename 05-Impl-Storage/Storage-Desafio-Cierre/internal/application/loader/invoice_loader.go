package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// InvoiceLoader is a struct that represents an implementation of the loader interface for invoice json files
type InvoiceLoader struct {
	FilePath   string
	Repository internal.RepositoryInvoice
}

// NewInvoiceLoaderJSON is a method that creates a new invoice loader json
func NewInvoiceLoaderJSON(filePath string, repository internal.RepositoryInvoice) *InvoiceLoader {
	return &InvoiceLoader{filePath, repository}
}

// Load is a method that loads the invoices into db
func (i *InvoiceLoader) Load() (err error) {
	// Read file
	data, err := os.ReadFile(i.FilePath)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	var invoices []internal.Invoice
	if err = json.Unmarshal(data, &invoices); err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	for _, invoice := range invoices {
		err = i.Repository.Save(&invoice)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
			return
		}
	}
	return
}
