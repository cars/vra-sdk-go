// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateProject creates project

Create project
*/
func (a *Client) CreateProject(params *CreateProjectParams) (*CreateProjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProject",
		Method:             "POST",
		PathPattern:        "/iaas/api/projects",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateProjectCreated), nil

}

/*
DeleteProject deletes project

Delete project with a given id
*/
func (a *Client) DeleteProject(params *DeleteProjectParams) (*DeleteProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProject",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/projects/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProjectOK), nil

}

/*
GetProject gets project

Get project with a given id
*/
func (a *Client) GetProject(params *GetProjectParams) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProject",
		Method:             "GET",
		PathPattern:        "/iaas/api/projects/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectOK), nil

}

/*
GetProjects gets projects

Get all projects
*/
func (a *Client) GetProjects(params *GetProjectsParams) (*GetProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjects",
		Method:             "GET",
		PathPattern:        "/iaas/api/projects",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectsOK), nil

}

/*
UpdateProject updates project

Update project
*/
func (a *Client) UpdateProject(params *UpdateProjectParams) (*UpdateProjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProject",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/projects/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProjectCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
