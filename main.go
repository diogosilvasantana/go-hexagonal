package main

import (
	"database/sql"

	db2 "github.com/diogosilvasantana/go-hexagonal/adapters/db"
	"github.com/diogosilvasantana/go-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Product 1", 10)
}
