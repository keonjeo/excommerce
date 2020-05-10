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
	"context"
	"log"
	"net/http"

	"github.com/Teelevision/excommerce/authentication"
	"github.com/Teelevision/excommerce/controller"
	openapi "github.com/Teelevision/excommerce/go"
	"github.com/Teelevision/excommerce/model"
	"github.com/Teelevision/excommerce/persistence"
	"github.com/Teelevision/excommerce/persistence/inmemory"
)

func main() {
	log.Printf("Server started")

	// persistence
	repo := inmemory.NewAdapter()
	initProducts(context.Background(), repo)

	// authentication
	authenticator := authentication.Authenticator{UserRepository: repo}

	// controllers
	userController := controller.User{UserRepository: repo}
	productController := controller.Product{ProductRepository: repo}
	cartController := controller.Cart{CartRepository: repo, ProductRepository: repo}

	// apis
	cartsAPI := &openapi.CartsAPI{
		Authenticator:     &authenticator,
		CartController:    &cartController,
		ProductController: &productController,
	}
	ordersAPI := &openapi.OrdersAPI{}
	productsAPI := &openapi.ProductsAPI{
		ProductController: &productController,
	}
	usersAPI := &openapi.UsersAPI{
		UserController: &userController,
	}

	router := openapi.NewRouter(cartsAPI, ordersAPI, productsAPI, usersAPI)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initProducts(ctx context.Context, r persistence.ProductRepository) {
	for _, product := range []model.Product{
		{ID: "a6da78f8-2be6-49ff-b40a-32aa86a6a986", Name: "Apple", Price: 49},
		{ID: "b16088e1-9603-4676-a8df-130823cf15a5", Name: "Banana", Price: 99},
		{ID: "5438bfe8-6bd2-4a88-ac36-ec29716eb6d7", Name: "Pear", Price: 109},
		{ID: "cfae533e-d9f2-4bbc-8fcb-24866fdca8fc", Name: "Orange", Price: 79},
	} {
		err := r.CreateProduct(context.Background(), product.ID, product.Name, product.Price)
		if err != nil {
			panic(err)
		}
	}
}
