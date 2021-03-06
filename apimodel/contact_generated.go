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

// Contact contact
//
// swagger:model Contact
type Contact struct {

	// embedded
	Embedded *ContactEmbedded `json:"_embedded,omitempty"`

	// account id
	AccountID int64 `json:"account_id,omitempty"`

	// closest task at
	ClosestTaskAt int64 `json:"closest_task_at,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// created by
	CreatedBy int64 `json:"created_by,omitempty"`

	// custom fields values
	CustomFieldsValues []*CustomFieldsValue `json:"custom_fields_values"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// group id
	GroupID int64 `json:"group_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is deleted
	IsDeleted bool `json:"is_deleted,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// responsible user id
	ResponsibleUserID int64 `json:"responsible_user_id,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy int64 `json:"updated_by,omitempty"`
}

// Validate validates this contact
func (m *Contact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFieldsValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contact) validateEmbedded(formats strfmt.Registry) error {
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

func (m *Contact) validateCustomFieldsValues(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFieldsValues) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFieldsValues); i++ {
		if swag.IsZero(m.CustomFieldsValues[i]) { // not required
			continue
		}

		if m.CustomFieldsValues[i] != nil {
			if err := m.CustomFieldsValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contact based on the context it is used
func (m *Contact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmbedded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFieldsValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contact) contextValidateEmbedded(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Contact) contextValidateCustomFieldsValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFieldsValues); i++ {

		if m.CustomFieldsValues[i] != nil {
			if err := m.CustomFieldsValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contact) UnmarshalBinary(b []byte) error {
	var res Contact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContactEmbedded contact embedded
//
// swagger:model ContactEmbedded
type ContactEmbedded struct {

	// companies
	Companies []*EmbCompany `json:"companies"`

	// customers
	Customers []*EmbCustomer `json:"customers"`

	// leads
	Leads []*EmbLead `json:"leads"`

	// tags
	Tags []*EmbTag `json:"tags"`
}

// Validate validates this contact embedded
func (m *ContactEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactEmbedded) validateCompanies(formats strfmt.Registry) error {
	if swag.IsZero(m.Companies) { // not required
		return nil
	}

	for i := 0; i < len(m.Companies); i++ {
		if swag.IsZero(m.Companies[i]) { // not required
			continue
		}

		if m.Companies[i] != nil {
			if err := m.Companies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "companies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) validateCustomers(formats strfmt.Registry) error {
	if swag.IsZero(m.Customers) { // not required
		return nil
	}

	for i := 0; i < len(m.Customers); i++ {
		if swag.IsZero(m.Customers[i]) { // not required
			continue
		}

		if m.Customers[i] != nil {
			if err := m.Customers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "customers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) validateLeads(formats strfmt.Registry) error {
	if swag.IsZero(m.Leads) { // not required
		return nil
	}

	for i := 0; i < len(m.Leads); i++ {
		if swag.IsZero(m.Leads[i]) { // not required
			continue
		}

		if m.Leads[i] != nil {
			if err := m.Leads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "leads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contact embedded based on the context it is used
func (m *ContactEmbedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompanies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLeads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactEmbedded) contextValidateCompanies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Companies); i++ {

		if m.Companies[i] != nil {
			if err := m.Companies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "companies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) contextValidateCustomers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Customers); i++ {

		if m.Customers[i] != nil {
			if err := m.Customers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "customers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) contextValidateLeads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Leads); i++ {

		if m.Leads[i] != nil {
			if err := m.Leads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "leads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactEmbedded) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_embedded" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactEmbedded) UnmarshalBinary(b []byte) error {
	var res ContactEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
