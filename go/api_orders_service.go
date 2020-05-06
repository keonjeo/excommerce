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
	"errors"
)

// OrdersApiService is a service that implents the logic for the OrdersApiServicer
// This service should implement the business logic for every endpoint for the OrdersApi API.
// Include any external packages or services that will be required by this service.
type OrdersApiService struct {
}

// NewOrdersApiService creates a default api service
func NewOrdersApiService() OrdersApiServicer {
	return &OrdersApiService{}
}

// CreateOrderFromCart - Create order from cart
func (s *OrdersApiService) CreateOrderFromCart(cartId string, order Order) (interface{}, error) {
	// TODO - update CreateOrderFromCart with the required logic for this service method.
	// Add api_orders_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'CreateOrderFromCart' not implemented")
}

// PlaceOrder - Place order
func (s *OrdersApiService) PlaceOrder(orderId string) (interface{}, error) {
	// TODO - update PlaceOrder with the required logic for this service method.
	// Add api_orders_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PlaceOrder' not implemented")
}