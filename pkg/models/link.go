// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Link Link
// swagger:model Link
type Link struct {

	// deprecation
	Deprecation string `json:"deprecation,omitempty" xml:"deprecation,attr"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr"`

	// hreflang
	Hreflang string `json:"hreflang,omitempty" xml:"hreflang,attr"`

	// media
	Media string `json:"media,omitempty" xml:"media,attr"`

	// rel
	Rel string `json:"rel,omitempty" xml:"rel,attr"`

	// templated
	Templated bool `json:"templated,omitempty"`

	// title
	Title string `json:"title,omitempty" xml:"title,attr"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr"`
}

// Validate validates this link
func (m *Link) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Link) UnmarshalBinary(b []byte) error {
	var res Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
