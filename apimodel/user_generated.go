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

// User user
//
// swagger:model User
type User struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// lang
	Lang string `json:"lang,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rights
	Rights *UserRights `json:"rights,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateRights(formats strfmt.Registry) error {
	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	if m.Rights != nil {
		if err := m.Rights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateRights(ctx context.Context, formats strfmt.Registry) error {

	if m.Rights != nil {
		if err := m.Rights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserRights user rights
//
// swagger:model UserRights
type UserRights struct {

	// catalog access
	CatalogAccess bool `json:"catalog_access,omitempty"`

	// companies
	Companies *UserRightsOptions `json:"companies,omitempty"`

	// contacts
	Contacts *UserRightsOptions `json:"contacts,omitempty"`

	// group id
	GroupID int64 `json:"group_id,omitempty"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is admin
	IsAdmin bool `json:"is_admin,omitempty"`

	// is free
	IsFree bool `json:"is_free,omitempty"`

	// leads
	Leads *UserRightsOptions `json:"leads,omitempty"`

	// mail access
	MailAccess bool `json:"mail_access,omitempty"`

	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// tasks
	Tasks *UserRightsOptions `json:"tasks,omitempty"`
}

// Validate validates this user rights
func (m *UserRights) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRights) validateCompanies(formats strfmt.Registry) error {
	if swag.IsZero(m.Companies) { // not required
		return nil
	}

	if m.Companies != nil {
		if err := m.Companies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "companies")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) validateContacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	if m.Contacts != nil {
		if err := m.Contacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "contacts")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) validateLeads(formats strfmt.Registry) error {
	if swag.IsZero(m.Leads) { // not required
		return nil
	}

	if m.Leads != nil {
		if err := m.Leads.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "leads")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	if m.Tasks != nil {
		if err := m.Tasks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "tasks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user rights based on the context it is used
func (m *UserRights) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompanies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLeads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRights) contextValidateCompanies(ctx context.Context, formats strfmt.Registry) error {

	if m.Companies != nil {
		if err := m.Companies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "companies")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) contextValidateContacts(ctx context.Context, formats strfmt.Registry) error {

	if m.Contacts != nil {
		if err := m.Contacts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "contacts")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) contextValidateLeads(ctx context.Context, formats strfmt.Registry) error {

	if m.Leads != nil {
		if err := m.Leads.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "leads")
			}
			return err
		}
	}

	return nil
}

func (m *UserRights) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	if m.Tasks != nil {
		if err := m.Tasks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + "tasks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRights) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRights) UnmarshalBinary(b []byte) error {
	var res UserRights
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
