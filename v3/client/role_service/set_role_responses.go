// Code generated by go-swagger; DO NOT EDIT.

package role_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// SetRoleReader is a Reader for the SetRole structure.
type SetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetRoleOK creates a SetRoleOK with default headers values
func NewSetRoleOK() *SetRoleOK {
	return &SetRoleOK{}
}

/* SetRoleOK describes a response with status code 200, with default header values.

SetRoleOK set role o k
*/
type SetRoleOK struct {
	Payload *models.IdmRole
}

func (o *SetRoleOK) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleOK  %+v", 200, o.Payload)
}
func (o *SetRoleOK) GetPayload() *models.IdmRole {
	return o.Payload
}

func (o *SetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
