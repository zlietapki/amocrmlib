// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomFieldEnum custom field enum
//
// swagger:model CustomFieldEnum
type CustomFieldEnum struct {

	// id
	ID int64 `json:"id,omitempty"`

	// sort
	Sort int64 `json:"sort,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this custom field enum
func (m *CustomFieldEnum) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom field enum based on context it is used
func (m *CustomFieldEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomFieldEnum) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomFieldEnum) UnmarshalBinary(b []byte) error {
	var res CustomFieldEnum
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
