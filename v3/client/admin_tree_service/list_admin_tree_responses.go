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

// ListAdminTreeReader is a Reader for the ListAdminTree structure.
type ListAdminTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAdminTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAdminTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAdminTreeOK creates a ListAdminTreeOK with default headers values
func NewListAdminTreeOK() *ListAdminTreeOK {
	return &ListAdminTreeOK{}
}

/* ListAdminTreeOK describes a response with status code 200, with default header values.

ListAdminTreeOK list admin tree o k
*/
type ListAdminTreeOK struct {
	Payload *models.RestNodesCollection
}

func (o *ListAdminTreeOK) Error() string {
	return fmt.Sprintf("[POST /tree/admin/list][%d] listAdminTreeOK  %+v", 200, o.Payload)
}
func (o *ListAdminTreeOK) GetPayload() *models.RestNodesCollection {
	return o.Payload
}

func (o *ListAdminTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestNodesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
