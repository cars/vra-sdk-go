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

// GetTerraformVersionUsingGET1Reader is a Reader for the GetTerraformVersionUsingGET1 structure.
type GetTerraformVersionUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerraformVersionUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerraformVersionUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTerraformVersionUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTerraformVersionUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTerraformVersionUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTerraformVersionUsingGET1OK creates a GetTerraformVersionUsingGET1OK with default headers values
func NewGetTerraformVersionUsingGET1OK() *GetTerraformVersionUsingGET1OK {
	return &GetTerraformVersionUsingGET1OK{}
}

/*GetTerraformVersionUsingGET1OK handles this case with default header values.

OK
*/
type GetTerraformVersionUsingGET1OK struct {
	Payload *models.TerraformVersion
}

func (o *GetTerraformVersionUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] getTerraformVersionUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetTerraformVersionUsingGET1OK) GetPayload() *models.TerraformVersion {
	return o.Payload
}

func (o *GetTerraformVersionUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TerraformVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformVersionUsingGET1Unauthorized creates a GetTerraformVersionUsingGET1Unauthorized with default headers values
func NewGetTerraformVersionUsingGET1Unauthorized() *GetTerraformVersionUsingGET1Unauthorized {
	return &GetTerraformVersionUsingGET1Unauthorized{}
}

/*GetTerraformVersionUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetTerraformVersionUsingGET1Unauthorized struct {
}

func (o *GetTerraformVersionUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] getTerraformVersionUsingGET1Unauthorized ", 401)
}

func (o *GetTerraformVersionUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTerraformVersionUsingGET1Forbidden creates a GetTerraformVersionUsingGET1Forbidden with default headers values
func NewGetTerraformVersionUsingGET1Forbidden() *GetTerraformVersionUsingGET1Forbidden {
	return &GetTerraformVersionUsingGET1Forbidden{}
}

/*GetTerraformVersionUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type GetTerraformVersionUsingGET1Forbidden struct {
}

func (o *GetTerraformVersionUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] getTerraformVersionUsingGET1Forbidden ", 403)
}

func (o *GetTerraformVersionUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTerraformVersionUsingGET1NotFound creates a GetTerraformVersionUsingGET1NotFound with default headers values
func NewGetTerraformVersionUsingGET1NotFound() *GetTerraformVersionUsingGET1NotFound {
	return &GetTerraformVersionUsingGET1NotFound{}
}

/*GetTerraformVersionUsingGET1NotFound handles this case with default header values.

Not Found
*/
type GetTerraformVersionUsingGET1NotFound struct {
	Payload *models.Error
}

func (o *GetTerraformVersionUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] getTerraformVersionUsingGET1NotFound  %+v", 404, o.Payload)
}

func (o *GetTerraformVersionUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTerraformVersionUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}