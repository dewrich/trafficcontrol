// Code generated by go-swagger; DO NOT EDIT.

package asn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/toclient/models"
)

// PutASNReader is a Reader for the PutASN structure.
type PutASNReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutASNReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutASNOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutASNOK creates a PutASNOK with default headers values
func NewPutASNOK() *PutASNOK {
	return &PutASNOK{}
}

/*PutASNOK handles this case with default header values.

ASN -  ASNResponse to get the "response" top level key
*/
type PutASNOK struct {
	Payload *models.ASNResponse
}

func (o *PutASNOK) Error() string {
	return fmt.Sprintf("[PUT /asns/{id}][%d] putASNOK  %+v", 200, o.Payload)
}

func (o *PutASNOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASNResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
