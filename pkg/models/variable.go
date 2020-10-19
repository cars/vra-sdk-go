// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Variable Variable
//
// swagger:model Variable
type Variable struct {

	// The variable's default value in the Terraform configuration.
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// The variable's description in the Terraform configuration.
	Description string `json:"description,omitempty"`

	// The variable's name in the Terraform configuration.
	Name string `json:"name,omitempty"`

	// Whether the variable should be obscured because of security concerns.
	Sensitive bool `json:"sensitive,omitempty"`

	// The variable's type in the Terraform configuration. Complex Terraform types may be treated as Strings.
	// Enum: [STRING NUMBER BOOL LIST MAP]
	Type string `json:"type,omitempty"`
}

// Validate validates this variable
func (m *Variable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var variableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","NUMBER","BOOL","LIST","MAP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		variableTypeTypePropEnum = append(variableTypeTypePropEnum, v)
	}
}

const (

	// VariableTypeSTRING captures enum value "STRING"
	VariableTypeSTRING string = "STRING"

	// VariableTypeNUMBER captures enum value "NUMBER"
	VariableTypeNUMBER string = "NUMBER"

	// VariableTypeBOOL captures enum value "BOOL"
	VariableTypeBOOL string = "BOOL"

	// VariableTypeLIST captures enum value "LIST"
	VariableTypeLIST string = "LIST"

	// VariableTypeMAP captures enum value "MAP"
	VariableTypeMAP string = "MAP"
)

// prop value enum
func (m *Variable) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, variableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Variable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Variable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Variable) UnmarshalBinary(b []byte) error {
	var res Variable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
