// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRemoveSpaceFromUserParams creates a new RemoveSpaceFromUserParams object
// no default values defined in spec.
func NewRemoveSpaceFromUserParams() RemoveSpaceFromUserParams {

	return RemoveSpaceFromUserParams{}
}

// RemoveSpaceFromUserParams contains all the bound params for the remove space from user operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeSpaceFromUser
type RemoveSpaceFromUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The guid parameter is used as a part of the request URL: '/v2/users/:guid/spaces/:space_guid'
	  Required: true
	  In: path
	*/
	GUID string
	/*The space_guid parameter is used as a part of the request URL: '/v2/users/:guid/spaces/:space_guid'
	  Required: true
	  In: path
	*/
	SpaceGUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveSpaceFromUserParams() beforehand.
func (o *RemoveSpaceFromUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSpaceGUID, rhkSpaceGUID, _ := route.Params.GetOK("space_guid")
	if err := o.bindSpaceGUID(rSpaceGUID, rhkSpaceGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveSpaceFromUserParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

// bindSpaceGUID binds and validates parameter SpaceGUID from path.
func (o *RemoveSpaceFromUserParams) bindSpaceGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SpaceGUID = raw

	return nil
}
