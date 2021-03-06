// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetVSphereCloudAccountReader is a Reader for the GetVSphereCloudAccount structure.
type GetVSphereCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVSphereCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVSphereCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVSphereCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVSphereCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVSphereCloudAccountOK creates a GetVSphereCloudAccountOK with default headers values
func NewGetVSphereCloudAccountOK() *GetVSphereCloudAccountOK {
	return &GetVSphereCloudAccountOK{}
}

/* GetVSphereCloudAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVSphereCloudAccountOK struct {
	Payload *models.CloudAccountVsphere
}

func (o *GetVSphereCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vsphere/{id}][%d] getVSphereCloudAccountOK  %+v", 200, o.Payload)
}
func (o *GetVSphereCloudAccountOK) GetPayload() *models.CloudAccountVsphere {
	return o.Payload
}

func (o *GetVSphereCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVsphere)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereCloudAccountForbidden creates a GetVSphereCloudAccountForbidden with default headers values
func NewGetVSphereCloudAccountForbidden() *GetVSphereCloudAccountForbidden {
	return &GetVSphereCloudAccountForbidden{}
}

/* GetVSphereCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVSphereCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetVSphereCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vsphere/{id}][%d] getVSphereCloudAccountForbidden  %+v", 403, o.Payload)
}
func (o *GetVSphereCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetVSphereCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereCloudAccountNotFound creates a GetVSphereCloudAccountNotFound with default headers values
func NewGetVSphereCloudAccountNotFound() *GetVSphereCloudAccountNotFound {
	return &GetVSphereCloudAccountNotFound{}
}

/* GetVSphereCloudAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVSphereCloudAccountNotFound struct {
	Payload *models.Error
}

func (o *GetVSphereCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vsphere/{id}][%d] getVSphereCloudAccountNotFound  %+v", 404, o.Payload)
}
func (o *GetVSphereCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVSphereCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
