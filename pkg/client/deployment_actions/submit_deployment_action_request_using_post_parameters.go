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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewSubmitDeploymentActionRequestUsingPOSTParams creates a new SubmitDeploymentActionRequestUsingPOSTParams object
// with the default values initialized.
func NewSubmitDeploymentActionRequestUsingPOSTParams() *SubmitDeploymentActionRequestUsingPOSTParams {
	var ()
	return &SubmitDeploymentActionRequestUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitDeploymentActionRequestUsingPOSTParamsWithTimeout creates a new SubmitDeploymentActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitDeploymentActionRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *SubmitDeploymentActionRequestUsingPOSTParams {
	var ()
	return &SubmitDeploymentActionRequestUsingPOSTParams{

		timeout: timeout,
	}
}

// NewSubmitDeploymentActionRequestUsingPOSTParamsWithContext creates a new SubmitDeploymentActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitDeploymentActionRequestUsingPOSTParamsWithContext(ctx context.Context) *SubmitDeploymentActionRequestUsingPOSTParams {
	var ()
	return &SubmitDeploymentActionRequestUsingPOSTParams{

		Context: ctx,
	}
}

// NewSubmitDeploymentActionRequestUsingPOSTParamsWithHTTPClient creates a new SubmitDeploymentActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitDeploymentActionRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *SubmitDeploymentActionRequestUsingPOSTParams {
	var ()
	return &SubmitDeploymentActionRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*SubmitDeploymentActionRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the submit deployment action request using p o s t operation typically these are written to a http.Request
*/
type SubmitDeploymentActionRequestUsingPOSTParams struct {

	/*ActionRequest
	  actionRequest

	*/
	ActionRequest *models.ResourceActionRequest
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

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

// WithTimeout adds the timeout to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithContext(ctx context.Context) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionRequest adds the actionRequest to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithActionRequest(actionRequest *models.ResourceActionRequest) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetActionRequest(actionRequest)
	return o
}

// SetActionRequest adds the actionRequest to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetActionRequest(actionRequest *models.ResourceActionRequest) {
	o.ActionRequest = actionRequest
}

// WithAPIVersion adds the aPIVersion to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithAPIVersion(aPIVersion *string) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WithDepID(depID strfmt.UUID) *SubmitDeploymentActionRequestUsingPOSTParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the submit deployment action request using p o s t params
func (o *SubmitDeploymentActionRequestUsingPOSTParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitDeploymentActionRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionRequest != nil {
		if err := r.SetBodyParam(o.ActionRequest); err != nil {
			return err
		}
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
