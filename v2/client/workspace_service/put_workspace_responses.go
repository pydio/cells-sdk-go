// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// PutWorkspaceReader is a Reader for the PutWorkspace structure.
type PutWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutWorkspaceOK creates a PutWorkspaceOK with default headers values
func NewPutWorkspaceOK() *PutWorkspaceOK {
	return &PutWorkspaceOK{}
}

/*PutWorkspaceOK handles this case with default header values.

PutWorkspaceOK put workspace o k
*/
type PutWorkspaceOK struct {
	Payload *models.IdmWorkspace
}

func (o *PutWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmWorkspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
