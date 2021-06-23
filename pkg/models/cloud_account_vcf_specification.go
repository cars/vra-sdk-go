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

// CloudAccountVcfSpecification Specification for a VCF cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountVcfSpecification
type CloudAccountVcfSpecification struct {

	// Accept self signed certificate when connecting to vSphere and NSX-T
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	DcID string `json:"dcId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// NSX Certificate
	NsxCertificate string `json:"nsxCertificate,omitempty"`

	// Host name for the NSX endpoint from the specified workload domain.
	// Example: nsxt.mycompany.com
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Password for the user used to authenticate with the NSX-T manager in VCF cloud account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	NsxPassword *string `json:"nsxPassword"`

	// User name for the NSX manager in the specified workload domain.
	// Example: administrator@mycompany.com
	// Required: true
	NsxUsername *string `json:"nsxUsername"`

	// A set of Region names to enable provisioning on.Refer to /iaas/cloud-accounts/region-enumeration. Deprecated - use regions to define enabled regions.
	// Example: [ \"us-east-1\", \"ap-northeast-1\" ]
	// Required: true
	RegionIds []string `json:"regionIds"`

	// A set of regions to enable provisioning on.Refer to /iaas/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"us-east-1\",\"externalRegionId\": \"us-east-1\"}]
	Regions []*RegionSpecification `json:"regions"`

	// SDDC manager integration id
	SddcManagerID string `json:"sddcManagerId,omitempty"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// vCenter Certificate
	VcenterCertificate string `json:"vcenterCertificate,omitempty"`

	// Host name for the vSphere from the specified workload domain.
	// Example: vc.mycompany.com
	// Required: true
	VcenterHostName *string `json:"vcenterHostName"`

	// Password for the user used to authenticate with the vCenter in VCF cloud account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	VcenterPassword *string `json:"vcenterPassword"`

	// vCenter user name for the specified workload domain.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	// Example: administrator@mycompany.com
	// Required: true
	VcenterUsername *string `json:"vcenterUsername"`

	// Id of the workload domain to add as VCF cloud account.
	// Required: true
	WorkloadDomainID *string `json:"workloadDomainId"`

	// Name of the workload domain to add as VCF cloud account.
	// Example: Management
	// Required: true
	WorkloadDomainName *string `json:"workloadDomainName"`
}

// Validate validates this cloud account vcf specification
func (m *CloudAccountVcfSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadDomainName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVcfSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateNsxHostName(formats strfmt.Registry) error {

	if err := validate.Required("nsxHostName", "body", m.NsxHostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateNsxPassword(formats strfmt.Registry) error {

	if err := validate.Required("nsxPassword", "body", m.NsxPassword); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateNsxUsername(formats strfmt.Registry) error {

	if err := validate.Required("nsxUsername", "body", m.NsxUsername); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateRegionIds(formats strfmt.Registry) error {

	if err := validate.Required("regionIds", "body", m.RegionIds); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateTags(formats strfmt.Registry) error {
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

func (m *CloudAccountVcfSpecification) validateVcenterHostName(formats strfmt.Registry) error {

	if err := validate.Required("vcenterHostName", "body", m.VcenterHostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateVcenterPassword(formats strfmt.Registry) error {

	if err := validate.Required("vcenterPassword", "body", m.VcenterPassword); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateVcenterUsername(formats strfmt.Registry) error {

	if err := validate.Required("vcenterUsername", "body", m.VcenterUsername); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateWorkloadDomainID(formats strfmt.Registry) error {

	if err := validate.Required("workloadDomainId", "body", m.WorkloadDomainID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcfSpecification) validateWorkloadDomainName(formats strfmt.Registry) error {

	if err := validate.Required("workloadDomainName", "body", m.WorkloadDomainName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud account vcf specification based on the context it is used
func (m *CloudAccountVcfSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVcfSpecification) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Regions); i++ {

		if m.Regions[i] != nil {
			if err := m.Regions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVcfSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountVcfSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVcfSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountVcfSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
