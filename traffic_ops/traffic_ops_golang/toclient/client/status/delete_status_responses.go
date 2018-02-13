// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/toclient/models"
)

// DeleteStatusReader is a Reader for the DeleteStatus structure.
type DeleteStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStatusOK creates a DeleteStatusOK with default headers values
func NewDeleteStatusOK() *DeleteStatusOK {
	return &DeleteStatusOK{}
}

/*DeleteStatusOK handles this case with default header values.

Alerts - informs the client of server side messages
*/
type DeleteStatusOK struct {
	Payload models.DeleteStatusOKBody
}

func (o *DeleteStatusOK) Error() string {
	return fmt.Sprintf("[DELETE /statuses/{id}][%d] deleteStatusOK  %+v", 200, o.Payload)
}

func (o *DeleteStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
