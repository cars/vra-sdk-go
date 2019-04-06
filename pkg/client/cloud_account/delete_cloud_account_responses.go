// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCloudAccountReader is a Reader for the DeleteCloudAccount structure.
type DeleteCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCloudAccountOK creates a DeleteCloudAccountOK with default headers values
func NewDeleteCloudAccountOK() *DeleteCloudAccountOK {
	return &DeleteCloudAccountOK{}
}

/*DeleteCloudAccountOK handles this case with default header values.

successful operation
*/
type DeleteCloudAccountOK struct {
}

func (o *DeleteCloudAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts/{id}][%d] deleteCloudAccountOK ", 200)
}

func (o *DeleteCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudAccountForbidden creates a DeleteCloudAccountForbidden with default headers values
func NewDeleteCloudAccountForbidden() *DeleteCloudAccountForbidden {
	return &DeleteCloudAccountForbidden{}
}

/*DeleteCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type DeleteCloudAccountForbidden struct {
}

func (o *DeleteCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts/{id}][%d] deleteCloudAccountForbidden ", 403)
}

func (o *DeleteCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudAccountNotFound creates a DeleteCloudAccountNotFound with default headers values
func NewDeleteCloudAccountNotFound() *DeleteCloudAccountNotFound {
	return &DeleteCloudAccountNotFound{}
}

/*DeleteCloudAccountNotFound handles this case with default header values.

Not Found
*/
type DeleteCloudAccountNotFound struct {
}

func (o *DeleteCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts/{id}][%d] deleteCloudAccountNotFound ", 404)
}

func (o *DeleteCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
