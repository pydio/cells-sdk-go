// Code generated by go-swagger; DO NOT EDIT.

package acl_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PutACLReader is a Reader for the PutACL structure.
type PutACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutACLOK creates a PutACLOK with default headers values
func NewPutACLOK() *PutACLOK {
	return &PutACLOK{}
}

/* PutACLOK describes a response with status code 200, with default header values.

PutACLOK put Acl o k
*/
type PutACLOK struct {
	Payload *models.IdmACL
}

func (o *PutACLOK) Error() string {
	return fmt.Sprintf("[PUT /acl][%d] putAclOK  %+v", 200, o.Payload)
}
func (o *PutACLOK) GetPayload() *models.IdmACL {
	return o.Payload
}

func (o *PutACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
