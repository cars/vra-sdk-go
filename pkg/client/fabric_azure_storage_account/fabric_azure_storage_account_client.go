// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fabric azure storage account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric azure storage account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFabricAzureStorageAccount gets fabric azure storage account

Get fabric Azure storage account with a given id
*/
func (a *Client) GetFabricAzureStorageAccount(params *GetFabricAzureStorageAccountParams) (*GetFabricAzureStorageAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricAzureStorageAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricAzureStorageAccount",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-azure-storage-accounts/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricAzureStorageAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFabricAzureStorageAccountOK), nil

}

/*
GetFabricAzureStorageAccounts gets fabric azure storage accounts

Get all fabric Azure storage accounts.
*/
func (a *Client) GetFabricAzureStorageAccounts(params *GetFabricAzureStorageAccountsParams) (*GetFabricAzureStorageAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricAzureStorageAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricAzureStorageAccounts",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-azure-storage-accounts",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricAzureStorageAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFabricAzureStorageAccountsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
