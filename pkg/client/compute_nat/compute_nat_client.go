// Code generated by go-swagger; DO NOT EDIT.

package compute_nat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new compute nat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute nat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateComputeNat(params *CreateComputeNatParams, opts ...ClientOption) (*CreateComputeNatAccepted, error)

	DeleteComputeNat(params *DeleteComputeNatParams, opts ...ClientOption) (*DeleteComputeNatAccepted, error)

	GetComputeNat(params *GetComputeNatParams, opts ...ClientOption) (*GetComputeNatOK, error)

	GetComputeNats(params *GetComputeNatsParams, opts ...ClientOption) (*GetComputeNatsOK, error)

	ReconfigureNat(params *ReconfigureNatParams, opts ...ClientOption) (*ReconfigureNatAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateComputeNat creates a compute nat

  Create a new Compute Nat.
*/
func (a *Client) CreateComputeNat(params *CreateComputeNatParams, opts ...ClientOption) (*CreateComputeNatAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateComputeNatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createComputeNat",
		Method:             "POST",
		PathPattern:        "/iaas/api/compute-nats",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateComputeNatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateComputeNatAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createComputeNat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteComputeNat deletes a compute nat

  Delete compute nat with a given id
*/
func (a *Client) DeleteComputeNat(params *DeleteComputeNatParams, opts ...ClientOption) (*DeleteComputeNatAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputeNatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComputeNat",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/compute-nats/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputeNatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteComputeNatAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComputeNat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComputeNat gets a compute nat

  Get Compute Nat with a given id
*/
func (a *Client) GetComputeNat(params *GetComputeNatParams, opts ...ClientOption) (*GetComputeNatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeNatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComputeNat",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-nats/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeNatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputeNatOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeNat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComputeNats gets compute nats

  Get all Compute Nats
*/
func (a *Client) GetComputeNats(params *GetComputeNatsParams, opts ...ClientOption) (*GetComputeNatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeNatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComputeNats",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-nats",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeNatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputeNatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeNats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReconfigureNat reconfigures operation for nat

  Day-2 reconfigure operation for nat
*/
func (a *Client) ReconfigureNat(params *ReconfigureNatParams, opts ...ClientOption) (*ReconfigureNatAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReconfigureNatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reconfigureNat",
		Method:             "POST",
		PathPattern:        "/iaas/api/compute-nats/{id}/operations/reconfigure",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReconfigureNatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReconfigureNatAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reconfigureNat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}