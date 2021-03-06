// Code generated by go-swagger; DO NOT EDIT.

package resource_types

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

// NewGetResourceTypeUsingGETMixin1Params creates a new GetResourceTypeUsingGETMixin1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceTypeUsingGETMixin1Params() *GetResourceTypeUsingGETMixin1Params {
	return &GetResourceTypeUsingGETMixin1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceTypeUsingGETMixin1ParamsWithTimeout creates a new GetResourceTypeUsingGETMixin1Params object
// with the ability to set a timeout on a request.
func NewGetResourceTypeUsingGETMixin1ParamsWithTimeout(timeout time.Duration) *GetResourceTypeUsingGETMixin1Params {
	return &GetResourceTypeUsingGETMixin1Params{
		timeout: timeout,
	}
}

// NewGetResourceTypeUsingGETMixin1ParamsWithContext creates a new GetResourceTypeUsingGETMixin1Params object
// with the ability to set a context for a request.
func NewGetResourceTypeUsingGETMixin1ParamsWithContext(ctx context.Context) *GetResourceTypeUsingGETMixin1Params {
	return &GetResourceTypeUsingGETMixin1Params{
		Context: ctx,
	}
}

// NewGetResourceTypeUsingGETMixin1ParamsWithHTTPClient creates a new GetResourceTypeUsingGETMixin1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceTypeUsingGETMixin1ParamsWithHTTPClient(client *http.Client) *GetResourceTypeUsingGETMixin1Params {
	return &GetResourceTypeUsingGETMixin1Params{
		HTTPClient: client,
	}
}

/* GetResourceTypeUsingGETMixin1Params contains all the parameters to send to the API endpoint
   for the get resource type using g e t mixin1 operation.

   Typically these are written to a http.Request.
*/
type GetResourceTypeUsingGETMixin1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ResourceTypeID.

	   resourceTypeId
	*/
	ResourceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource type using g e t mixin1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceTypeUsingGETMixin1Params) WithDefaults() *GetResourceTypeUsingGETMixin1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource type using g e t mixin1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceTypeUsingGETMixin1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) WithTimeout(timeout time.Duration) *GetResourceTypeUsingGETMixin1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) WithContext(ctx context.Context) *GetResourceTypeUsingGETMixin1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) WithHTTPClient(client *http.Client) *GetResourceTypeUsingGETMixin1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) WithAPIVersion(aPIVersion *string) *GetResourceTypeUsingGETMixin1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithResourceTypeID adds the resourceTypeID to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) WithResourceTypeID(resourceTypeID string) *GetResourceTypeUsingGETMixin1Params {
	o.SetResourceTypeID(resourceTypeID)
	return o
}

// SetResourceTypeID adds the resourceTypeId to the get resource type using g e t mixin1 params
func (o *GetResourceTypeUsingGETMixin1Params) SetResourceTypeID(resourceTypeID string) {
	o.ResourceTypeID = resourceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceTypeUsingGETMixin1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceTypeId
	if err := r.SetPathParam("resourceTypeId", o.ResourceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
