// Code generated by go-swagger; DO NOT EDIT.

package symptoms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSymptomsHandlerFunc turns a function with the right signature into a get symptoms handler
type GetSymptomsHandlerFunc func(GetSymptomsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSymptomsHandlerFunc) Handle(params GetSymptomsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSymptomsHandler interface for that can handle valid get symptoms params
type GetSymptomsHandler interface {
	Handle(GetSymptomsParams, interface{}) middleware.Responder
}

// NewGetSymptoms creates a new http.Handler for the get symptoms operation
func NewGetSymptoms(ctx *middleware.Context, handler GetSymptomsHandler) *GetSymptoms {
	return &GetSymptoms{Context: ctx, Handler: handler}
}

/*GetSymptoms swagger:route GET /symptoms symptoms getSymptoms

Get symptoms

*/
type GetSymptoms struct {
	Context *middleware.Context
	Handler GetSymptomsHandler
}

func (o *GetSymptoms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSymptomsParams()

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