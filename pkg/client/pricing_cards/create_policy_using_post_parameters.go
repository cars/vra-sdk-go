// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

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

// NewCreatePolicyUsingPOSTParams creates a new CreatePolicyUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePolicyUsingPOSTParams() *CreatePolicyUsingPOSTParams {
	return &CreatePolicyUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyUsingPOSTParamsWithTimeout creates a new CreatePolicyUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreatePolicyUsingPOSTParamsWithTimeout(timeout time.Duration) *CreatePolicyUsingPOSTParams {
	return &CreatePolicyUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreatePolicyUsingPOSTParamsWithContext creates a new CreatePolicyUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreatePolicyUsingPOSTParamsWithContext(ctx context.Context) *CreatePolicyUsingPOSTParams {
	return &CreatePolicyUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreatePolicyUsingPOSTParamsWithHTTPClient creates a new CreatePolicyUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePolicyUsingPOSTParamsWithHTTPClient(client *http.Client) *CreatePolicyUsingPOSTParams {
	return &CreatePolicyUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CreatePolicyUsingPOSTParams contains all the parameters to send to the API endpoint
   for the create policy using p o s t operation.

   Typically these are written to a http.Request.
*/
type CreatePolicyUsingPOSTParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* MeteringPolicy.

	   The pricing card to be created
	*/
	MeteringPolicy *models.MeteringPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policy using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyUsingPOSTParams) WithDefaults() *CreatePolicyUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policy using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithTimeout(timeout time.Duration) *CreatePolicyUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithContext(ctx context.Context) *CreatePolicyUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithHTTPClient(client *http.Client) *CreatePolicyUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreatePolicyUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithMeteringPolicy adds the meteringPolicy to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithMeteringPolicy(meteringPolicy *models.MeteringPolicy) *CreatePolicyUsingPOSTParams {
	o.SetMeteringPolicy(meteringPolicy)
	return o
}

// SetMeteringPolicy adds the meteringPolicy to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetMeteringPolicy(meteringPolicy *models.MeteringPolicy) {
	o.MeteringPolicy = meteringPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.MeteringPolicy != nil {
		if err := r.SetBodyParam(o.MeteringPolicy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
