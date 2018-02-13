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

// PostStatusReader is a Reader for the PostStatus structure.
type PostStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStatusOK creates a PostStatusOK with default headers values
func NewPostStatusOK() *PostStatusOK {
	return &PostStatusOK{}
}

/*PostStatusOK handles this case with default header values.

Alerts - informs the client of server side messages
*/
type PostStatusOK struct {
	Payload models.PostStatusOKBody
}

func (o *PostStatusOK) Error() string {
	return fmt.Sprintf("[POST /statuses][%d] postStatusOK  %+v", 200, o.Payload)
}

func (o *PostStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
