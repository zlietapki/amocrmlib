// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Link link
//
// swagger:model Link
type Link struct {

	// metadata
	Metadata *LinkMetadata `json:"metadata,omitempty"`

	// to entity id
	ToEntityID int64 `json:"to_entity_id,omitempty"`

	// to entity type
	ToEntityType string `json:"to_entity_type,omitempty"`
}

// Validate validates this link
func (m *Link) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Link) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this link based on the context it is used
func (m *Link) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Link) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Link) UnmarshalBinary(b []byte) error {
	var res Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinkMetadata link metadata
//
// swagger:model LinkMetadata
type LinkMetadata struct {

	// catalog id
	CatalogID int64 `json:"catalog_id,omitempty"`

	// main contact
	MainContact bool `json:"main_contact,omitempty"`

	// price id
	PriceID int64 `json:"price_id,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`
}

// Validate validates this link metadata
func (m *LinkMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link metadata based on context it is used
func (m *LinkMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkMetadata) UnmarshalBinary(b []byte) error {
	var res LinkMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
