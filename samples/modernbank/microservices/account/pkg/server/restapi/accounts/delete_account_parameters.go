// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAccountParams creates a new DeleteAccountParams object
// no default values defined in spec.
func NewDeleteAccountParams() DeleteAccountParams {

	return DeleteAccountParams{}
}

// DeleteAccountParams contains all the bound params for the delete account operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteAccount
type DeleteAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The number of the account that is to be deleted.
	  Required: true
	  In: path
	*/
	Number int64
	/*Owner of the account
	  Required: true
	  In: path
	*/
	Owner string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteAccountParams() beforehand.
func (o *DeleteAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNumber, rhkNumber, _ := route.Params.GetOK("number")
	if err := o.bindNumber(rNumber, rhkNumber, route.Formats); err != nil {
		res = append(res, err)
	}

	rOwner, rhkOwner, _ := route.Params.GetOK("owner")
	if err := o.bindOwner(rOwner, rhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNumber binds and validates parameter Number from path.
func (o *DeleteAccountParams) bindNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("number", "path", "int64", raw)
	}
	o.Number = value

	return nil
}

// bindOwner binds and validates parameter Owner from path.
func (o *DeleteAccountParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Owner = raw

	return nil
}
