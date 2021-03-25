// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// CreateSelectionReader is a Reader for the CreateSelection structure.
type CreateSelectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSelectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSelectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSelectionOK creates a CreateSelectionOK with default headers values
func NewCreateSelectionOK() *CreateSelectionOK {
	return &CreateSelectionOK{}
}

/*CreateSelectionOK handles this case with default header values.

CreateSelectionOK create selection o k
*/
type CreateSelectionOK struct {
	Payload *models.RestCreateSelectionResponse
}

func (o *CreateSelectionOK) Error() string {
	return fmt.Sprintf("[POST /tree/selection][%d] createSelectionOK  %+v", 200, o.Payload)
}

func (o *CreateSelectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCreateSelectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
