// Code generated by go-swagger; DO NOT EDIT.

package meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// SetMetaReader is a Reader for the SetMeta structure.
type SetMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetMetaOK creates a SetMetaOK with default headers values
func NewSetMetaOK() *SetMetaOK {
	return &SetMetaOK{}
}

/*SetMetaOK handles this case with default header values.

SetMetaOK set meta o k
*/
type SetMetaOK struct {
	Payload *models.TreeNode
}

func (o *SetMetaOK) Error() string {
	return fmt.Sprintf("[POST /meta/set/{NodePath}][%d] setMetaOK  %+v", 200, o.Payload)
}

func (o *SetMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreeNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
