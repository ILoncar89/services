package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"services/database"
	"services/product"
	"services/receipt"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	receipt.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
