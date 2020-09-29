// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RateFactor RateFactor
// swagger:model RateFactor
type RateFactor struct {

	// context metering item
	ContextMeteringItem string `json:"contextMeteringItem,omitempty"`

	// rate factor
	RateFactor float64 `json:"rateFactor,omitempty"`
}

// Validate validates this rate factor
func (m *RateFactor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RateFactor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateFactor) UnmarshalBinary(b []byte) error {
	var res RateFactor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}