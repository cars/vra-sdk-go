// Code generated by go-swagger; DO NOT EDIT.

package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetResourceTypeUsingGETMixin1Reader is a Reader for the GetResourceTypeUsingGETMixin1 structure.
type GetResourceTypeUsingGETMixin1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceTypeUsingGETMixin1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceTypeUsingGETMixin1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceTypeUsingGETMixin1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceTypeUsingGETMixin1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceTypeUsingGETMixin1OK creates a GetResourceTypeUsingGETMixin1OK with default headers values
func NewGetResourceTypeUsingGETMixin1OK() *GetResourceTypeUsingGETMixin1OK {
	return &GetResourceTypeUsingGETMixin1OK{}
}

/*GetResourceTypeUsingGETMixin1OK handles this case with default header values.

OK
*/
type GetResourceTypeUsingGETMixin1OK struct {
	Payload *models.DeploymentResourceType
}

func (o *GetResourceTypeUsingGETMixin1OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resource-types/{resourceTypeId}][%d] getResourceTypeUsingGETMixin1OK  %+v", 200, o.Payload)
}

func (o *GetResourceTypeUsingGETMixin1OK) GetPayload() *models.DeploymentResourceType {
	return o.Payload
}

func (o *GetResourceTypeUsingGETMixin1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceTypeUsingGETMixin1Unauthorized creates a GetResourceTypeUsingGETMixin1Unauthorized with default headers values
func NewGetResourceTypeUsingGETMixin1Unauthorized() *GetResourceTypeUsingGETMixin1Unauthorized {
	return &GetResourceTypeUsingGETMixin1Unauthorized{}
}

/*GetResourceTypeUsingGETMixin1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetResourceTypeUsingGETMixin1Unauthorized struct {
}

func (o *GetResourceTypeUsingGETMixin1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resource-types/{resourceTypeId}][%d] getResourceTypeUsingGETMixin1Unauthorized ", 401)
}

func (o *GetResourceTypeUsingGETMixin1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceTypeUsingGETMixin1NotFound creates a GetResourceTypeUsingGETMixin1NotFound with default headers values
func NewGetResourceTypeUsingGETMixin1NotFound() *GetResourceTypeUsingGETMixin1NotFound {
	return &GetResourceTypeUsingGETMixin1NotFound{}
}

/*GetResourceTypeUsingGETMixin1NotFound handles this case with default header values.

Not Found
*/
type GetResourceTypeUsingGETMixin1NotFound struct {
	Payload *models.Error
}

func (o *GetResourceTypeUsingGETMixin1NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resource-types/{resourceTypeId}][%d] getResourceTypeUsingGETMixin1NotFound  %+v", 404, o.Payload)
}

func (o *GetResourceTypeUsingGETMixin1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourceTypeUsingGETMixin1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
