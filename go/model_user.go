/*
 * ExCommerce
 *
 * ExCommerce is an example commerce system.
 *
 * API version: beta
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// User - A user of the shop.
type User struct {

	// The UUID of the user.
	ID string `json:"id"`

	// The unique name of the user.
	Name string `json:"name"`

	// The plain text password of the user.
	Password string `json:"password,omitempty"`
}
