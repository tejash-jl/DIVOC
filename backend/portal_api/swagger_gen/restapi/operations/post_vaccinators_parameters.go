// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostVaccinatorsParams creates a new PostVaccinatorsParams object
// no default values defined in spec.
func NewPostVaccinatorsParams() PostVaccinatorsParams {

	return PostVaccinatorsParams{}
}

// PostVaccinatorsParams contains all the bound params for the post vaccinators operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostVaccinators
type PostVaccinatorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Facility data in the form of csv
	  In: formData
	*/
	File io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostVaccinatorsParams() beforehand.
func (o *PostVaccinatorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFile(file, fileHeader); err != nil {
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PostVaccinatorsParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
