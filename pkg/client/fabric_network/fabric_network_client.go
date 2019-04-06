// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fabric network API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric network API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFabricNetwork gets fabric network

Get fabric network with a given id
*/
func (a *Client) GetFabricNetwork(params *GetFabricNetworkParams) (*GetFabricNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricNetwork",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-networks/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFabricNetworkOK), nil

}

/*
GetFabricNetworks gets fabric networks

Get all fabric networks.
*/
func (a *Client) GetFabricNetworks(params *GetFabricNetworksParams) (*GetFabricNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricNetworks",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-networks",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricNetworksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFabricNetworksOK), nil

}

/*
UpdateFabricNetwork updates fabric network

Update fabric network. Only tag updates are supported.
*/
func (a *Client) UpdateFabricNetwork(params *UpdateFabricNetworkParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFabricNetworkParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFabricNetwork",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/fabric-networks/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFabricNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
