// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// UpdateUserMetaReader is a Reader for the UpdateUserMeta structure.
type UpdateUserMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserMetaOK creates a UpdateUserMetaOK with default headers values
func NewUpdateUserMetaOK() *UpdateUserMetaOK {
	return &UpdateUserMetaOK{}
}

/* UpdateUserMetaOK describes a response with status code 200, with default header values.

UpdateUserMetaOK update user meta o k
*/
type UpdateUserMetaOK struct {
	Payload *models.IdmUpdateUserMetaResponse
}

func (o *UpdateUserMetaOK) Error() string {
	return fmt.Sprintf("[PUT /user-meta/update][%d] updateUserMetaOK  %+v", 200, o.Payload)
}
func (o *UpdateUserMetaOK) GetPayload() *models.IdmUpdateUserMetaResponse {
	return o.Payload
}

func (o *UpdateUserMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmUpdateUserMetaResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
