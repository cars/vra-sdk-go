// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUsingDELETEReader is a Reader for the DeleteUsingDELETE structure.
type DeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsingDELETENoContent creates a DeleteUsingDELETENoContent with default headers values
func NewDeleteUsingDELETENoContent() *DeleteUsingDELETENoContent {
	return &DeleteUsingDELETENoContent{}
}

/* DeleteUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUsingDELETENoContent struct {
}

func (o *DeleteUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/sources/{sourceId}][%d] deleteUsingDELETENoContent ", 204)
}

func (o *DeleteUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETEUnauthorized creates a DeleteUsingDELETEUnauthorized with default headers values
func NewDeleteUsingDELETEUnauthorized() *DeleteUsingDELETEUnauthorized {
	return &DeleteUsingDELETEUnauthorized{}
}

/* DeleteUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUsingDELETEUnauthorized struct {
}

func (o *DeleteUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/sources/{sourceId}][%d] deleteUsingDELETEUnauthorized ", 401)
}

func (o *DeleteUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
