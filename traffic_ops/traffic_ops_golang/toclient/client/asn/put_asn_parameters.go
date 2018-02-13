// Code generated by go-swagger; DO NOT EDIT.

package asn

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

// NewPutASNParams creates a new PutASNParams object
// with the default values initialized.
func NewPutASNParams() *PutASNParams {
	var ()
	return &PutASNParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutASNParamsWithTimeout creates a new PutASNParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutASNParamsWithTimeout(timeout time.Duration) *PutASNParams {
	var ()
	return &PutASNParams{

		timeout: timeout,
	}
}

// NewPutASNParamsWithContext creates a new PutASNParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutASNParamsWithContext(ctx context.Context) *PutASNParams {
	var ()
	return &PutASNParams{

		Context: ctx,
	}
}

// NewPutASNParamsWithHTTPClient creates a new PutASNParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutASNParamsWithHTTPClient(client *http.Client) *PutASNParams {
	var ()
	return &PutASNParams{
		HTTPClient: client,
	}
}

/*PutASNParams contains all the parameters to send to the API endpoint
for the put a s n operation typically these are written to a http.Request
*/
type PutASNParams struct {

	/*ASN
	  ASN Request Body

	*/
	ASN *models.ASN
	/*ID
	  ID

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put a s n params
func (o *PutASNParams) WithTimeout(timeout time.Duration) *PutASNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put a s n params
func (o *PutASNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put a s n params
func (o *PutASNParams) WithContext(ctx context.Context) *PutASNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put a s n params
func (o *PutASNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put a s n params
func (o *PutASNParams) WithHTTPClient(client *http.Client) *PutASNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put a s n params
func (o *PutASNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithASN adds the aSN to the put a s n params
func (o *PutASNParams) WithASN(aSN *models.ASN) *PutASNParams {
	o.SetASN(aSN)
	return o
}

// SetASN adds the aSN to the put a s n params
func (o *PutASNParams) SetASN(aSN *models.ASN) {
	o.ASN = aSN
}

// WithID adds the id to the put a s n params
func (o *PutASNParams) WithID(id int64) *PutASNParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put a s n params
func (o *PutASNParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutASNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ASN != nil {
		if err := r.SetBodyParam(o.ASN); err != nil {
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
