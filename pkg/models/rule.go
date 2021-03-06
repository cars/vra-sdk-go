// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rule A rule used in a security group.
//
// swagger:model Rule
type Rule struct {

	// Type of access (Allow, Deny or Drop) for the security rule. Allow is default. Traffic that does not match any rules will be denied.
	// Example: Allow
	// Required: true
	// Enum: [Allow Deny Drop]
	Access *string `json:"access"`

	// Direction of the security rule (inbound or outboud).
	// Example: Outbound
	// Required: true
	// Enum: [Inbound Outbound]
	Direction *string `json:"direction"`

	// IP address(es) in CIDR format which the security rule applies to.
	// Example: 66.170.99.2/32
	// Required: true
	IPRangeCidr *string `json:"ipRangeCidr"`

	// Name of security rule.
	// Example: 5756f7e2
	Name string `json:"name,omitempty"`

	// Ports the security rule applies to.
	// Example: 443, 1-655535
	// Required: true
	Ports *string `json:"ports"`

	// Protocol the security rule applies to.
	// Example: ANY, TCP, UDP
	Protocol string `json:"protocol,omitempty"`

	// Service defined by the provider (such as: SSH, HTTPS). Either service or protocol have to be specified.
	// Example: HTTPS, SSH
	Service string `json:"service,omitempty"`
}

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPRangeCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ruleTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Allow","Deny","Drop"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleTypeAccessPropEnum = append(ruleTypeAccessPropEnum, v)
	}
}

const (

	// RuleAccessAllow captures enum value "Allow"
	RuleAccessAllow string = "Allow"

	// RuleAccessDeny captures enum value "Deny"
	RuleAccessDeny string = "Deny"

	// RuleAccessDrop captures enum value "Drop"
	RuleAccessDrop string = "Drop"
)

// prop value enum
func (m *Rule) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ruleTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Rule) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

var ruleTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleTypeDirectionPropEnum = append(ruleTypeDirectionPropEnum, v)
	}
}

const (

	// RuleDirectionInbound captures enum value "Inbound"
	RuleDirectionInbound string = "Inbound"

	// RuleDirectionOutbound captures enum value "Outbound"
	RuleDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *Rule) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ruleTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Rule) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateIPRangeCidr(formats strfmt.Registry) error {

	if err := validate.Required("ipRangeCidr", "body", m.IPRangeCidr); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rule based on context it is used
func (m *Rule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule) UnmarshalBinary(b []byte) error {
	var res Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
