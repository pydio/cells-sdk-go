// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// UpdateUserMetaNamespaceReader is a Reader for the UpdateUserMetaNamespace structure.
type UpdateUserMetaNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserMetaNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserMetaNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserMetaNamespaceOK creates a UpdateUserMetaNamespaceOK with default headers values
func NewUpdateUserMetaNamespaceOK() *UpdateUserMetaNamespaceOK {
	return &UpdateUserMetaNamespaceOK{}
}

/*UpdateUserMetaNamespaceOK handles this case with default header values.

UpdateUserMetaNamespaceOK update user meta namespace o k
*/
type UpdateUserMetaNamespaceOK struct {
	Payload *models.IdmUpdateUserMetaNamespaceResponse
}

func (o *UpdateUserMetaNamespaceOK) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceOK  %+v", 200, o.Payload)
}

func (o *UpdateUserMetaNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmUpdateUserMetaNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
