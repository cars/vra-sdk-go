// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetTerraformConfigurationSourcesUsingGET1Reader is a Reader for the GetTerraformConfigurationSourcesUsingGET1 structure.
type GetTerraformConfigurationSourcesUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerraformConfigurationSourcesUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerraformConfigurationSourcesUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTerraformConfigurationSourcesUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTerraformConfigurationSourcesUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTerraformConfigurationSourcesUsingGET1OK creates a GetTerraformConfigurationSourcesUsingGET1OK with default headers values
func NewGetTerraformConfigurationSourcesUsingGET1OK() *GetTerraformConfigurationSourcesUsingGET1OK {
	return &GetTerraformConfigurationSourcesUsingGET1OK{}
}

/* GetTerraformConfigurationSourcesUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetTerraformConfigurationSourcesUsingGET1OK struct {
	Payload *models.PageOfBlueprintContentSource
}

func (o *GetTerraformConfigurationSourcesUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetTerraformConfigurationSourcesUsingGET1OK) GetPayload() *models.PageOfBlueprintContentSource {
	return o.Payload
}

func (o *GetTerraformConfigurationSourcesUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfBlueprintContentSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformConfigurationSourcesUsingGET1Unauthorized creates a GetTerraformConfigurationSourcesUsingGET1Unauthorized with default headers values
func NewGetTerraformConfigurationSourcesUsingGET1Unauthorized() *GetTerraformConfigurationSourcesUsingGET1Unauthorized {
	return &GetTerraformConfigurationSourcesUsingGET1Unauthorized{}
}

/* GetTerraformConfigurationSourcesUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTerraformConfigurationSourcesUsingGET1Unauthorized struct {
}

func (o *GetTerraformConfigurationSourcesUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGET1Unauthorized ", 401)
}

func (o *GetTerraformConfigurationSourcesUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTerraformConfigurationSourcesUsingGET1Forbidden creates a GetTerraformConfigurationSourcesUsingGET1Forbidden with default headers values
func NewGetTerraformConfigurationSourcesUsingGET1Forbidden() *GetTerraformConfigurationSourcesUsingGET1Forbidden {
	return &GetTerraformConfigurationSourcesUsingGET1Forbidden{}
}

/* GetTerraformConfigurationSourcesUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTerraformConfigurationSourcesUsingGET1Forbidden struct {
}

func (o *GetTerraformConfigurationSourcesUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGET1Forbidden ", 403)
}

func (o *GetTerraformConfigurationSourcesUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
