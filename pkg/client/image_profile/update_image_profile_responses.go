// Code generated by go-swagger; DO NOT EDIT.

package image_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateImageProfileReader is a Reader for the UpdateImageProfile structure.
type UpdateImageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateImageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateImageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateImageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateImageProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateImageProfileOK creates a UpdateImageProfileOK with default headers values
func NewUpdateImageProfileOK() *UpdateImageProfileOK {
	return &UpdateImageProfileOK{}
}

/* UpdateImageProfileOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateImageProfileOK struct {
	Payload *models.ImageProfile
}

func (o *UpdateImageProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/image-profiles/{id}][%d] updateImageProfileOK  %+v", 200, o.Payload)
}
func (o *UpdateImageProfileOK) GetPayload() *models.ImageProfile {
	return o.Payload
}

func (o *UpdateImageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImageProfileForbidden creates a UpdateImageProfileForbidden with default headers values
func NewUpdateImageProfileForbidden() *UpdateImageProfileForbidden {
	return &UpdateImageProfileForbidden{}
}

/* UpdateImageProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateImageProfileForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateImageProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/image-profiles/{id}][%d] updateImageProfileForbidden  %+v", 403, o.Payload)
}
func (o *UpdateImageProfileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateImageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImageProfileNotFound creates a UpdateImageProfileNotFound with default headers values
func NewUpdateImageProfileNotFound() *UpdateImageProfileNotFound {
	return &UpdateImageProfileNotFound{}
}

/* UpdateImageProfileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateImageProfileNotFound struct {
	Payload *models.Error
}

func (o *UpdateImageProfileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/image-profiles/{id}][%d] updateImageProfileNotFound  %+v", 404, o.Payload)
}
func (o *UpdateImageProfileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateImageProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
