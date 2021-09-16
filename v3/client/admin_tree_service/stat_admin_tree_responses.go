// Code generated by go-swagger; DO NOT EDIT.

package admin_tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// StatAdminTreeReader is a Reader for the StatAdminTree structure.
type StatAdminTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatAdminTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatAdminTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStatAdminTreeOK creates a StatAdminTreeOK with default headers values
func NewStatAdminTreeOK() *StatAdminTreeOK {
	return &StatAdminTreeOK{}
}

/* StatAdminTreeOK describes a response with status code 200, with default header values.

StatAdminTreeOK stat admin tree o k
*/
type StatAdminTreeOK struct {
	Payload *models.TreeReadNodeResponse
}

func (o *StatAdminTreeOK) Error() string {
	return fmt.Sprintf("[POST /tree/admin/stat][%d] statAdminTreeOK  %+v", 200, o.Payload)
}
func (o *StatAdminTreeOK) GetPayload() *models.TreeReadNodeResponse {
	return o.Payload
}

func (o *StatAdminTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreeReadNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
