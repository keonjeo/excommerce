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

// A UsersApiController binds http requests to an api service and writes the service results to the http response
type UsersApiController struct {
	service UsersApiServicer
}

// NewUsersApiController creates a default api controller
func NewUsersApiController(s UsersApiServicer) Router {
	return &UsersApiController{service: s}
}

// Routes returns all of the api route for the UsersApiController
func (c *UsersApiController) Routes() Routes {
	return Routes{
		{
			"Login",
			strings.ToUpper("Post"),
			"/beta/users/login",
			c.Login,
		},
		{
			"Register",
			strings.ToUpper("Post"),
			"/beta/users",
			c.Register,
		},
	}
}

// Login - Login a user
func (c *UsersApiController) Login(w http.ResponseWriter, r *http.Request) {
	loginForm := &LoginForm{}
	if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.Login(*loginForm)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// Register - Register a user
func (c *UsersApiController) Register(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.Register(*user)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
