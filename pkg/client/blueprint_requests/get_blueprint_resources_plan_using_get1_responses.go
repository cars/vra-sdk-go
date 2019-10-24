// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetBlueprintResourcesPlanUsingGET1Reader is a Reader for the GetBlueprintResourcesPlanUsingGET1 structure.
type GetBlueprintResourcesPlanUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintResourcesPlanUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprintResourcesPlanUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBlueprintResourcesPlanUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlueprintResourcesPlanUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlueprintResourcesPlanUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintResourcesPlanUsingGET1OK creates a GetBlueprintResourcesPlanUsingGET1OK with default headers values
func NewGetBlueprintResourcesPlanUsingGET1OK() *GetBlueprintResourcesPlanUsingGET1OK {
	return &GetBlueprintResourcesPlanUsingGET1OK{}
}

/*GetBlueprintResourcesPlanUsingGET1OK handles this case with default header values.

OK
*/
type GetBlueprintResourcesPlanUsingGET1OK struct {
	Payload *models.BlueprintResourcesPlan
}

func (o *GetBlueprintResourcesPlanUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests/{requestId}/resources-plan][%d] getBlueprintResourcesPlanUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintResourcesPlanUsingGET1OK) GetPayload() *models.BlueprintResourcesPlan {
	return o.Payload
}

func (o *GetBlueprintResourcesPlanUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintResourcesPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintResourcesPlanUsingGET1Unauthorized creates a GetBlueprintResourcesPlanUsingGET1Unauthorized with default headers values
func NewGetBlueprintResourcesPlanUsingGET1Unauthorized() *GetBlueprintResourcesPlanUsingGET1Unauthorized {
	return &GetBlueprintResourcesPlanUsingGET1Unauthorized{}
}

/*GetBlueprintResourcesPlanUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetBlueprintResourcesPlanUsingGET1Unauthorized struct {
}

func (o *GetBlueprintResourcesPlanUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests/{requestId}/resources-plan][%d] getBlueprintResourcesPlanUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintResourcesPlanUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintResourcesPlanUsingGET1Forbidden creates a GetBlueprintResourcesPlanUsingGET1Forbidden with default headers values
func NewGetBlueprintResourcesPlanUsingGET1Forbidden() *GetBlueprintResourcesPlanUsingGET1Forbidden {
	return &GetBlueprintResourcesPlanUsingGET1Forbidden{}
}

/*GetBlueprintResourcesPlanUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type GetBlueprintResourcesPlanUsingGET1Forbidden struct {
}

func (o *GetBlueprintResourcesPlanUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests/{requestId}/resources-plan][%d] getBlueprintResourcesPlanUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintResourcesPlanUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintResourcesPlanUsingGET1NotFound creates a GetBlueprintResourcesPlanUsingGET1NotFound with default headers values
func NewGetBlueprintResourcesPlanUsingGET1NotFound() *GetBlueprintResourcesPlanUsingGET1NotFound {
	return &GetBlueprintResourcesPlanUsingGET1NotFound{}
}

/*GetBlueprintResourcesPlanUsingGET1NotFound handles this case with default header values.

Not Found
*/
type GetBlueprintResourcesPlanUsingGET1NotFound struct {
}

func (o *GetBlueprintResourcesPlanUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests/{requestId}/resources-plan][%d] getBlueprintResourcesPlanUsingGET1NotFound ", 404)
}

func (o *GetBlueprintResourcesPlanUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}