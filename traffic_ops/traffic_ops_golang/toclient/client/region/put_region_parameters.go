// Code generated by go-swagger; DO NOT EDIT.

package region

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

	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/toclient/models"
)

// NewPutRegionParams creates a new PutRegionParams object
// with the default values initialized.
func NewPutRegionParams() *PutRegionParams {
	var ()
	return &PutRegionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRegionParamsWithTimeout creates a new PutRegionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRegionParamsWithTimeout(timeout time.Duration) *PutRegionParams {
	var ()
	return &PutRegionParams{

		timeout: timeout,
	}
}

// NewPutRegionParamsWithContext creates a new PutRegionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRegionParamsWithContext(ctx context.Context) *PutRegionParams {
	var ()
	return &PutRegionParams{

		Context: ctx,
	}
}

// NewPutRegionParamsWithHTTPClient creates a new PutRegionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRegionParamsWithHTTPClient(client *http.Client) *PutRegionParams {
	var ()
	return &PutRegionParams{
		HTTPClient: client,
	}
}

/*PutRegionParams contains all the parameters to send to the API endpoint
for the put region operation typically these are written to a http.Request
*/
type PutRegionParams struct {

	/*Region
	  Region Request Body

	*/
	Region *models.Region
	/*ID
	  ID

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put region params
func (o *PutRegionParams) WithTimeout(timeout time.Duration) *PutRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put region params
func (o *PutRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put region params
func (o *PutRegionParams) WithContext(ctx context.Context) *PutRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put region params
func (o *PutRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put region params
func (o *PutRegionParams) WithHTTPClient(client *http.Client) *PutRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put region params
func (o *PutRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the put region params
func (o *PutRegionParams) WithRegion(region *models.Region) *PutRegionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the put region params
func (o *PutRegionParams) SetRegion(region *models.Region) {
	o.Region = region
}

// WithID adds the id to the put region params
func (o *PutRegionParams) WithID(id int64) *PutRegionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put region params
func (o *PutRegionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Region != nil {
		if err := r.SetBodyParam(o.Region); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
