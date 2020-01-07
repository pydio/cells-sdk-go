// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// ListStorageBucketsReader is a Reader for the ListStorageBuckets structure.
type ListStorageBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStorageBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListStorageBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListStorageBucketsOK creates a ListStorageBucketsOK with default headers values
func NewListStorageBucketsOK() *ListStorageBucketsOK {
	return &ListStorageBucketsOK{}
}

/*ListStorageBucketsOK handles this case with default header values.

ListStorageBucketsOK list storage buckets o k
*/
type ListStorageBucketsOK struct {
	Payload *models.RestNodesCollection
}

func (o *ListStorageBucketsOK) Error() string {
	return fmt.Sprintf("[POST /config/buckets][%d] listStorageBucketsOK  %+v", 200, o.Payload)
}

func (o *ListStorageBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestNodesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
