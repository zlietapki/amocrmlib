// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomFieldsValueItem custom fields value item
//
// swagger:model CustomFieldsValueItem
type CustomFieldsValueItem struct {

	// enum code
	EnumCode string `json:"enum_code,omitempty"`

	// enum id
	EnumID int64 `json:"enum_id,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this custom fields value item
func (m *CustomFieldsValueItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom fields value item based on context it is used
func (m *CustomFieldsValueItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomFieldsValueItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomFieldsValueItem) UnmarshalBinary(b []byte) error {
	var res CustomFieldsValueItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}