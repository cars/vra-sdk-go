// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContentAbout ContentAbout
// swagger:model ContentAbout
type ContentAbout struct {

	// Latest API Version
	// Read Only: true
	LatestAPIVersion string `json:"latestApiVersion,omitempty"`

	// Supported APIs
	// Read Only: true
	SupportedApis []*SupportedAPI `json:"supportedApis"`
}

// Validate validates this content about
func (m *ContentAbout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportedApis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentAbout) validateSupportedApis(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedApis) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedApis); i++ {
		if swag.IsZero(m.SupportedApis[i]) { // not required
			continue
		}

		if m.SupportedApis[i] != nil {
			if err := m.SupportedApis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedApis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentAbout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentAbout) UnmarshalBinary(b []byte) error {
	var res ContentAbout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
