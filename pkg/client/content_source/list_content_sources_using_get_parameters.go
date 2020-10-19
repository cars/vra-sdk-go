// Code generated by go-swagger; DO NOT EDIT.

package content_source

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

// NewListContentSourcesUsingGETParams creates a new ListContentSourcesUsingGETParams object
// with the default values initialized.
func NewListContentSourcesUsingGETParams() *ListContentSourcesUsingGETParams {
	var ()
	return &ListContentSourcesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListContentSourcesUsingGETParamsWithTimeout creates a new ListContentSourcesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListContentSourcesUsingGETParamsWithTimeout(timeout time.Duration) *ListContentSourcesUsingGETParams {
	var ()
	return &ListContentSourcesUsingGETParams{

		timeout: timeout,
	}
}

// NewListContentSourcesUsingGETParamsWithContext creates a new ListContentSourcesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListContentSourcesUsingGETParamsWithContext(ctx context.Context) *ListContentSourcesUsingGETParams {
	var ()
	return &ListContentSourcesUsingGETParams{

		Context: ctx,
	}
}

// NewListContentSourcesUsingGETParamsWithHTTPClient creates a new ListContentSourcesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListContentSourcesUsingGETParamsWithHTTPClient(client *http.Client) *ListContentSourcesUsingGETParams {
	var ()
	return &ListContentSourcesUsingGETParams{
		HTTPClient: client,
	}
}

/*ListContentSourcesUsingGETParams contains all the parameters to send to the API endpoint
for the list content sources using g e t operation typically these are written to a http.Request
*/
type ListContentSourcesUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about

	*/
	APIVersion *string
	/*ContentType
	  Search based on Content Type

	*/
	ContentType *string
	/*IntegrationID
	  Search based on integrationId

	*/
	IntegrationID *string
	/*Name
	  Search based on Content Source name

	*/
	Name *string
	/*ProjectIds
	  Search based on associated projects

	*/
	ProjectIds []string
	/*Repository
	  Search based on repository

	*/
	Repository *string
	/*Search
	  Search based on name or repository

	*/
	Search *string
	/*SyncEnabled
	  Search based on whether sync is enabled or not

	*/
	SyncEnabled *bool
	/*TypeIds
	  Search based on Content Source Type Ids

	*/
	TypeIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithTimeout(timeout time.Duration) *ListContentSourcesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithContext(ctx context.Context) *ListContentSourcesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithHTTPClient(client *http.Client) *ListContentSourcesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithAPIVersion(aPIVersion *string) *ListContentSourcesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithContentType adds the contentType to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithContentType(contentType *string) *ListContentSourcesUsingGETParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithIntegrationID adds the integrationID to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithIntegrationID(integrationID *string) *ListContentSourcesUsingGETParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetIntegrationID(integrationID *string) {
	o.IntegrationID = integrationID
}

// WithName adds the name to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithName(name *string) *ListContentSourcesUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WithProjectIds adds the projectIds to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithProjectIds(projectIds []string) *ListContentSourcesUsingGETParams {
	o.SetProjectIds(projectIds)
	return o
}

// SetProjectIds adds the projectIds to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetProjectIds(projectIds []string) {
	o.ProjectIds = projectIds
}

// WithRepository adds the repository to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithRepository(repository *string) *ListContentSourcesUsingGETParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetRepository(repository *string) {
	o.Repository = repository
}

// WithSearch adds the search to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithSearch(search *string) *ListContentSourcesUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WithSyncEnabled adds the syncEnabled to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithSyncEnabled(syncEnabled *bool) *ListContentSourcesUsingGETParams {
	o.SetSyncEnabled(syncEnabled)
	return o
}

// SetSyncEnabled adds the syncEnabled to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetSyncEnabled(syncEnabled *bool) {
	o.SyncEnabled = syncEnabled
}

// WithTypeIds adds the typeIds to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) WithTypeIds(typeIds []string) *ListContentSourcesUsingGETParams {
	o.SetTypeIds(typeIds)
	return o
}

// SetTypeIds adds the typeIds to the list content sources using get params
func (o *ListContentSourcesUsingGETParams) SetTypeIds(typeIds []string) {
	o.TypeIds = typeIds
}

// WriteToRequest writes these params to a swagger request
func (o *ListContentSourcesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ContentType != nil {

		// query param contentType
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("contentType", qContentType); err != nil {
				return err
			}
		}

	}

	if o.IntegrationID != nil {

		// query param integrationId
		var qrIntegrationID string
		if o.IntegrationID != nil {
			qrIntegrationID = *o.IntegrationID
		}
		qIntegrationID := qrIntegrationID
		if qIntegrationID != "" {
			if err := r.SetQueryParam("integrationId", qIntegrationID); err != nil {
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

	valuesProjectIds := o.ProjectIds

	joinedProjectIds := swag.JoinByFormat(valuesProjectIds, "multi")
	// query array param projectIds
	if err := r.SetQueryParam("projectIds", joinedProjectIds...); err != nil {
		return err
	}

	if o.Repository != nil {

		// query param repository
		var qrRepository string
		if o.Repository != nil {
			qrRepository = *o.Repository
		}
		qRepository := qrRepository
		if qRepository != "" {
			if err := r.SetQueryParam("repository", qRepository); err != nil {
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

	if o.SyncEnabled != nil {

		// query param syncEnabled
		var qrSyncEnabled bool
		if o.SyncEnabled != nil {
			qrSyncEnabled = *o.SyncEnabled
		}
		qSyncEnabled := swag.FormatBool(qrSyncEnabled)
		if qSyncEnabled != "" {
			if err := r.SetQueryParam("syncEnabled", qSyncEnabled); err != nil {
				return err
			}
		}

	}

	valuesTypeIds := o.TypeIds

	joinedTypeIds := swag.JoinByFormat(valuesTypeIds, "multi")
	// query array param typeIds
	if err := r.SetQueryParam("typeIds", joinedTypeIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
