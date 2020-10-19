// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Deployment Deployment
//
// A group of resources such as machines, network, software, etc... typically provisioned together to deliver a complete/workable application.
//
// swagger:model Deployment
type Deployment struct {

	// Expanded deployment blueprint
	Blueprint *ResourceReference `json:"blueprint,omitempty"`

	// Deployment blueprint id
	BlueprintID string `json:"blueprintId,omitempty"`

	// Deployment blueprint version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Expanded deployment catalog
	Catalog *ResourceReference `json:"catalog,omitempty"`

	// Deployment catalog item id
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Deployment catalog version
	CatalogItemVersion string `json:"catalogItemVersion,omitempty"`

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	CreatedBy string `json:"createdBy,omitempty"`

	// Indicates whether the deployment is deleted or not.
	Deleted bool `json:"deleted,omitempty"`

	// Description of the deployment
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	Expense *Expense `json:"expense,omitempty"`

	// Deployment icon id
	IconID string `json:"iconId,omitempty"`

	// Id of the deployment
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The inputs that were used to request this deployment
	Inputs interface{} `json:"inputs,omitempty"`

	// Last request
	LastRequest *Request `json:"lastRequest,omitempty"`

	// Update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Lease expiration time
	// Format: date-time
	LeaseExpireAt strfmt.DateTime `json:"leaseExpireAt,omitempty"`

	// Name of the deployment
	// Required: true
	Name *string `json:"name"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// Owned by
	OwnedBy string `json:"ownedBy,omitempty"`

	// Expanded deployment project
	Project *ResourceReference `json:"project,omitempty"`

	// Deployment project id
	ProjectID string `json:"projectId,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible
	Resources []*Resource `json:"resources"`

	// Deployment status.
	// Enum: [CREATE_SUCCESSFUL CREATE_INPROGRESS CREATE_FAILED UPDATE_SUCCESSFUL UPDATE_INPROGRESS UPDATE_FAILED DELETE_SUCCESSFUL DELETE_INPROGRESS DELETE_FAILED ACTION_SUCCESSFUL ACTION_INPROGRESS ACTION_FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this deployment
func (m *Deployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaseExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deployment) validateBlueprint(formats strfmt.Registry) error {

	if swag.IsZero(m.Blueprint) { // not required
		return nil
	}

	if m.Blueprint != nil {
		if err := m.Blueprint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blueprint")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalog")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateExpense(formats strfmt.Registry) error {

	if swag.IsZero(m.Expense) { // not required
		return nil
	}

	if m.Expense != nil {
		if err := m.Expense.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expense")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateLastRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.LastRequest) { // not required
		return nil
	}

	if m.LastRequest != nil {
		if err := m.LastRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastRequest")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateLastUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedAt", "body", "date-time", m.LastUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateLeaseExpireAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LeaseExpireAt) { // not required
		return nil
	}

	if err := validate.FormatOf("leaseExpireAt", "body", "date-time", m.LeaseExpireAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var deploymentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE_SUCCESSFUL","CREATE_INPROGRESS","CREATE_FAILED","UPDATE_SUCCESSFUL","UPDATE_INPROGRESS","UPDATE_FAILED","DELETE_SUCCESSFUL","DELETE_INPROGRESS","DELETE_FAILED","ACTION_SUCCESSFUL","ACTION_INPROGRESS","ACTION_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentTypeStatusPropEnum = append(deploymentTypeStatusPropEnum, v)
	}
}

const (

	// DeploymentStatusCREATESUCCESSFUL captures enum value "CREATE_SUCCESSFUL"
	DeploymentStatusCREATESUCCESSFUL string = "CREATE_SUCCESSFUL"

	// DeploymentStatusCREATEINPROGRESS captures enum value "CREATE_INPROGRESS"
	DeploymentStatusCREATEINPROGRESS string = "CREATE_INPROGRESS"

	// DeploymentStatusCREATEFAILED captures enum value "CREATE_FAILED"
	DeploymentStatusCREATEFAILED string = "CREATE_FAILED"

	// DeploymentStatusUPDATESUCCESSFUL captures enum value "UPDATE_SUCCESSFUL"
	DeploymentStatusUPDATESUCCESSFUL string = "UPDATE_SUCCESSFUL"

	// DeploymentStatusUPDATEINPROGRESS captures enum value "UPDATE_INPROGRESS"
	DeploymentStatusUPDATEINPROGRESS string = "UPDATE_INPROGRESS"

	// DeploymentStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	DeploymentStatusUPDATEFAILED string = "UPDATE_FAILED"

	// DeploymentStatusDELETESUCCESSFUL captures enum value "DELETE_SUCCESSFUL"
	DeploymentStatusDELETESUCCESSFUL string = "DELETE_SUCCESSFUL"

	// DeploymentStatusDELETEINPROGRESS captures enum value "DELETE_INPROGRESS"
	DeploymentStatusDELETEINPROGRESS string = "DELETE_INPROGRESS"

	// DeploymentStatusDELETEFAILED captures enum value "DELETE_FAILED"
	DeploymentStatusDELETEFAILED string = "DELETE_FAILED"

	// DeploymentStatusACTIONSUCCESSFUL captures enum value "ACTION_SUCCESSFUL"
	DeploymentStatusACTIONSUCCESSFUL string = "ACTION_SUCCESSFUL"

	// DeploymentStatusACTIONINPROGRESS captures enum value "ACTION_INPROGRESS"
	DeploymentStatusACTIONINPROGRESS string = "ACTION_INPROGRESS"

	// DeploymentStatusACTIONFAILED captures enum value "ACTION_FAILED"
	DeploymentStatusACTIONFAILED string = "ACTION_FAILED"
)

// prop value enum
func (m *Deployment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Deployment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deployment) UnmarshalBinary(b []byte) error {
	var res Deployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
