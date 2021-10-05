// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmbLead emb lead
//
// swagger:model EmbLead
type EmbLead struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this emb lead
func (m *EmbLead) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this emb lead based on context it is used
func (m *EmbLead) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmbLead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbLead) UnmarshalBinary(b []byte) error {
	var res EmbLead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}