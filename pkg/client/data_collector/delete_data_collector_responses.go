// Code generated by go-swagger; DO NOT EDIT.

package data_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteDataCollectorReader is a Reader for the DeleteDataCollector structure.
type DeleteDataCollectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDataCollectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDataCollectorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteDataCollectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDataCollectorNoContent creates a DeleteDataCollectorNoContent with default headers values
func NewDeleteDataCollectorNoContent() *DeleteDataCollectorNoContent {
	return &DeleteDataCollectorNoContent{}
}

/* DeleteDataCollectorNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteDataCollectorNoContent struct {
}

func (o *DeleteDataCollectorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/data-collectors/{id}][%d] deleteDataCollectorNoContent ", 204)
}

func (o *DeleteDataCollectorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDataCollectorForbidden creates a DeleteDataCollectorForbidden with default headers values
func NewDeleteDataCollectorForbidden() *DeleteDataCollectorForbidden {
	return &DeleteDataCollectorForbidden{}
}

/* DeleteDataCollectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDataCollectorForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *DeleteDataCollectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/data-collectors/{id}][%d] deleteDataCollectorForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDataCollectorForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteDataCollectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
