// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkDomain The network domain object is the parent of related subnets that are configured together. Subnets within a domain must have non-overlapping address spaces.<br>**HATEOAS** links:<br>**cloud-accounts** - array[CloudAccount] - List of cloud accounts where this network domain exists.<br>**self** - NetworkDomain - Self link to this network domain
// swagger:model NetworkDomain
type NetworkDomain struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// IPv4 address range of the network domain in CIDR format
	// Required: true
	Cidr *string `json:"cidr"`

	// Set of ids of the cloud accounts this entity belongs to.
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the network domain.
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// A set of tag keys and optional values that were set on this network domain.
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this network domain
func (m *NetworkDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *NetworkDomain) validateLinks(formats strfmt.Registry) error {

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *NetworkDomain) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDomain) validateCloudAccountIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDomain) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDomain) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDomain) validateTags(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *NetworkDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDomain) UnmarshalBinary(b []byte) error {
	var res NetworkDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
