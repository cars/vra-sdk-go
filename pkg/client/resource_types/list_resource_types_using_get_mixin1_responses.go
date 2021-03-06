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

// ListResourceTypesUsingGETMixin1Reader is a Reader for the ListResourceTypesUsingGETMixin1 structure.
type ListResourceTypesUsingGETMixin1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourceTypesUsingGETMixin1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourceTypesUsingGETMixin1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListResourceTypesUsingGETMixin1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListResourceTypesUsingGETMixin1OK creates a ListResourceTypesUsingGETMixin1OK with default headers values
func NewListResourceTypesUsingGETMixin1OK() *ListResourceTypesUsingGETMixin1OK {
	return &ListResourceTypesUsingGETMixin1OK{}
}

/* ListResourceTypesUsingGETMixin1OK describes a response with status code 200, with default header values.

OK
*/
type ListResourceTypesUsingGETMixin1OK struct {
	Payload *models.PageOfDeploymentResourceType
}

func (o *ListResourceTypesUsingGETMixin1OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resource-types][%d] listResourceTypesUsingGETMixin1OK  %+v", 200, o.Payload)
}
func (o *ListResourceTypesUsingGETMixin1OK) GetPayload() *models.PageOfDeploymentResourceType {
	return o.Payload
}

func (o *ListResourceTypesUsingGETMixin1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfDeploymentResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourceTypesUsingGETMixin1Unauthorized creates a ListResourceTypesUsingGETMixin1Unauthorized with default headers values
func NewListResourceTypesUsingGETMixin1Unauthorized() *ListResourceTypesUsingGETMixin1Unauthorized {
	return &ListResourceTypesUsingGETMixin1Unauthorized{}
}

/* ListResourceTypesUsingGETMixin1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListResourceTypesUsingGETMixin1Unauthorized struct {
}

func (o *ListResourceTypesUsingGETMixin1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resource-types][%d] listResourceTypesUsingGETMixin1Unauthorized ", 401)
}

func (o *ListResourceTypesUsingGETMixin1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
