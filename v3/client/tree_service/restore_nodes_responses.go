// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// RestoreNodesReader is a Reader for the RestoreNodes structure.
type RestoreNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestoreNodesOK creates a RestoreNodesOK with default headers values
func NewRestoreNodesOK() *RestoreNodesOK {
	return &RestoreNodesOK{}
}

/* RestoreNodesOK describes a response with status code 200, with default header values.

RestoreNodesOK restore nodes o k
*/
type RestoreNodesOK struct {
	Payload *models.RestRestoreNodesResponse
}

func (o *RestoreNodesOK) Error() string {
	return fmt.Sprintf("[POST /tree/restore][%d] restoreNodesOK  %+v", 200, o.Payload)
}
func (o *RestoreNodesOK) GetPayload() *models.RestRestoreNodesResponse {
	return o.Payload
}

func (o *RestoreNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestRestoreNodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}