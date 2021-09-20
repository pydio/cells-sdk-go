// Code generated by go-swagger; DO NOT EDIT.

package search_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// NodesReader is a Reader for the Nodes structure.
type NodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodesOK creates a NodesOK with default headers values
func NewNodesOK() *NodesOK {
	return &NodesOK{}
}

/* NodesOK describes a response with status code 200, with default header values.

NodesOK nodes o k
*/
type NodesOK struct {
	Payload *models.RestSearchResults
}

func (o *NodesOK) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesOK  %+v", 200, o.Payload)
}
func (o *NodesOK) GetPayload() *models.RestSearchResults {
	return o.Payload
}

func (o *NodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}