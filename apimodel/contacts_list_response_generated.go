// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactsListResponse contacts list response
//
// swagger:model ContactsListResponse
type ContactsListResponse struct {

	// embedded
	Embedded *ContactsListResponseEmbedded `json:"_embedded,omitempty"`

	// page
	Page int64 `json:"_page,omitempty"`
}

// Validate validates this contacts list response
func (m *ContactsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsListResponse) validateEmbedded(formats strfmt.Registry) error {
	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contacts list response based on the context it is used
func (m *ContactsListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmbedded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsListResponse) contextValidateEmbedded(ctx context.Context, formats strfmt.Registry) error {

	if m.Embedded != nil {
		if err := m.Embedded.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactsListResponse) UnmarshalBinary(b []byte) error {
	var res ContactsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContactsListResponseEmbedded contacts list response embedded
//
// swagger:model ContactsListResponseEmbedded
type ContactsListResponseEmbedded struct {

	// contacts
	Contacts []*Contact `json:"contacts"`
}

// Validate validates this contacts list response embedded
func (m *ContactsListResponseEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContacts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsListResponseEmbedded) validateContacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contacts); i++ {
		if swag.IsZero(m.Contacts[i]) { // not required
			continue
		}

		if m.Contacts[i] != nil {
			if err := m.Contacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "contacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contacts list response embedded based on the context it is used
func (m *ContactsListResponseEmbedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsListResponseEmbedded) contextValidateContacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Contacts); i++ {

		if m.Contacts[i] != nil {
			if err := m.Contacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "contacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactsListResponseEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactsListResponseEmbedded) UnmarshalBinary(b []byte) error {
	var res ContactsListResponseEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
