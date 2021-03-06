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

// RegionSpecification Specification for a region in a cloud account.
//
// swagger:model RegionSpecification
type RegionSpecification struct {

	// Unique identifier of region on the provider side.
	// Example: us-west
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// Name of region on the provider side. In vSphere, the name of the region is different from its id.
	// Example: us-west
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this region specification
func (m *RegionSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionSpecification) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *RegionSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this region specification based on context it is used
func (m *RegionSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegionSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionSpecification) UnmarshalBinary(b []byte) error {
	var res RegionSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
