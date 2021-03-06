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

// NetworkInterfaceSpecification Specification for attaching nic to machine
//
// swagger:model NetworkInterfaceSpecification
type NetworkInterfaceSpecification struct {

	// A list of IP addresses allocated or in use by this network interface.
	// Example: [ \"10.1.2.190\" ]
	Addresses []string `json:"addresses"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"awaitIp\" : \"true\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The device index of this network interface.
	// Example: 1
	DeviceIndex int32 `json:"deviceIndex,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Id of the network instance that this network interface plugs into.
	// Example: dcd9
	// Required: true
	NetworkID *string `json:"networkId"`

	// A list of security group ids which this network interface will be assigned to.
	SecurityGroupIds []string `json:"securityGroupIds"`
}

// Validate validates this network interface specification
func (m *NetworkInterfaceSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInterfaceSpecification) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkId", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network interface specification based on context it is used
func (m *NetworkInterfaceSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaceSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaceSpecification) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaceSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
