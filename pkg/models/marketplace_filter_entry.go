// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketplaceFilterEntry MarketplaceFilterEntry
//
// swagger:model MarketplaceFilterEntry
type MarketplaceFilterEntry struct {

	// Entry count for this filter
	// Example: 0
	Count int64 `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Filter Entry Display Name
	Name string `json:"name,omitempty"`
}

// Validate validates this marketplace filter entry
func (m *MarketplaceFilterEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this marketplace filter entry based on context it is used
func (m *MarketplaceFilterEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceFilterEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceFilterEntry) UnmarshalBinary(b []byte) error {
	var res MarketplaceFilterEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
