// Code generated by go-swagger; DO NOT EDIT.

package deployment_requests

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
	"github.com/go-openapi/swag"
)

// NewGetRequestEventsUsingGETParams creates a new GetRequestEventsUsingGETParams object
// with the default values initialized.
func NewGetRequestEventsUsingGETParams() *GetRequestEventsUsingGETParams {
	var ()
	return &GetRequestEventsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRequestEventsUsingGETParamsWithTimeout creates a new GetRequestEventsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRequestEventsUsingGETParamsWithTimeout(timeout time.Duration) *GetRequestEventsUsingGETParams {
	var ()
	return &GetRequestEventsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetRequestEventsUsingGETParamsWithContext creates a new GetRequestEventsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRequestEventsUsingGETParamsWithContext(ctx context.Context) *GetRequestEventsUsingGETParams {
	var ()
	return &GetRequestEventsUsingGETParams{

		Context: ctx,
	}
}

// NewGetRequestEventsUsingGETParamsWithHTTPClient creates a new GetRequestEventsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRequestEventsUsingGETParamsWithHTTPClient(client *http.Client) *GetRequestEventsUsingGETParams {
	var ()
	return &GetRequestEventsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetRequestEventsUsingGETParams contains all the parameters to send to the API endpoint
for the get request events using g e t operation typically these are written to a http.Request
*/
type GetRequestEventsUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.

	*/
	DollarOrderby []string
	/*DollarSkip
	  Number of records you want to skip

	*/
	DollarSkip *int32
	/*DollarTop
	  Number of records you want

	*/
	DollarTop *int32
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*Deleted
	  Retrieves the soft-deleted events of the request that have not yet been completely deleted.

	*/
	Deleted *bool
	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*RequestID
	  Request ID

	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithTimeout(timeout time.Duration) *GetRequestEventsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithContext(ctx context.Context) *GetRequestEventsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithHTTPClient(client *http.Client) *GetRequestEventsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetRequestEventsUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetRequestEventsUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithDollarTop(dollarTop *int32) *GetRequestEventsUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithAPIVersion(aPIVersion *string) *GetRequestEventsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeleted adds the deleted to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithDeleted(deleted *bool) *GetRequestEventsUsingGETParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDepID adds the depID to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithDepID(depID strfmt.UUID) *GetRequestEventsUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithRequestID adds the requestID to the get request events using get params
func (o *GetRequestEventsUsingGETParams) WithRequestID(requestID strfmt.UUID) *GetRequestEventsUsingGETParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get request events using get params
func (o *GetRequestEventsUsingGETParams) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRequestEventsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarOrderby := o.DollarOrderby

	joinedDollarOrderby := swag.JoinByFormat(valuesDollarOrderby, "multi")
	// query array param $orderby
	if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
		return err
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32
		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {
			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
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

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
