// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateAzureCloudAccountReader is a Reader for the CreateAzureCloudAccount structure.
type CreateAzureCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAzureCloudAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAzureCloudAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAzureCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAzureCloudAccountCreated creates a CreateAzureCloudAccountCreated with default headers values
func NewCreateAzureCloudAccountCreated() *CreateAzureCloudAccountCreated {
	return &CreateAzureCloudAccountCreated{}
}

/*CreateAzureCloudAccountCreated handles this case with default header values.

successful operation
*/
type CreateAzureCloudAccountCreated struct {
	Payload *models.CloudAccount
}

func (o *CreateAzureCloudAccountCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-azure][%d] createAzureCloudAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateAzureCloudAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureCloudAccountBadRequest creates a CreateAzureCloudAccountBadRequest with default headers values
func NewCreateAzureCloudAccountBadRequest() *CreateAzureCloudAccountBadRequest {
	return &CreateAzureCloudAccountBadRequest{}
}

/*CreateAzureCloudAccountBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateAzureCloudAccountBadRequest struct {
}

func (o *CreateAzureCloudAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-azure][%d] createAzureCloudAccountBadRequest ", 400)
}

func (o *CreateAzureCloudAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAzureCloudAccountForbidden creates a CreateAzureCloudAccountForbidden with default headers values
func NewCreateAzureCloudAccountForbidden() *CreateAzureCloudAccountForbidden {
	return &CreateAzureCloudAccountForbidden{}
}

/*CreateAzureCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateAzureCloudAccountForbidden struct {
}

func (o *CreateAzureCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-azure][%d] createAzureCloudAccountForbidden ", 403)
}

func (o *CreateAzureCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
