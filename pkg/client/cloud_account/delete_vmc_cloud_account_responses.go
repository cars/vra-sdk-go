// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteVmcCloudAccountReader is a Reader for the DeleteVmcCloudAccount structure.
type DeleteVmcCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVmcCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteVmcCloudAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteVmcCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVmcCloudAccountNoContent creates a DeleteVmcCloudAccountNoContent with default headers values
func NewDeleteVmcCloudAccountNoContent() *DeleteVmcCloudAccountNoContent {
	return &DeleteVmcCloudAccountNoContent{}
}

/* DeleteVmcCloudAccountNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteVmcCloudAccountNoContent struct {
}

func (o *DeleteVmcCloudAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vmc/{id}][%d] deleteVmcCloudAccountNoContent ", 204)
}

func (o *DeleteVmcCloudAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVmcCloudAccountForbidden creates a DeleteVmcCloudAccountForbidden with default headers values
func NewDeleteVmcCloudAccountForbidden() *DeleteVmcCloudAccountForbidden {
	return &DeleteVmcCloudAccountForbidden{}
}

/* DeleteVmcCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteVmcCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *DeleteVmcCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vmc/{id}][%d] deleteVmcCloudAccountForbidden  %+v", 403, o.Payload)
}
func (o *DeleteVmcCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteVmcCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
