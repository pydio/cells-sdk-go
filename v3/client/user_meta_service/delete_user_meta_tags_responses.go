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

// DeleteUserMetaTagsReader is a Reader for the DeleteUserMetaTags structure.
type DeleteUserMetaTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserMetaTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserMetaTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserMetaTagsOK creates a DeleteUserMetaTagsOK with default headers values
func NewDeleteUserMetaTagsOK() *DeleteUserMetaTagsOK {
	return &DeleteUserMetaTagsOK{}
}

/* DeleteUserMetaTagsOK describes a response with status code 200, with default header values.

DeleteUserMetaTagsOK delete user meta tags o k
*/
type DeleteUserMetaTagsOK struct {
	Payload *models.RestDeleteUserMetaTagsResponse
}

func (o *DeleteUserMetaTagsOK) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsOK  %+v", 200, o.Payload)
}
func (o *DeleteUserMetaTagsOK) GetPayload() *models.RestDeleteUserMetaTagsResponse {
	return o.Payload
}

func (o *DeleteUserMetaTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteUserMetaTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
