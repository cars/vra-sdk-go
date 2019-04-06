// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RevertMachineSnapshotReader is a Reader for the RevertMachineSnapshot structure.
type RevertMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevertMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 403:
		result := NewRevertMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRevertMachineSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevertMachineSnapshotForbidden creates a RevertMachineSnapshotForbidden with default headers values
func NewRevertMachineSnapshotForbidden() *RevertMachineSnapshotForbidden {
	return &RevertMachineSnapshotForbidden{}
}

/*RevertMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type RevertMachineSnapshotForbidden struct {
}

func (o *RevertMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/revert][%d] revertMachineSnapshotForbidden ", 403)
}

func (o *RevertMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevertMachineSnapshotNotFound creates a RevertMachineSnapshotNotFound with default headers values
func NewRevertMachineSnapshotNotFound() *RevertMachineSnapshotNotFound {
	return &RevertMachineSnapshotNotFound{}
}

/*RevertMachineSnapshotNotFound handles this case with default header values.

Not Found
*/
type RevertMachineSnapshotNotFound struct {
}

func (o *RevertMachineSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/revert][%d] revertMachineSnapshotNotFound ", 404)
}

func (o *RevertMachineSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
