// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/api/swagger_gen/models"
)

// GetLoggedInUserInfoHandlerFunc turns a function with the right signature into a get logged in user info handler
type GetLoggedInUserInfoHandlerFunc func(GetLoggedInUserInfoParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLoggedInUserInfoHandlerFunc) Handle(params GetLoggedInUserInfoParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetLoggedInUserInfoHandler interface for that can handle valid get logged in user info params
type GetLoggedInUserInfoHandler interface {
	Handle(GetLoggedInUserInfoParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetLoggedInUserInfo creates a new http.Handler for the get logged in user info operation
func NewGetLoggedInUserInfo(ctx *middleware.Context, handler GetLoggedInUserInfoHandler) *GetLoggedInUserInfo {
	return &GetLoggedInUserInfo{Context: ctx, Handler: handler}
}

/*GetLoggedInUserInfo swagger:route GET /v1/users/me vaccination getLoggedInUserInfo

Get User information

*/
type GetLoggedInUserInfo struct {
	Context *middleware.Context
	Handler GetLoggedInUserInfoHandler
}

func (o *GetLoggedInUserInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLoggedInUserInfoParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
