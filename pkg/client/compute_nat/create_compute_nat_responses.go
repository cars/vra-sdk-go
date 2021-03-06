// Code generated by go-swagger; DO NOT EDIT.

package compute_nat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateComputeNatReader is a Reader for the CreateComputeNat structure.
type CreateComputeNatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateComputeNatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateComputeNatAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateComputeNatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateComputeNatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateComputeNatAccepted creates a CreateComputeNatAccepted with default headers values
func NewCreateComputeNatAccepted() *CreateComputeNatAccepted {
	return &CreateComputeNatAccepted{}
}

/* CreateComputeNatAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateComputeNatAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateComputeNatAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/compute-nats][%d] createComputeNatAccepted  %+v", 202, o.Payload)
}
func (o *CreateComputeNatAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateComputeNatAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComputeNatBadRequest creates a CreateComputeNatBadRequest with default headers values
func NewCreateComputeNatBadRequest() *CreateComputeNatBadRequest {
	return &CreateComputeNatBadRequest{}
}

/* CreateComputeNatBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateComputeNatBadRequest struct {
	Payload *models.Error
}

func (o *CreateComputeNatBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/compute-nats][%d] createComputeNatBadRequest  %+v", 400, o.Payload)
}
func (o *CreateComputeNatBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateComputeNatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComputeNatForbidden creates a CreateComputeNatForbidden with default headers values
func NewCreateComputeNatForbidden() *CreateComputeNatForbidden {
	return &CreateComputeNatForbidden{}
}

/* CreateComputeNatForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateComputeNatForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *CreateComputeNatForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/compute-nats][%d] createComputeNatForbidden  %+v", 403, o.Payload)
}
func (o *CreateComputeNatForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateComputeNatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
