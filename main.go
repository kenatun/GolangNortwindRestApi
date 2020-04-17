package main

import (
	"net/http"

	"github.com/GolangNortwindRestApi/database"
	"github.com/GolangNortwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var productRepository = product.NewRepository(databaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)

}
