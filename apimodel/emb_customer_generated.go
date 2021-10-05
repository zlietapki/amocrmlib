// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmbCustomer emb customer
//
// swagger:model EmbCustomer
type EmbCustomer struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this emb customer
func (m *EmbCustomer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this emb customer based on context it is used
func (m *EmbCustomer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmbCustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbCustomer) UnmarshalBinary(b []byte) error {
	var res EmbCustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}