// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewDeleteNsxVCloudAccountParams creates a new DeleteNsxVCloudAccountParams object
// with the default values initialized.
func NewDeleteNsxVCloudAccountParams() *DeleteNsxVCloudAccountParams {
	var ()
	return &DeleteNsxVCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNsxVCloudAccountParamsWithTimeout creates a new DeleteNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNsxVCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteNsxVCloudAccountParams {
	var ()
	return &DeleteNsxVCloudAccountParams{

		timeout: timeout,
	}
}

// NewDeleteNsxVCloudAccountParamsWithContext creates a new DeleteNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNsxVCloudAccountParamsWithContext(ctx context.Context) *DeleteNsxVCloudAccountParams {
	var ()
	return &DeleteNsxVCloudAccountParams{

		Context: ctx,
	}
}

// NewDeleteNsxVCloudAccountParamsWithHTTPClient creates a new DeleteNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNsxVCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteNsxVCloudAccountParams {
	var ()
	return &DeleteNsxVCloudAccountParams{
		HTTPClient: client,
	}
}

/*DeleteNsxVCloudAccountParams contains all the parameters to send to the API endpoint
for the delete nsx v cloud account operation typically these are written to a http.Request
*/
type DeleteNsxVCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*ID
	  The ID of the Cloud Account

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteNsxVCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) WithContext(ctx context.Context) *DeleteNsxVCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteNsxVCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) WithAPIVersion(aPIVersion string) *DeleteNsxVCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) WithID(id string) *DeleteNsxVCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete nsx v cloud account params
func (o *DeleteNsxVCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNsxVCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
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
