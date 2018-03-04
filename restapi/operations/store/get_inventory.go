// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
)

// GetInventoryHandlerFunc turns a function with the right signature into a get inventory handler
type GetInventoryHandlerFunc func(GetInventoryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInventoryHandlerFunc) Handle(params GetInventoryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetInventoryHandler interface for that can handle valid get inventory params
type GetInventoryHandler interface {
	Handle(GetInventoryParams, interface{}) middleware.Responder
}

// NewGetInventory creates a new http.Handler for the get inventory operation
func NewGetInventory(ctx *middleware.Context, handler GetInventoryHandler) *GetInventory {
	return &GetInventory{Context: ctx, Handler: handler}
}

/*GetInventory swagger:route GET /store/inventory store getInventory

Returns pet inventories by status

Returns a map of status codes to quantities

*/
type GetInventory struct {
	Context *middleware.Context
	Handler GetInventoryHandler
}

func (o *GetInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetInventoryParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetInventoryOKBody get inventory o k body
// swagger:model GetInventoryOKBody
type GetInventoryOKBody map[string]int32

// Validate validates this get inventory o k body
func (o GetInventoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
