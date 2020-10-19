// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

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

// NewGetDeploymentActionUsingGETParams creates a new GetDeploymentActionUsingGETParams object
// with the default values initialized.
func NewGetDeploymentActionUsingGETParams() *GetDeploymentActionUsingGETParams {
	var ()
	return &GetDeploymentActionUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentActionUsingGETParamsWithTimeout creates a new GetDeploymentActionUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentActionUsingGETParamsWithTimeout(timeout time.Duration) *GetDeploymentActionUsingGETParams {
	var ()
	return &GetDeploymentActionUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDeploymentActionUsingGETParamsWithContext creates a new GetDeploymentActionUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentActionUsingGETParamsWithContext(ctx context.Context) *GetDeploymentActionUsingGETParams {
	var ()
	return &GetDeploymentActionUsingGETParams{

		Context: ctx,
	}
}

// NewGetDeploymentActionUsingGETParamsWithHTTPClient creates a new GetDeploymentActionUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentActionUsingGETParamsWithHTTPClient(client *http.Client) *GetDeploymentActionUsingGETParams {
	var ()
	return &GetDeploymentActionUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDeploymentActionUsingGETParams contains all the parameters to send to the API endpoint
for the get deployment action using g e t operation typically these are written to a http.Request
*/
type GetDeploymentActionUsingGETParams struct {

	/*ActionID
	  Action ID

	*/
	ActionID string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithTimeout(timeout time.Duration) *GetDeploymentActionUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithContext(ctx context.Context) *GetDeploymentActionUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithHTTPClient(client *http.Client) *GetDeploymentActionUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithActionID(actionID string) *GetDeploymentActionUsingGETParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithAPIVersion adds the aPIVersion to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithAPIVersion(aPIVersion *string) *GetDeploymentActionUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) WithDepID(depID strfmt.UUID) *GetDeploymentActionUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get deployment action using get params
func (o *GetDeploymentActionUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentActionUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

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

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
