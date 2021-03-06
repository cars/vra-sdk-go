// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

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

// NewPatchMeteringPolicyAssignmentUsingPATCHParams creates a new PatchMeteringPolicyAssignmentUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchMeteringPolicyAssignmentUsingPATCHParams() *PatchMeteringPolicyAssignmentUsingPATCHParams {
	return &PatchMeteringPolicyAssignmentUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithTimeout creates a new PatchMeteringPolicyAssignmentUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	return &PatchMeteringPolicyAssignmentUsingPATCHParams{
		timeout: timeout,
	}
}

// NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithContext creates a new PatchMeteringPolicyAssignmentUsingPATCHParams object
// with the ability to set a context for a request.
func NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithContext(ctx context.Context) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	return &PatchMeteringPolicyAssignmentUsingPATCHParams{
		Context: ctx,
	}
}

// NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithHTTPClient creates a new PatchMeteringPolicyAssignmentUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchMeteringPolicyAssignmentUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	return &PatchMeteringPolicyAssignmentUsingPATCHParams{
		HTTPClient: client,
	}
}

/* PatchMeteringPolicyAssignmentUsingPATCHParams contains all the parameters to send to the API endpoint
   for the patch metering policy assignment using p a t c h operation.

   Typically these are written to a http.Request.
*/
type PatchMeteringPolicyAssignmentUsingPATCHParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ID.

	   pricing card Assignment Id

	   Format: uuid
	*/
	ID strfmt.UUID

	/* MeteringPolicyAssignment.

	   A pricing card assignment with pricing card Id to override existing pricing card Id
	*/
	MeteringPolicyAssignment *models.MeteringPolicyAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch metering policy assignment using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithDefaults() *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch metering policy assignment using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithContext(ctx context.Context) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithAPIVersion(aPIVersion *string) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithID(id strfmt.UUID) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithMeteringPolicyAssignment adds the meteringPolicyAssignment to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WithMeteringPolicyAssignment(meteringPolicyAssignment *models.MeteringPolicyAssignment) *PatchMeteringPolicyAssignmentUsingPATCHParams {
	o.SetMeteringPolicyAssignment(meteringPolicyAssignment)
	return o
}

// SetMeteringPolicyAssignment adds the meteringPolicyAssignment to the patch metering policy assignment using p a t c h params
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) SetMeteringPolicyAssignment(meteringPolicyAssignment *models.MeteringPolicyAssignment) {
	o.MeteringPolicyAssignment = meteringPolicyAssignment
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMeteringPolicyAssignmentUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.MeteringPolicyAssignment != nil {
		if err := r.SetBodyParam(o.MeteringPolicyAssignment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
