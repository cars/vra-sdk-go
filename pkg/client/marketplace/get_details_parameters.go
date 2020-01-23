// Code generated by go-swagger; DO NOT EDIT.

package marketplace

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
)

// NewGetDetailsParams creates a new GetDetailsParams object
// with the default values initialized.
func NewGetDetailsParams() *GetDetailsParams {
	var ()
	return &GetDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDetailsParamsWithTimeout creates a new GetDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDetailsParamsWithTimeout(timeout time.Duration) *GetDetailsParams {
	var ()
	return &GetDetailsParams{

		timeout: timeout,
	}
}

// NewGetDetailsParamsWithContext creates a new GetDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDetailsParamsWithContext(ctx context.Context) *GetDetailsParams {
	var ()
	return &GetDetailsParams{

		Context: ctx,
	}
}

// NewGetDetailsParamsWithHTTPClient creates a new GetDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDetailsParamsWithHTTPClient(client *http.Client) *GetDetailsParams {
	var ()
	return &GetDetailsParams{
		HTTPClient: client,
	}
}

/*GetDetailsParams contains all the parameters to send to the API endpoint
for the get details operation typically these are written to a http.Request
*/
type GetDetailsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about

	*/
	APIVersion *string
	/*ContentID
	  Content Id

	*/
	ContentID string
	/*SourceID
	  Content Source Id

	*/
	SourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get details params
func (o *GetDetailsParams) WithTimeout(timeout time.Duration) *GetDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get details params
func (o *GetDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get details params
func (o *GetDetailsParams) WithContext(ctx context.Context) *GetDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get details params
func (o *GetDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get details params
func (o *GetDetailsParams) WithHTTPClient(client *http.Client) *GetDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get details params
func (o *GetDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get details params
func (o *GetDetailsParams) WithAPIVersion(aPIVersion *string) *GetDetailsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get details params
func (o *GetDetailsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithContentID adds the contentID to the get details params
func (o *GetDetailsParams) WithContentID(contentID string) *GetDetailsParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the get details params
func (o *GetDetailsParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithSourceID adds the sourceID to the get details params
func (o *GetDetailsParams) WithSourceID(sourceID strfmt.UUID) *GetDetailsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get details params
func (o *GetDetailsParams) SetSourceID(sourceID strfmt.UUID) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
		return err
	}

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
