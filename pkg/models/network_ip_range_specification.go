// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkIPRangeSpecification Specification for creating or updating a NetworkIPRange
//
// swagger:model NetworkIPRangeSpecification
type NetworkIPRangeSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// End IP address of the range.
	// Required: true
	EndIPAddress *string `json:"endIPAddress"`

	// Deprecated. Use 'fabricNetworkIds'.
	FabricNetworkID string `json:"fabricNetworkId,omitempty"`

	// A list of fabric network Ids that this IP range should be associated with.
	FabricNetworkIds []string `json:"fabricNetworkIds"`

	// IP address version: IPv4 or IPv6. Default: IPv4.
	// Enum: [IPv4 IPv6]
	IPVersion string `json:"ipVersion,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Start IP address of the range.
	// Required: true
	StartIPAddress *string `json:"startIPAddress"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this network IP range specification
func (m *NetworkIPRangeSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkIPRangeSpecification) validateEndIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("endIPAddress", "body", m.EndIPAddress); err != nil {
		return err
	}

	return nil
}

var networkIpRangeSpecificationTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkIpRangeSpecificationTypeIPVersionPropEnum = append(networkIpRangeSpecificationTypeIPVersionPropEnum, v)
	}
}

const (

	// NetworkIPRangeSpecificationIPVersionIPV4 captures enum value "IPv4"
	NetworkIPRangeSpecificationIPVersionIPV4 string = "IPv4"

	// NetworkIPRangeSpecificationIPVersionIPV6 captures enum value "IPv6"
	NetworkIPRangeSpecificationIPVersionIPV6 string = "IPv6"
)

// prop value enum
func (m *NetworkIPRangeSpecification) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkIpRangeSpecificationTypeIPVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkIPRangeSpecification) validateIPVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPVersionEnum("ipVersion", "body", m.IPVersion); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPRangeSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPRangeSpecification) validateStartIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("startIPAddress", "body", m.StartIPAddress); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPRangeSpecification) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network IP range specification based on the context it is used
func (m *NetworkIPRangeSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkIPRangeSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIPRangeSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIPRangeSpecification) UnmarshalBinary(b []byte) error {
	var res NetworkIPRangeSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
