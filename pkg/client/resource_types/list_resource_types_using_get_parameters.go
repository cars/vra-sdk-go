// Code generated by go-swagger; DO NOT EDIT.

package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListResourceTypesUsingGETParams creates a new ListResourceTypesUsingGETParams object
// with the default values initialized.
func NewListResourceTypesUsingGETParams() *ListResourceTypesUsingGETParams {
	var (
		expandDefault = bool(true)
	)
	return &ListResourceTypesUsingGETParams{
		Expand: &expandDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListResourceTypesUsingGETParamsWithTimeout creates a new ListResourceTypesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListResourceTypesUsingGETParamsWithTimeout(timeout time.Duration) *ListResourceTypesUsingGETParams {
	var (
		expandDefault = bool(true)
	)
	return &ListResourceTypesUsingGETParams{
		Expand: &expandDefault,

		timeout: timeout,
	}
}

// NewListResourceTypesUsingGETParamsWithContext creates a new ListResourceTypesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListResourceTypesUsingGETParamsWithContext(ctx context.Context) *ListResourceTypesUsingGETParams {
	var (
		expandDefault = bool(true)
	)
	return &ListResourceTypesUsingGETParams{
		Expand: &expandDefault,

		Context: ctx,
	}
}

// NewListResourceTypesUsingGETParamsWithHTTPClient creates a new ListResourceTypesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListResourceTypesUsingGETParamsWithHTTPClient(client *http.Client) *ListResourceTypesUsingGETParams {
	var (
		expandDefault = bool(true)
	)
	return &ListResourceTypesUsingGETParams{
		Expand:     &expandDefault,
		HTTPClient: client,
	}
}

/*ListResourceTypesUsingGETParams contains all the parameters to send to the API endpoint
for the list resource types using g e t operation typically these are written to a http.Request
*/
type ListResourceTypesUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is ascending on name. Sorting is supported on fields createdAt, updatedAt, createdBy, updatedBy, name.

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
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about

	*/
	APIVersion *string
	/*Expand
	  Expand with content

	*/
	Expand *bool
	/*Name
	  Filter by name

	*/
	Name *string
	/*ProviderID
	  Filter by provider ID

	*/
	ProviderID *string
	/*Search
	  Search by name and description

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithTimeout(timeout time.Duration) *ListResourceTypesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithContext(ctx context.Context) *ListResourceTypesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithHTTPClient(client *http.Client) *ListResourceTypesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithDollarOrderby(dollarOrderby []string) *ListResourceTypesUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithDollarSkip(dollarSkip *int32) *ListResourceTypesUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithDollarTop(dollarTop *int32) *ListResourceTypesUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithAPIVersion(aPIVersion *string) *ListResourceTypesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpand adds the expand to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithExpand(expand *bool) *ListResourceTypesUsingGETParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetExpand(expand *bool) {
	o.Expand = expand
}

// WithName adds the name to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithName(name *string) *ListResourceTypesUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WithProviderID adds the providerID to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithProviderID(providerID *string) *ListResourceTypesUsingGETParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithSearch adds the search to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) WithSearch(search *string) *ListResourceTypesUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the list resource types using get params
func (o *ListResourceTypesUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ListResourceTypesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Expand != nil {

		// query param expand
		var qrExpand bool
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := swag.FormatBool(qrExpand)
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.ProviderID != nil {

		// query param providerId
		var qrProviderID string
		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {
			if err := r.SetQueryParam("providerId", qProviderID); err != nil {
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