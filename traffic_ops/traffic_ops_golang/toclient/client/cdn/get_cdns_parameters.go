// Code generated by go-swagger; DO NOT EDIT.

package cdn

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
)

// NewGetCDNsParams creates a new GetCDNsParams object
// with the default values initialized.
func NewGetCDNsParams() *GetCDNsParams {
	var ()
	return &GetCDNsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCDNsParamsWithTimeout creates a new GetCDNsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCDNsParamsWithTimeout(timeout time.Duration) *GetCDNsParams {
	var ()
	return &GetCDNsParams{

		timeout: timeout,
	}
}

// NewGetCDNsParamsWithContext creates a new GetCDNsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCDNsParamsWithContext(ctx context.Context) *GetCDNsParams {
	var ()
	return &GetCDNsParams{

		Context: ctx,
	}
}

// NewGetCDNsParamsWithHTTPClient creates a new GetCDNsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCDNsParamsWithHTTPClient(client *http.Client) *GetCDNsParams {
	var ()
	return &GetCDNsParams{
		HTTPClient: client,
	}
}

/*GetCDNsParams contains all the parameters to send to the API endpoint
for the get c d ns operation typically these are written to a http.Request
*/
type GetCDNsParams struct {

	/*DnssecEnabled
	  Enables Domain Name System Security Extensions (DNSSEC) for the CDN

	*/
	DnssecEnabled *string
	/*DomainName
	  The domain name for the CDN

	*/
	DomainName *string
	/*ID
	  Unique identifier for the CDN

	*/
	ID *string
	/*Name
	  The CDN name for the CDN

	*/
	Name *string
	/*Orderby*/
	Orderby *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get c d ns params
func (o *GetCDNsParams) WithTimeout(timeout time.Duration) *GetCDNsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c d ns params
func (o *GetCDNsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c d ns params
func (o *GetCDNsParams) WithContext(ctx context.Context) *GetCDNsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c d ns params
func (o *GetCDNsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c d ns params
func (o *GetCDNsParams) WithHTTPClient(client *http.Client) *GetCDNsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c d ns params
func (o *GetCDNsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDnssecEnabled adds the dnssecEnabled to the get c d ns params
func (o *GetCDNsParams) WithDnssecEnabled(dnssecEnabled *string) *GetCDNsParams {
	o.SetDnssecEnabled(dnssecEnabled)
	return o
}

// SetDnssecEnabled adds the dnssecEnabled to the get c d ns params
func (o *GetCDNsParams) SetDnssecEnabled(dnssecEnabled *string) {
	o.DnssecEnabled = dnssecEnabled
}

// WithDomainName adds the domainName to the get c d ns params
func (o *GetCDNsParams) WithDomainName(domainName *string) *GetCDNsParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the get c d ns params
func (o *GetCDNsParams) SetDomainName(domainName *string) {
	o.DomainName = domainName
}

// WithID adds the id to the get c d ns params
func (o *GetCDNsParams) WithID(id *string) *GetCDNsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get c d ns params
func (o *GetCDNsParams) SetID(id *string) {
	o.ID = id
}

// WithName adds the name to the get c d ns params
func (o *GetCDNsParams) WithName(name *string) *GetCDNsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get c d ns params
func (o *GetCDNsParams) SetName(name *string) {
	o.Name = name
}

// WithOrderby adds the orderby to the get c d ns params
func (o *GetCDNsParams) WithOrderby(orderby *string) *GetCDNsParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the get c d ns params
func (o *GetCDNsParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WriteToRequest writes these params to a swagger request
func (o *GetCDNsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DnssecEnabled != nil {

		// query param dnssecEnabled
		var qrDnssecEnabled string
		if o.DnssecEnabled != nil {
			qrDnssecEnabled = *o.DnssecEnabled
		}
		qDnssecEnabled := qrDnssecEnabled
		if qDnssecEnabled != "" {
			if err := r.SetQueryParam("dnssecEnabled", qDnssecEnabled); err != nil {
				return err
			}
		}

	}

	if o.DomainName != nil {

		// query param domainName
		var qrDomainName string
		if o.DomainName != nil {
			qrDomainName = *o.DomainName
		}
		qDomainName := qrDomainName
		if qDomainName != "" {
			if err := r.SetQueryParam("domainName", qDomainName); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
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

	if o.Orderby != nil {

		// query param orderby
		var qrOrderby string
		if o.Orderby != nil {
			qrOrderby = *o.Orderby
		}
		qOrderby := qrOrderby
		if qOrderby != "" {
			if err := r.SetQueryParam("orderby", qOrderby); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
