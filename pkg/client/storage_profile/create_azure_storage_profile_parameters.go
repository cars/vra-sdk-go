// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// NewCreateAzureStorageProfileParams creates a new CreateAzureStorageProfileParams object
// with the default values initialized.
func NewCreateAzureStorageProfileParams() *CreateAzureStorageProfileParams {
	var ()
	return &CreateAzureStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureStorageProfileParamsWithTimeout creates a new CreateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAzureStorageProfileParamsWithTimeout(timeout time.Duration) *CreateAzureStorageProfileParams {
	var ()
	return &CreateAzureStorageProfileParams{

		timeout: timeout,
	}
}

// NewCreateAzureStorageProfileParamsWithContext creates a new CreateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAzureStorageProfileParamsWithContext(ctx context.Context) *CreateAzureStorageProfileParams {
	var ()
	return &CreateAzureStorageProfileParams{

		Context: ctx,
	}
}

// NewCreateAzureStorageProfileParamsWithHTTPClient creates a new CreateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAzureStorageProfileParamsWithHTTPClient(client *http.Client) *CreateAzureStorageProfileParams {
	var ()
	return &CreateAzureStorageProfileParams{
		HTTPClient: client,
	}
}

/*CreateAzureStorageProfileParams contains all the parameters to send to the API endpoint
for the create azure storage profile operation typically these are written to a http.Request
*/
type CreateAzureStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  StorageProfileAzureSpecification instance

	*/
	Body *models.StorageProfileAzureSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) WithTimeout(timeout time.Duration) *CreateAzureStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) WithContext(ctx context.Context) *CreateAzureStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) WithHTTPClient(client *http.Client) *CreateAzureStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) WithAPIVersion(aPIVersion string) *CreateAzureStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) WithBody(body *models.StorageProfileAzureSpecification) *CreateAzureStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create azure storage profile params
func (o *CreateAzureStorageProfileParams) SetBody(body *models.StorageProfileAzureSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
