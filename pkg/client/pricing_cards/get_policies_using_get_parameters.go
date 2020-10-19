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
	"github.com/go-openapi/swag"
)

// NewGetPoliciesUsingGETParams creates a new GetPoliciesUsingGETParams object
// with the default values initialized.
func NewGetPoliciesUsingGETParams() *GetPoliciesUsingGETParams {
	var ()
	return &GetPoliciesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoliciesUsingGETParamsWithTimeout creates a new GetPoliciesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoliciesUsingGETParamsWithTimeout(timeout time.Duration) *GetPoliciesUsingGETParams {
	var ()
	return &GetPoliciesUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPoliciesUsingGETParamsWithContext creates a new GetPoliciesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoliciesUsingGETParamsWithContext(ctx context.Context) *GetPoliciesUsingGETParams {
	var ()
	return &GetPoliciesUsingGETParams{

		Context: ctx,
	}
}

// NewGetPoliciesUsingGETParamsWithHTTPClient creates a new GetPoliciesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoliciesUsingGETParamsWithHTTPClient(client *http.Client) *GetPoliciesUsingGETParams {
	var ()
	return &GetPoliciesUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPoliciesUsingGETParams contains all the parameters to send to the API endpoint
for the get policies using g e t operation typically these are written to a http.Request
*/
type GetPoliciesUsingGETParams struct {

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
	/*ExpandAssignmentInfo
	  Whether or not returns count of assignments.

	*/
	ExpandAssignmentInfo *bool
	/*ExpandPricingCard
	  Whether or not returns detailed pricing card for each result.

	*/
	ExpandPricingCard *bool
	/*Search
	  Search by name and description

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithTimeout(timeout time.Duration) *GetPoliciesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithContext(ctx context.Context) *GetPoliciesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithHTTPClient(client *http.Client) *GetPoliciesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetPoliciesUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetPoliciesUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithDollarTop(dollarTop *int32) *GetPoliciesUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithAPIVersion(aPIVersion *string) *GetPoliciesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpandAssignmentInfo adds the expandAssignmentInfo to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithExpandAssignmentInfo(expandAssignmentInfo *bool) *GetPoliciesUsingGETParams {
	o.SetExpandAssignmentInfo(expandAssignmentInfo)
	return o
}

// SetExpandAssignmentInfo adds the expandAssignmentInfo to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetExpandAssignmentInfo(expandAssignmentInfo *bool) {
	o.ExpandAssignmentInfo = expandAssignmentInfo
}

// WithExpandPricingCard adds the expandPricingCard to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithExpandPricingCard(expandPricingCard *bool) *GetPoliciesUsingGETParams {
	o.SetExpandPricingCard(expandPricingCard)
	return o
}

// SetExpandPricingCard adds the expandPricingCard to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetExpandPricingCard(expandPricingCard *bool) {
	o.ExpandPricingCard = expandPricingCard
}

// WithSearch adds the search to the get policies using get params
func (o *GetPoliciesUsingGETParams) WithSearch(search *string) *GetPoliciesUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get policies using get params
func (o *GetPoliciesUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoliciesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExpandAssignmentInfo != nil {

		// query param expandAssignmentInfo
		var qrExpandAssignmentInfo bool
		if o.ExpandAssignmentInfo != nil {
			qrExpandAssignmentInfo = *o.ExpandAssignmentInfo
		}
		qExpandAssignmentInfo := swag.FormatBool(qrExpandAssignmentInfo)
		if qExpandAssignmentInfo != "" {
			if err := r.SetQueryParam("expandAssignmentInfo", qExpandAssignmentInfo); err != nil {
				return err
			}
		}

	}

	if o.ExpandPricingCard != nil {

		// query param expandPricingCard
		var qrExpandPricingCard bool
		if o.ExpandPricingCard != nil {
			qrExpandPricingCard = *o.ExpandPricingCard
		}
		qExpandPricingCard := swag.FormatBool(qrExpandPricingCard)
		if qExpandPricingCard != "" {
			if err := r.SetQueryParam("expandPricingCard", qExpandPricingCard); err != nil {
				return err
			}
		}

	}

	if o.Search != nil {

		// query param search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
