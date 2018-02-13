// Code generated by go-swagger; DO NOT EDIT.

package division

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDivisionParams creates a new DeleteDivisionParams object
// with the default values initialized.
func NewDeleteDivisionParams() *DeleteDivisionParams {
	var ()
	return &DeleteDivisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDivisionParamsWithTimeout creates a new DeleteDivisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDivisionParamsWithTimeout(timeout time.Duration) *DeleteDivisionParams {
	var ()
	return &DeleteDivisionParams{

		timeout: timeout,
	}
}

// NewDeleteDivisionParamsWithContext creates a new DeleteDivisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDivisionParamsWithContext(ctx context.Context) *DeleteDivisionParams {
	var ()
	return &DeleteDivisionParams{

		Context: ctx,
	}
}

// NewDeleteDivisionParamsWithHTTPClient creates a new DeleteDivisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDivisionParamsWithHTTPClient(client *http.Client) *DeleteDivisionParams {
	var ()
	return &DeleteDivisionParams{
		HTTPClient: client,
	}
}

/*DeleteDivisionParams contains all the parameters to send to the API endpoint
for the delete division operation typically these are written to a http.Request
*/
type DeleteDivisionParams struct {

	/*ID
	  Id associated to the Division

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete division params
func (o *DeleteDivisionParams) WithTimeout(timeout time.Duration) *DeleteDivisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete division params
func (o *DeleteDivisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete division params
func (o *DeleteDivisionParams) WithContext(ctx context.Context) *DeleteDivisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete division params
func (o *DeleteDivisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete division params
func (o *DeleteDivisionParams) WithHTTPClient(client *http.Client) *DeleteDivisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete division params
func (o *DeleteDivisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete division params
func (o *DeleteDivisionParams) WithID(id int64) *DeleteDivisionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete division params
func (o *DeleteDivisionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDivisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
