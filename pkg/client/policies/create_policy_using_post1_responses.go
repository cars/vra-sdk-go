// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreatePolicyUsingPOST1Reader is a Reader for the CreatePolicyUsingPOST1 structure.
type CreatePolicyUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePolicyUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePolicyUsingPOST1OK creates a CreatePolicyUsingPOST1OK with default headers values
func NewCreatePolicyUsingPOST1OK() *CreatePolicyUsingPOST1OK {
	return &CreatePolicyUsingPOST1OK{}
}

/*CreatePolicyUsingPOST1OK handles this case with default header values.

OK
*/
type CreatePolicyUsingPOST1OK struct {
	Payload *models.Policy
}

func (o *CreatePolicyUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /policy/api/policies][%d] createPolicyUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *CreatePolicyUsingPOST1OK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *CreatePolicyUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}