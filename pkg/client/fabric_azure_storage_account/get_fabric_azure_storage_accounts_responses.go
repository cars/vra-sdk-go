// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetFabricAzureStorageAccountsReader is a Reader for the GetFabricAzureStorageAccounts structure.
type GetFabricAzureStorageAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricAzureStorageAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFabricAzureStorageAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetFabricAzureStorageAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricAzureStorageAccountsOK creates a GetFabricAzureStorageAccountsOK with default headers values
func NewGetFabricAzureStorageAccountsOK() *GetFabricAzureStorageAccountsOK {
	return &GetFabricAzureStorageAccountsOK{}
}

/*GetFabricAzureStorageAccountsOK handles this case with default header values.

successful operation
*/
type GetFabricAzureStorageAccountsOK struct {
	Payload *models.FabricAzureStorageAccountResult
}

func (o *GetFabricAzureStorageAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts][%d] getFabricAzureStorageAccountsOK  %+v", 200, o.Payload)
}

func (o *GetFabricAzureStorageAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricAzureStorageAccountResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricAzureStorageAccountsForbidden creates a GetFabricAzureStorageAccountsForbidden with default headers values
func NewGetFabricAzureStorageAccountsForbidden() *GetFabricAzureStorageAccountsForbidden {
	return &GetFabricAzureStorageAccountsForbidden{}
}

/*GetFabricAzureStorageAccountsForbidden handles this case with default header values.

Forbidden
*/
type GetFabricAzureStorageAccountsForbidden struct {
}

func (o *GetFabricAzureStorageAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts][%d] getFabricAzureStorageAccountsForbidden ", 403)
}

func (o *GetFabricAzureStorageAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
