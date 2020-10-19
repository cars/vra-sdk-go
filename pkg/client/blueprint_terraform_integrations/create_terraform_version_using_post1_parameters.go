// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateTerraformVersionUsingPOST1Params creates a new CreateTerraformVersionUsingPOST1Params object
// with the default values initialized.
func NewCreateTerraformVersionUsingPOST1Params() *CreateTerraformVersionUsingPOST1Params {
	var ()
	return &CreateTerraformVersionUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTerraformVersionUsingPOST1ParamsWithTimeout creates a new CreateTerraformVersionUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTerraformVersionUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateTerraformVersionUsingPOST1Params {
	var ()
	return &CreateTerraformVersionUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateTerraformVersionUsingPOST1ParamsWithContext creates a new CreateTerraformVersionUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTerraformVersionUsingPOST1ParamsWithContext(ctx context.Context) *CreateTerraformVersionUsingPOST1Params {
	var ()
	return &CreateTerraformVersionUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateTerraformVersionUsingPOST1ParamsWithHTTPClient creates a new CreateTerraformVersionUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTerraformVersionUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateTerraformVersionUsingPOST1Params {
	var ()
	return &CreateTerraformVersionUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateTerraformVersionUsingPOST1Params contains all the parameters to send to the API endpoint
for the create terraform version using p o s t 1 operation typically these are written to a http.Request
*/
type CreateTerraformVersionUsingPOST1Params struct {

	/*TerraformVersion
	  Terraform version

	*/
	TerraformVersion *models.TerraformVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateTerraformVersionUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) WithContext(ctx context.Context) *CreateTerraformVersionUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateTerraformVersionUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTerraformVersion adds the terraformVersion to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) WithTerraformVersion(terraformVersion *models.TerraformVersion) *CreateTerraformVersionUsingPOST1Params {
	o.SetTerraformVersion(terraformVersion)
	return o
}

// SetTerraformVersion adds the terraformVersion to the create terraform version using p o s t 1 params
func (o *CreateTerraformVersionUsingPOST1Params) SetTerraformVersion(terraformVersion *models.TerraformVersion) {
	o.TerraformVersion = terraformVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTerraformVersionUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TerraformVersion != nil {
		if err := r.SetBodyParam(o.TerraformVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
