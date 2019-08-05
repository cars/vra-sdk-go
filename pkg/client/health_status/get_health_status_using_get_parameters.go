// Code generated by go-swagger; DO NOT EDIT.

package health_status

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

// NewGetHealthStatusUsingGETParams creates a new GetHealthStatusUsingGETParams object
// with the default values initialized.
func NewGetHealthStatusUsingGETParams() *GetHealthStatusUsingGETParams {

	return &GetHealthStatusUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthStatusUsingGETParamsWithTimeout creates a new GetHealthStatusUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHealthStatusUsingGETParamsWithTimeout(timeout time.Duration) *GetHealthStatusUsingGETParams {

	return &GetHealthStatusUsingGETParams{

		timeout: timeout,
	}
}

// NewGetHealthStatusUsingGETParamsWithContext creates a new GetHealthStatusUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHealthStatusUsingGETParamsWithContext(ctx context.Context) *GetHealthStatusUsingGETParams {

	return &GetHealthStatusUsingGETParams{

		Context: ctx,
	}
}

// NewGetHealthStatusUsingGETParamsWithHTTPClient creates a new GetHealthStatusUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHealthStatusUsingGETParamsWithHTTPClient(client *http.Client) *GetHealthStatusUsingGETParams {

	return &GetHealthStatusUsingGETParams{
		HTTPClient: client,
	}
}

/*GetHealthStatusUsingGETParams contains all the parameters to send to the API endpoint
for the get health status using g e t operation typically these are written to a http.Request
*/
type GetHealthStatusUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get health status using get params
func (o *GetHealthStatusUsingGETParams) WithTimeout(timeout time.Duration) *GetHealthStatusUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health status using get params
func (o *GetHealthStatusUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health status using get params
func (o *GetHealthStatusUsingGETParams) WithContext(ctx context.Context) *GetHealthStatusUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health status using get params
func (o *GetHealthStatusUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health status using get params
func (o *GetHealthStatusUsingGETParams) WithHTTPClient(client *http.Client) *GetHealthStatusUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health status using get params
func (o *GetHealthStatusUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthStatusUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}