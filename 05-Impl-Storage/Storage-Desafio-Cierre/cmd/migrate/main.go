package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationLoader{
		Db: &mysql.Config{
			User:   "root",
			Passwd: "%VU99Asp",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationLoader(cfg)

	// - run
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data is loaded successfully!")
}
