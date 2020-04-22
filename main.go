package main

import (
	"net/http"

	"github.com/GolangNortwindRestApi/customer"
	"github.com/GolangNortwindRestApi/database"
	_ "github.com/GolangNortwindRestApi/docs"
	"github.com/GolangNortwindRestApi/employee"
	"github.com/GolangNortwindRestApi/order"
	"github.com/GolangNortwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Northwind API
// @version 1.0
// @description This is a sample server called server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository  = product.NewRepository(databaseConnection)
		employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
		orderRepository    = order.NewRepository(databaseConnection)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
	)
	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)

}
