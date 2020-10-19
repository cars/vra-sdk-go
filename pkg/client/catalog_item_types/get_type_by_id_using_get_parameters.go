// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetTypeByIDUsingGETParams creates a new GetTypeByIDUsingGETParams object
// with the default values initialized.
func NewGetTypeByIDUsingGETParams() *GetTypeByIDUsingGETParams {
	var ()
	return &GetTypeByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTypeByIDUsingGETParamsWithTimeout creates a new GetTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTypeByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetTypeByIDUsingGETParams {
	var ()
	return &GetTypeByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetTypeByIDUsingGETParamsWithContext creates a new GetTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTypeByIDUsingGETParamsWithContext(ctx context.Context) *GetTypeByIDUsingGETParams {
	var ()
	return &GetTypeByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetTypeByIDUsingGETParamsWithHTTPClient creates a new GetTypeByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTypeByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetTypeByIDUsingGETParams {
	var ()
	return &GetTypeByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetTypeByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get type by Id using g e t operation typically these are written to a http.Request
*/
type GetTypeByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*ID
	  Catalog Type ID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetTypeByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) WithContext(ctx context.Context) *GetTypeByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetTypeByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetTypeByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) WithID(id string) *GetTypeByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get type by Id using get params
func (o *GetTypeByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTypeByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
