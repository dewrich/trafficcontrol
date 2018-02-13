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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/toclient/models"
)

// NewPostASNParams creates a new PostASNParams object
// with the default values initialized.
func NewPostASNParams() *PostASNParams {
	var ()
	return &PostASNParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostASNParamsWithTimeout creates a new PostASNParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostASNParamsWithTimeout(timeout time.Duration) *PostASNParams {
	var ()
	return &PostASNParams{

		timeout: timeout,
	}
}

// NewPostASNParamsWithContext creates a new PostASNParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostASNParamsWithContext(ctx context.Context) *PostASNParams {
	var ()
	return &PostASNParams{

		Context: ctx,
	}
}

// NewPostASNParamsWithHTTPClient creates a new PostASNParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostASNParamsWithHTTPClient(client *http.Client) *PostASNParams {
	var ()
	return &PostASNParams{
		HTTPClient: client,
	}
}

/*PostASNParams contains all the parameters to send to the API endpoint
for the post a s n operation typically these are written to a http.Request
*/
type PostASNParams struct {

	/*ASN
	  ASN Request Body

	*/
	ASN *models.ASN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post a s n params
func (o *PostASNParams) WithTimeout(timeout time.Duration) *PostASNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post a s n params
func (o *PostASNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post a s n params
func (o *PostASNParams) WithContext(ctx context.Context) *PostASNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post a s n params
func (o *PostASNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post a s n params
func (o *PostASNParams) WithHTTPClient(client *http.Client) *PostASNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post a s n params
func (o *PostASNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithASN adds the aSN to the post a s n params
func (o *PostASNParams) WithASN(aSN *models.ASN) *PostASNParams {
	o.SetASN(aSN)
	return o
}

// SetASN adds the aSN to the post a s n params
func (o *PostASNParams) SetASN(aSN *models.ASN) {
	o.ASN = aSN
}

// WriteToRequest writes these params to a swagger request
func (o *PostASNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ASN != nil {
		if err := r.SetBodyParam(o.ASN); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
