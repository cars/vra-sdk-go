// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetDeploymentExpenseHistoryByIDUsingGETParams creates a new GetDeploymentExpenseHistoryByIDUsingGETParams object
// with the default values initialized.
func NewGetDeploymentExpenseHistoryByIDUsingGETParams() *GetDeploymentExpenseHistoryByIDUsingGETParams {
	var ()
	return &GetDeploymentExpenseHistoryByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithTimeout creates a new GetDeploymentExpenseHistoryByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	var ()
	return &GetDeploymentExpenseHistoryByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithContext creates a new GetDeploymentExpenseHistoryByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithContext(ctx context.Context) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	var ()
	return &GetDeploymentExpenseHistoryByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithHTTPClient creates a new GetDeploymentExpenseHistoryByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentExpenseHistoryByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	var ()
	return &GetDeploymentExpenseHistoryByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDeploymentExpenseHistoryByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get deployment expense history by Id using g e t operation typically these are written to a http.Request
*/
type GetDeploymentExpenseHistoryByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*From
	  The timestamp from when history is requested. Should be of ISO_INSTANT format.

	*/
	From *strfmt.DateTime
	/*Interval
	  The interval of the expense history. Should be one of daily, weekly or monthly.

	*/
	Interval *string
	/*To
	  The timestamp until when history is requested. Should be of ISO_INSTANT format.

	*/
	To *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithContext(ctx context.Context) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithDepID(depID strfmt.UUID) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithFrom adds the from to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithFrom(from *strfmt.DateTime) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithInterval adds the interval to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithInterval(interval *string) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithTo adds the to to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WithTo(to *strfmt.DateTime) *GetDeploymentExpenseHistoryByIDUsingGETParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get deployment expense history by Id using get params
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentExpenseHistoryByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	if o.From != nil {

		// query param from
		var qrFrom strfmt.DateTime
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom.String()
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Interval != nil {

		// query param interval
		var qrInterval string
		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {
			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}

	}

	if o.To != nil {

		// query param to
		var qrTo strfmt.DateTime
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo.String()
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
