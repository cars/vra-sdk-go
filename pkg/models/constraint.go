// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Constraint A constraint that is conveyed to the policy engine.
//
// swagger:model Constraint
type Constraint struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	//
	// Example: ha:strong
	// Required: true
	Expression *string `json:"expression"`

	// Indicates whether this constraint should be strictly enforced or not.
	// Required: true
	Mandatory *bool `json:"mandatory"`
}

// Validate validates this constraint
func (m *Constraint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMandatory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Constraint) validateExpression(formats strfmt.Registry) error {

	if err := validate.Required("expression", "body", m.Expression); err != nil {
		return err
	}

	return nil
}

func (m *Constraint) validateMandatory(formats strfmt.Registry) error {

	if err := validate.Required("mandatory", "body", m.Mandatory); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this constraint based on context it is used
func (m *Constraint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Constraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Constraint) UnmarshalBinary(b []byte) error {
	var res Constraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
