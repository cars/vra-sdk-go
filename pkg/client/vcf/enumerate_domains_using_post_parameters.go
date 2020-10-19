// Code generated by go-swagger; DO NOT EDIT.

package vcf

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

// NewEnumerateDomainsUsingPOSTParams creates a new EnumerateDomainsUsingPOSTParams object
// with the default values initialized.
func NewEnumerateDomainsUsingPOSTParams() *EnumerateDomainsUsingPOSTParams {
	var ()
	return &EnumerateDomainsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateDomainsUsingPOSTParamsWithTimeout creates a new EnumerateDomainsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnumerateDomainsUsingPOSTParamsWithTimeout(timeout time.Duration) *EnumerateDomainsUsingPOSTParams {
	var ()
	return &EnumerateDomainsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewEnumerateDomainsUsingPOSTParamsWithContext creates a new EnumerateDomainsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnumerateDomainsUsingPOSTParamsWithContext(ctx context.Context) *EnumerateDomainsUsingPOSTParams {
	var ()
	return &EnumerateDomainsUsingPOSTParams{

		Context: ctx,
	}
}

// NewEnumerateDomainsUsingPOSTParamsWithHTTPClient creates a new EnumerateDomainsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnumerateDomainsUsingPOSTParamsWithHTTPClient(client *http.Client) *EnumerateDomainsUsingPOSTParams {
	var ()
	return &EnumerateDomainsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*EnumerateDomainsUsingPOSTParams contains all the parameters to send to the API endpoint
for the enumerate domains using p o s t operation typically these are written to a http.Request
*/
type EnumerateDomainsUsingPOSTParams struct {

	/*Input
	  input

	*/
	Input *models.PhotonModelEndpointConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) WithTimeout(timeout time.Duration) *EnumerateDomainsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) WithContext(ctx context.Context) *EnumerateDomainsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) WithHTTPClient(client *http.Client) *EnumerateDomainsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) WithInput(input *models.PhotonModelEndpointConfigRequest) *EnumerateDomainsUsingPOSTParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the enumerate domains using p o s t params
func (o *EnumerateDomainsUsingPOSTParams) SetInput(input *models.PhotonModelEndpointConfigRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateDomainsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
