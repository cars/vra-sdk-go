// Code generated by go-swagger; DO NOT EDIT.

package policy_types

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

// NewGetPolicyTypeByIDUsingGETParams creates a new GetPolicyTypeByIDUsingGETParams object
// with the default values initialized.
func NewGetPolicyTypeByIDUsingGETParams() *GetPolicyTypeByIDUsingGETParams {
	var ()
	return &GetPolicyTypeByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyTypeByIDUsingGETParamsWithTimeout creates a new GetPolicyTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyTypeByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetPolicyTypeByIDUsingGETParams {
	var ()
	return &GetPolicyTypeByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPolicyTypeByIDUsingGETParamsWithContext creates a new GetPolicyTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyTypeByIDUsingGETParamsWithContext(ctx context.Context) *GetPolicyTypeByIDUsingGETParams {
	var ()
	return &GetPolicyTypeByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetPolicyTypeByIDUsingGETParamsWithHTTPClient creates a new GetPolicyTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyTypeByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetPolicyTypeByIDUsingGETParams {
	var ()
	return &GetPolicyTypeByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPolicyTypeByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get policy type by Id using g e t operation typically these are written to a http.Request
*/
type GetPolicyTypeByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*ID
	  Policy type ID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetPolicyTypeByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) WithContext(ctx context.Context) *GetPolicyTypeByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetPolicyTypeByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetPolicyTypeByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) WithID(id string) *GetPolicyTypeByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get policy type by Id using get params
func (o *GetPolicyTypeByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyTypeByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
