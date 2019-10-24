// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEntitlementUsingDELETEReader is a Reader for the DeleteEntitlementUsingDELETE structure.
type DeleteEntitlementUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntitlementUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEntitlementUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEntitlementUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEntitlementUsingDELETENoContent creates a DeleteEntitlementUsingDELETENoContent with default headers values
func NewDeleteEntitlementUsingDELETENoContent() *DeleteEntitlementUsingDELETENoContent {
	return &DeleteEntitlementUsingDELETENoContent{}
}

/*DeleteEntitlementUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteEntitlementUsingDELETENoContent struct {
}

func (o *DeleteEntitlementUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/entitlements/{id}][%d] deleteEntitlementUsingDELETENoContent ", 204)
}

func (o *DeleteEntitlementUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntitlementUsingDELETEUnauthorized creates a DeleteEntitlementUsingDELETEUnauthorized with default headers values
func NewDeleteEntitlementUsingDELETEUnauthorized() *DeleteEntitlementUsingDELETEUnauthorized {
	return &DeleteEntitlementUsingDELETEUnauthorized{}
}

/*DeleteEntitlementUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteEntitlementUsingDELETEUnauthorized struct {
}

func (o *DeleteEntitlementUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/entitlements/{id}][%d] deleteEntitlementUsingDELETEUnauthorized ", 401)
}

func (o *DeleteEntitlementUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}