// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPoliciesUsingGETReader is a Reader for the GetPoliciesUsingGET structure.
type GetPoliciesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPoliciesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPoliciesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPoliciesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPoliciesUsingGETOK creates a GetPoliciesUsingGETOK with default headers values
func NewGetPoliciesUsingGETOK() *GetPoliciesUsingGETOK {
	return &GetPoliciesUsingGETOK{}
}

/* GetPoliciesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetPoliciesUsingGETOK struct {
	Payload *models.PageOfMeteringPolicy
}

func (o *GetPoliciesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-cards][%d] getPoliciesUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetPoliciesUsingGETOK) GetPayload() *models.PageOfMeteringPolicy {
	return o.Payload
}

func (o *GetPoliciesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfMeteringPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPoliciesUsingGETUnauthorized creates a GetPoliciesUsingGETUnauthorized with default headers values
func NewGetPoliciesUsingGETUnauthorized() *GetPoliciesUsingGETUnauthorized {
	return &GetPoliciesUsingGETUnauthorized{}
}

/* GetPoliciesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPoliciesUsingGETUnauthorized struct {
}

func (o *GetPoliciesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-cards][%d] getPoliciesUsingGETUnauthorized ", 401)
}

func (o *GetPoliciesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
