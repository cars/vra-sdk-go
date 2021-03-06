// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pricing card assignments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pricing card assignments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeMeteringAssignmentStrategyUsingPATCH(params *ChangeMeteringAssignmentStrategyUsingPATCHParams, opts ...ClientOption) (*ChangeMeteringAssignmentStrategyUsingPATCHOK, error)

	CreateMeteringAssignmentStrategyUsingPOST(params *CreateMeteringAssignmentStrategyUsingPOSTParams, opts ...ClientOption) (*CreateMeteringAssignmentStrategyUsingPOSTOK, *CreateMeteringAssignmentStrategyUsingPOSTCreated, error)

	CreateMeteringPolicyAssignmentUsingPOST(params *CreateMeteringPolicyAssignmentUsingPOSTParams, opts ...ClientOption) (*CreateMeteringPolicyAssignmentUsingPOSTOK, *CreateMeteringPolicyAssignmentUsingPOSTCreated, error)

	DeleteMeteringPolicyAssignmentUsingDELETE(params *DeleteMeteringPolicyAssignmentUsingDELETEParams, opts ...ClientOption) (*DeleteMeteringPolicyAssignmentUsingDELETEOK, error)

	GetAllMeteringPolicyAssignmentsUsingGET(params *GetAllMeteringPolicyAssignmentsUsingGETParams, opts ...ClientOption) (*GetAllMeteringPolicyAssignmentsUsingGETOK, error)

	GetMeteringAssignmentStrategyUsingGET(params *GetMeteringAssignmentStrategyUsingGETParams, opts ...ClientOption) (*GetMeteringAssignmentStrategyUsingGETOK, error)

	GetMeteringPolicyAssignmentUsingGET(params *GetMeteringPolicyAssignmentUsingGETParams, opts ...ClientOption) (*GetMeteringPolicyAssignmentUsingGETOK, error)

	PatchMeteringPolicyAssignmentUsingPATCH(params *PatchMeteringPolicyAssignmentUsingPATCHParams, opts ...ClientOption) (*PatchMeteringPolicyAssignmentUsingPATCHOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeMeteringAssignmentStrategyUsingPATCH updates the pricing card assignment strategy for the org
*/
func (a *Client) ChangeMeteringAssignmentStrategyUsingPATCH(params *ChangeMeteringAssignmentStrategyUsingPATCHParams, opts ...ClientOption) (*ChangeMeteringAssignmentStrategyUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMeteringAssignmentStrategyUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeMeteringAssignmentStrategyUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeMeteringAssignmentStrategyUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*ChangeMeteringAssignmentStrategyUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeMeteringAssignmentStrategyUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateMeteringAssignmentStrategyUsingPOST selectings the new pricing card assignment strategy p r o j e c t or c l o u d z o n e are possible values can be used while creating strategy also there can be only one strategy for a given org at a given point of time

  Create a new pricing card assignment strategy based on request body and validate its field according to business rules.
*/
func (a *Client) CreateMeteringAssignmentStrategyUsingPOST(params *CreateMeteringAssignmentStrategyUsingPOSTParams, opts ...ClientOption) (*CreateMeteringAssignmentStrategyUsingPOSTOK, *CreateMeteringAssignmentStrategyUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMeteringAssignmentStrategyUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMeteringAssignmentStrategyUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMeteringAssignmentStrategyUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateMeteringAssignmentStrategyUsingPOSTOK:
		return value, nil, nil
	case *CreateMeteringAssignmentStrategyUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_card_assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateMeteringPolicyAssignmentUsingPOST creates a new pricing card assignment

  Create a new pricing card policy assignment based on request body and validate its field according to business rules. Request body with ALL entityType will delete the older assignments for the given pricingCardId
*/
func (a *Client) CreateMeteringPolicyAssignmentUsingPOST(params *CreateMeteringPolicyAssignmentUsingPOSTParams, opts ...ClientOption) (*CreateMeteringPolicyAssignmentUsingPOSTOK, *CreateMeteringPolicyAssignmentUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMeteringPolicyAssignmentUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMeteringPolicyAssignmentUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-card-assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMeteringPolicyAssignmentUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateMeteringPolicyAssignmentUsingPOSTOK:
		return value, nil, nil
	case *CreateMeteringPolicyAssignmentUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_card_assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteMeteringPolicyAssignmentUsingDELETE deletes the pricing card assignment with specified id

  Deletes the pricing card assignment with the specified id
*/
func (a *Client) DeleteMeteringPolicyAssignmentUsingDELETE(params *DeleteMeteringPolicyAssignmentUsingDELETEParams, opts ...ClientOption) (*DeleteMeteringPolicyAssignmentUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMeteringPolicyAssignmentUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMeteringPolicyAssignmentUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMeteringPolicyAssignmentUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeleteMeteringPolicyAssignmentUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMeteringPolicyAssignmentUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllMeteringPolicyAssignmentsUsingGET fetches all pricing card assignment for private cloud

  Returns a paginated list of pricing card assignments
*/
func (a *Client) GetAllMeteringPolicyAssignmentsUsingGET(params *GetAllMeteringPolicyAssignmentsUsingGETParams, opts ...ClientOption) (*GetAllMeteringPolicyAssignmentsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllMeteringPolicyAssignmentsUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllMeteringPolicyAssignmentsUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllMeteringPolicyAssignmentsUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllMeteringPolicyAssignmentsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllMeteringPolicyAssignmentsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMeteringAssignmentStrategyUsingGET fetches pricing card assignment strategy for the org

  Returns a pricing card assignment strategy for the Org
*/
func (a *Client) GetMeteringAssignmentStrategyUsingGET(params *GetMeteringAssignmentStrategyUsingGETParams, opts ...ClientOption) (*GetMeteringAssignmentStrategyUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeteringAssignmentStrategyUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMeteringAssignmentStrategyUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMeteringAssignmentStrategyUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetMeteringAssignmentStrategyUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMeteringAssignmentStrategyUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMeteringPolicyAssignmentUsingGET fetches pricing card assignment for private cloud by id

  Returns a pricing card assignments by id
*/
func (a *Client) GetMeteringPolicyAssignmentUsingGET(params *GetMeteringPolicyAssignmentUsingGETParams, opts ...ClientOption) (*GetMeteringPolicyAssignmentUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeteringPolicyAssignmentUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMeteringPolicyAssignmentUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMeteringPolicyAssignmentUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetMeteringPolicyAssignmentUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMeteringPolicyAssignmentUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchMeteringPolicyAssignmentUsingPATCH updates the pricing card assignment id with the supplied id request body with a l l entity type will delete the older assignments for the given pricing card Id
*/
func (a *Client) PatchMeteringPolicyAssignmentUsingPATCH(params *PatchMeteringPolicyAssignmentUsingPATCHParams, opts ...ClientOption) (*PatchMeteringPolicyAssignmentUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchMeteringPolicyAssignmentUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchMeteringPolicyAssignmentUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchMeteringPolicyAssignmentUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*PatchMeteringPolicyAssignmentUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchMeteringPolicyAssignmentUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
