/*
 * ExCommerce
 *
 * ExCommerce is an example commerce system.
 *
 * API version: beta
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ProductsApiController binds http requests to an api service and writes the service results to the http response
type ProductsApiController struct {
	service ProductsApiServicer
}

// NewProductsApiController creates a default api controller
func NewProductsApiController(s ProductsApiServicer) Router {
	return &ProductsApiController{service: s}
}

// Routes returns all of the api route for the ProductsApiController
func (c *ProductsApiController) Routes() Routes {
	return Routes{
		{
			"GetAllProducts",
			strings.ToUpper("Get"),
			"/beta/products",
			c.GetAllProducts,
		},
		{
			"StoreCouponForProduct",
			strings.ToUpper("Put"),
			"/beta/products/{productId}/coupon/{couponCode}",
			c.StoreCouponForProduct,
		},
	}
}

// GetAllProducts - Get all products
func (c *ProductsApiController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllProducts()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// StoreCouponForProduct - Create product coupon
func (c *ProductsApiController) StoreCouponForProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["productId"]
	couponCode := params["couponCode"]
	coupon := &Coupon{}
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.StoreCouponForProduct(productId, couponCode, *coupon)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}