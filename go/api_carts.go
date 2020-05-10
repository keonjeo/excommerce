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
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Teelevision/excommerce/authentication"
	"github.com/Teelevision/excommerce/controller"
	"github.com/Teelevision/excommerce/model"
	"github.com/gorilla/mux"
)

var _ Router = (*CartsAPI)(nil)

// A CartsAPI binds http requests to an api service and writes the service results to the http response
type CartsAPI struct {
	service CartsAPIServicer

	Authenticator     *authentication.Authenticator
	CartController    *controller.Cart
	ProductController *controller.Product
}

// Routes returns all of the api route for the CartsApiController
func (c *CartsAPI) Routes() Routes {
	return Routes{
		{
			"DeleteCart",
			strings.ToUpper("Delete"),
			"/beta/carts/{cartId}",
			c.DeleteCart,
		},
		{
			"GetAllCarts",
			strings.ToUpper("Get"),
			"/beta/carts",
			c.GetAllCarts,
		},
		{
			"GetCart",
			strings.ToUpper("Get"),
			"/beta/carts/{cartId}",
			c.GetCart,
		},
		{
			"StoreCart",
			strings.ToUpper("Put"),
			"/beta/carts/{cartId}",
			c.Authenticator.HandlerFunc(c.StoreCart),
		},
	}
}

// DeleteCart - Delete a cart
func (c *CartsAPI) DeleteCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	result, err := c.service.DeleteCart(cartID)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// GetAllCarts - Get all carts
func (c *CartsAPI) GetAllCarts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	locked := query.Get("locked") == "true"
	result, err := c.service.GetAllCarts(locked)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// GetCart - Get a cart
func (c *CartsAPI) GetCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	result, err := c.service.GetCart(cartID)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// StoreCart - Store a cart
func (c *CartsAPI) StoreCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// validation
	params := mux.Vars(r)
	cartID := params["cartId"]
	if !uuidPattern.Match([]byte(cartID)) {
		invalidInput("The cartId of the path is not a UUID.", uuidPattern.String(), w)
		return
	}
	input := &Cart{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		invalidJSON(err, w)
		return
	}
	for i, position := range input.Positions {
		if !uuidPattern.Match([]byte(position.Product.ID)) {
			failValidation("The product id is not a UUID.", fmt.Sprintf("/positions/%d/product/id", i), w)
			return
		}
	}

	// convert to internal model
	cartInput := model.Cart{
		ID:        cartID,
		Positions: make([]model.Position, len(input.Positions)),
	}
	for i, position := range input.Positions {
		cartInput.Positions[i].ProductID = position.Product.ID
		cartInput.Positions[i].Quantity = int(position.Quantity)
		// load product
		product, err := c.ProductController.Get(ctx, position.Product.ID)
		switch {
		case errors.Is(err, controller.ErrNotFound):
			failValidation("The product is not available.", fmt.Sprintf("/positions/%d/product/id", i), w)
			return
		case err == nil:
			cartInput.Positions[i].Product = product
		default:
			panic(err)
		}
	}

	// action
	var existed bool
	// create (or update if cart already exists)
	cart, err := c.CartController.CreateAndGet(ctx, &cartInput)
	if errors.Is(err, controller.ErrConflict) {
		existed = true
		cart, err = c.CartController.UpdateAndGet(ctx, &cartInput)
	}
	switch {
	case errors.Is(err, controller.ErrForbidden):
		w.WriteHeader(http.StatusForbidden) // 403
	case err == nil:
		status := http.StatusOK // 200
		if !existed {
			status = http.StatusCreated // 201
		}
		cartOutput := Cart{
			ID:        cart.ID,
			Positions: make([]Position, len(cart.Positions)),
			Locked:    cart.Locked,
		}
		for i, position := range cart.Positions {
			cartOutput.Positions[i].Product.ID = position.Product.ID
			cartOutput.Positions[i].Product.Name = position.Product.Name
			cartOutput.Positions[i].Product.Price = float32(position.Product.Price) / 100
			cartOutput.Positions[i].Quantity = int32(position.Quantity)
			cartOutput.Positions[i].Price = float32(position.Price) / 100
		}
		EncodeJSONResponse(cartOutput, &status, w)
	default:
		panic(err)
	}
}

var uuidPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
