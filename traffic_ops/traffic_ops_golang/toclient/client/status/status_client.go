// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteStatus Delete a Status
*/
func (a *Client) DeleteStatus(params *DeleteStatusParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteStatus",
		Method:             "DELETE",
		PathPattern:        "/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStatusOK), nil

}

/*
GetStatusByID Retrieve a specific Status
*/
func (a *Client) GetStatusByID(params *GetStatusByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatusByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStatusById",
		Method:             "GET",
		PathPattern:        "/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatusByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatusByIDOK), nil

}

/*
GetStatuses Retrieve a list of Statuses
*/
func (a *Client) GetStatuses(params *GetStatusesParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatusesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStatuses",
		Method:             "GET",
		PathPattern:        "/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatusesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatusesOK), nil

}

/*
PostStatus Create a Status
*/
func (a *Client) PostStatus(params *PostStatusParams, authInfo runtime.ClientAuthInfoWriter) (*PostStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStatus",
		Method:             "POST",
		PathPattern:        "/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStatusOK), nil

}

/*
PutStatus Update a Status
*/
func (a *Client) PutStatus(params *PutStatusParams, authInfo runtime.ClientAuthInfoWriter) (*PutStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutStatus",
		Method:             "PUT",
		PathPattern:        "/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutStatusOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
