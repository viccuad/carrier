// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilterResultSetByProviderDeprecatedResponse filter result set by provider deprecated response
//
// swagger:model filterResultSetByProviderDeprecatedResponse
type FilterResultSetByProviderDeprecatedResponse struct {

	// The label
	Label string `json:"label,omitempty"`

	// The provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this filter result set by provider deprecated response
func (m *FilterResultSetByProviderDeprecatedResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilterResultSetByProviderDeprecatedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterResultSetByProviderDeprecatedResponse) UnmarshalBinary(b []byte) error {
	var res FilterResultSetByProviderDeprecatedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}