// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PutShareLinkReader is a Reader for the PutShareLink structure.
type PutShareLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutShareLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutShareLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutShareLinkOK creates a PutShareLinkOK with default headers values
func NewPutShareLinkOK() *PutShareLinkOK {
	return &PutShareLinkOK{}
}

/* PutShareLinkOK describes a response with status code 200, with default header values.

PutShareLinkOK put share link o k
*/
type PutShareLinkOK struct {
	Payload *models.RestShareLink
}

func (o *PutShareLinkOK) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkOK  %+v", 200, o.Payload)
}
func (o *PutShareLinkOK) GetPayload() *models.RestShareLink {
	return o.Payload
}

func (o *PutShareLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestShareLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}