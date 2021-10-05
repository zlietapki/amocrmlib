// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserRightsOptions user rights options
//
// swagger:model UserRightsOptions
type UserRightsOptions struct {

	// add
	Add string `json:"add,omitempty"`

	// delete
	Delete string `json:"delete,omitempty"`

	// edit
	Edit string `json:"edit,omitempty"`

	// export
	Export string `json:"export,omitempty"`

	// view
	View string `json:"view,omitempty"`
}

// Validate validates this user rights options
func (m *UserRightsOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user rights options based on context it is used
func (m *UserRightsOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserRightsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRightsOptions) UnmarshalBinary(b []byte) error {
	var res UserRightsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
