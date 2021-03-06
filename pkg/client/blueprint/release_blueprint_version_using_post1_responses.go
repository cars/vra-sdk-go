// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ReleaseBlueprintVersionUsingPOST1Reader is a Reader for the ReleaseBlueprintVersionUsingPOST1 structure.
type ReleaseBlueprintVersionUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseBlueprintVersionUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReleaseBlueprintVersionUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReleaseBlueprintVersionUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReleaseBlueprintVersionUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReleaseBlueprintVersionUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReleaseBlueprintVersionUsingPOST1OK creates a ReleaseBlueprintVersionUsingPOST1OK with default headers values
func NewReleaseBlueprintVersionUsingPOST1OK() *ReleaseBlueprintVersionUsingPOST1OK {
	return &ReleaseBlueprintVersionUsingPOST1OK{}
}

/* ReleaseBlueprintVersionUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type ReleaseBlueprintVersionUsingPOST1OK struct {
	Payload *models.BlueprintVersion
}

func (o *ReleaseBlueprintVersionUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/release][%d] releaseBlueprintVersionUsingPOST1OK  %+v", 200, o.Payload)
}
func (o *ReleaseBlueprintVersionUsingPOST1OK) GetPayload() *models.BlueprintVersion {
	return o.Payload
}

func (o *ReleaseBlueprintVersionUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseBlueprintVersionUsingPOST1Unauthorized creates a ReleaseBlueprintVersionUsingPOST1Unauthorized with default headers values
func NewReleaseBlueprintVersionUsingPOST1Unauthorized() *ReleaseBlueprintVersionUsingPOST1Unauthorized {
	return &ReleaseBlueprintVersionUsingPOST1Unauthorized{}
}

/* ReleaseBlueprintVersionUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ReleaseBlueprintVersionUsingPOST1Unauthorized struct {
}

func (o *ReleaseBlueprintVersionUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/release][%d] releaseBlueprintVersionUsingPOST1Unauthorized ", 401)
}

func (o *ReleaseBlueprintVersionUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseBlueprintVersionUsingPOST1Forbidden creates a ReleaseBlueprintVersionUsingPOST1Forbidden with default headers values
func NewReleaseBlueprintVersionUsingPOST1Forbidden() *ReleaseBlueprintVersionUsingPOST1Forbidden {
	return &ReleaseBlueprintVersionUsingPOST1Forbidden{}
}

/* ReleaseBlueprintVersionUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReleaseBlueprintVersionUsingPOST1Forbidden struct {
}

func (o *ReleaseBlueprintVersionUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/release][%d] releaseBlueprintVersionUsingPOST1Forbidden ", 403)
}

func (o *ReleaseBlueprintVersionUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseBlueprintVersionUsingPOST1NotFound creates a ReleaseBlueprintVersionUsingPOST1NotFound with default headers values
func NewReleaseBlueprintVersionUsingPOST1NotFound() *ReleaseBlueprintVersionUsingPOST1NotFound {
	return &ReleaseBlueprintVersionUsingPOST1NotFound{}
}

/* ReleaseBlueprintVersionUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReleaseBlueprintVersionUsingPOST1NotFound struct {
	Payload *models.Error
}

func (o *ReleaseBlueprintVersionUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/release][%d] releaseBlueprintVersionUsingPOST1NotFound  %+v", 404, o.Payload)
}
func (o *ReleaseBlueprintVersionUsingPOST1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReleaseBlueprintVersionUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
