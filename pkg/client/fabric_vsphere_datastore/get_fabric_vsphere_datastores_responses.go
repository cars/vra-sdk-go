// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_datastore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricVSphereDatastoresReader is a Reader for the GetFabricVSphereDatastores structure.
type GetFabricVSphereDatastoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricVSphereDatastoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricVSphereDatastoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricVSphereDatastoresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFabricVSphereDatastoresOK creates a GetFabricVSphereDatastoresOK with default headers values
func NewGetFabricVSphereDatastoresOK() *GetFabricVSphereDatastoresOK {
	return &GetFabricVSphereDatastoresOK{}
}

/* GetFabricVSphereDatastoresOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFabricVSphereDatastoresOK struct {
	Payload *models.FabricVsphereDatastoreResult
}

func (o *GetFabricVSphereDatastoresOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores][%d] getFabricVSphereDatastoresOK  %+v", 200, o.Payload)
}
func (o *GetFabricVSphereDatastoresOK) GetPayload() *models.FabricVsphereDatastoreResult {
	return o.Payload
}

func (o *GetFabricVSphereDatastoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricVsphereDatastoreResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricVSphereDatastoresForbidden creates a GetFabricVSphereDatastoresForbidden with default headers values
func NewGetFabricVSphereDatastoresForbidden() *GetFabricVSphereDatastoresForbidden {
	return &GetFabricVSphereDatastoresForbidden{}
}

/* GetFabricVSphereDatastoresForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFabricVSphereDatastoresForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetFabricVSphereDatastoresForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores][%d] getFabricVSphereDatastoresForbidden  %+v", 403, o.Payload)
}
func (o *GetFabricVSphereDatastoresForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetFabricVSphereDatastoresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
