/*
 * ExCommerce
 *
 * ExCommerce is an example commerce system.
 *
 * API version: beta
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/Teelevision/excommerce/controller"
	openapi "github.com/Teelevision/excommerce/go"
	"github.com/Teelevision/excommerce/persistence/inmemory"
)

func main() {
	log.Printf("Server started")

	repo := inmemory.NewAdapter()

	CartsAPIService := openapi.NewCartsAPIService()
	CartsAPIController := openapi.NewCartsAPIController(CartsAPIService)

	OrdersAPIService := openapi.NewOrdersAPIService()
	OrdersAPIController := openapi.NewOrdersAPIController(OrdersAPIService)

	ProductsAPIService := openapi.NewProductsAPIService()
	ProductsAPIController := openapi.NewProductsAPIController(ProductsAPIService)

	UsersAPIController := &openapi.UsersAPIController{
		CreateUserController: &controller.CreateUser{
			UserRepository: repo,
		},
		GetUserController: &controller.GetUser{
			UserRepository: repo,
		},
	}

	router := openapi.NewRouter(CartsAPIController, OrdersAPIController, ProductsAPIController, UsersAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
