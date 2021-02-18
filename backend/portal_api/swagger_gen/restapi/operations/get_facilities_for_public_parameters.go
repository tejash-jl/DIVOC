// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetFacilitiesForPublicParams creates a new GetFacilitiesForPublicParams object
//
// There are no default values defined in the spec.
func NewGetFacilitiesForPublicParams() GetFacilitiesForPublicParams {

	return GetFacilitiesForPublicParams{}
}

// GetFacilitiesForPublicParams contains all the bound params for the get facilities for public operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFacilitiesForPublic
type GetFacilitiesForPublicParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Limit *float64
	/*
	  In: query
	*/
	Offset *float64
	/*Facility pincode
	  In: query
	*/
	Pincode *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFacilitiesForPublicParams() beforehand.
func (o *GetFacilitiesForPublicParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qPincode, qhkPincode, _ := qs.GetOK("pincode")
	if err := o.bindPincode(qPincode, qhkPincode, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetFacilitiesForPublicParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "float64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetFacilitiesForPublicParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "float64", raw)
	}
	o.Offset = &value

	return nil
}

// bindPincode binds and validates parameter Pincode from query.
func (o *GetFacilitiesForPublicParams) bindPincode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Pincode = &raw

	return nil
}
