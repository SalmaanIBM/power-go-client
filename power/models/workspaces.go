// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Workspaces workspaces
//
// swagger:model Workspaces
type Workspaces struct {

	// The list of available workspaces
	// Required: true
	Workspace []*Workspace `json:"workspace"`
}

// Validate validates this workspaces
func (m *Workspaces) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspaces) validateWorkspace(formats strfmt.Registry) error {

	if err := validate.Required("workspace", "body", m.Workspace); err != nil {
		return err
	}

	for i := 0; i < len(m.Workspace); i++ {
		if swag.IsZero(m.Workspace[i]) { // not required
			continue
		}

		if m.Workspace[i] != nil {
			if err := m.Workspace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workspaces based on the context it is used
func (m *Workspaces) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspaces) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workspace); i++ {

		if m.Workspace[i] != nil {

			if swag.IsZero(m.Workspace[i]) { // not required
				return nil
			}

			if err := m.Workspace[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workspaces) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspaces) UnmarshalBinary(b []byte) error {
	var res Workspaces
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
