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

// Property Property
//
// swagger:model Property
type Property struct {

	// Path that can be used to retrieve permissible values
	DollarData string `json:"$data,omitempty"`

	// Path that can be used to retrieve single permissible default value
	DollarDynamicDefault string `json:"$dynamicDefault,omitempty"`

	// Constant value for the property
	Const interface{} `json:"const,omitempty"`

	// Default value for the property
	Default interface{} `json:"default,omitempty"`

	// Property description
	Description string `json:"description,omitempty"`

	// Flag indicating if property is encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// Restrict a value to a fixed set of values for the property
	Enum []interface{} `json:"enum"`

	// The expected (and pre-defined) format to be used for string validations. Supported format values include: date-time, email, ip-address, alphanumeric, phone, etc.
	Format string `json:"format,omitempty"`

	// Property items defines a schema for the contents of an array
	Items *Property `json:"items,omitempty"`

	// Maximum number of items for the property
	MaxItems int64 `json:"maxItems,omitempty"`

	// String instance is validated against this keyword for its maximum length
	MaxLength int32 `json:"maxLength,omitempty"`

	// Maximum value for the property
	Maximum int64 `json:"maximum,omitempty"`

	// Minimum number of items for the property
	MinItems int64 `json:"minItems,omitempty"`

	// String instance is validated against this keyword for its minimum length
	MinLength int32 `json:"minLength,omitempty"`

	// Minimum value for the property
	Minimum int64 `json:"minimum,omitempty"`

	// Ensures the given data is valid against only one of the specified schemas
	OneOf []*Property `json:"oneOf"`

	// Regular expression to express constraint for the property
	Pattern string `json:"pattern,omitempty"`

	// Properties within the property
	Properties map[string]Property `json:"properties,omitempty"`

	// Flag indicating if property is read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// Property title
	Title string `json:"title,omitempty"`

	// Property type should be among valid types (number, integer, string, array, boolean, object)
	Type string `json:"type,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	if m.Items != nil {
		if err := m.Items.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

func (m *Property) validateOneOf(formats strfmt.Registry) error {
	if swag.IsZero(m.OneOf) { // not required
		return nil
	}

	for i := 0; i < len(m.OneOf); i++ {
		if swag.IsZero(m.OneOf[i]) { // not required
			continue
		}

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Property) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}
		if val, ok := m.Properties[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this property based on the context it is used
func (m *Property) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if m.Items != nil {
		if err := m.Items.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

func (m *Property) contextValidateOneOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OneOf); i++ {

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Property) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Properties {

		if val, ok := m.Properties[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}