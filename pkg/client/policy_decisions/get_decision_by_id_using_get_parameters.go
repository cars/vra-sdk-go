// Code generated by go-swagger; DO NOT EDIT.

package policy_decisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDecisionByIDUsingGETParams creates a new GetDecisionByIDUsingGETParams object
// with the default values initialized.
func NewGetDecisionByIDUsingGETParams() *GetDecisionByIDUsingGETParams {
	var ()
	return &GetDecisionByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDecisionByIDUsingGETParamsWithTimeout creates a new GetDecisionByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDecisionByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetDecisionByIDUsingGETParams {
	var ()
	return &GetDecisionByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDecisionByIDUsingGETParamsWithContext creates a new GetDecisionByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDecisionByIDUsingGETParamsWithContext(ctx context.Context) *GetDecisionByIDUsingGETParams {
	var ()
	return &GetDecisionByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetDecisionByIDUsingGETParamsWithHTTPClient creates a new GetDecisionByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDecisionByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetDecisionByIDUsingGETParams {
	var ()
	return &GetDecisionByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDecisionByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get decision by Id using g e t operation typically these are written to a http.Request
*/
type GetDecisionByIDUsingGETParams struct {

	/*ID
	  Policy decision Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetDecisionByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) WithContext(ctx context.Context) *GetDecisionByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetDecisionByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) WithID(id strfmt.UUID) *GetDecisionByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get decision by Id using get params
func (o *GetDecisionByIDUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDecisionByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}