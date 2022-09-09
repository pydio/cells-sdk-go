// Code generated by go-swagger; DO NOT EDIT.

package search_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
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
	case 401:
		result := NewNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodesOK creates a NodesOK with default headers values
func NewNodesOK() *NodesOK {
	return &NodesOK{}
}

/*
NodesOK describes a response with status code 200, with default header values.

A successful response.
*/
type NodesOK struct {
	Payload *models.RestSearchResults
}

// IsSuccess returns true when this nodes o k response has a 2xx status code
func (o *NodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes o k response has a 3xx status code
func (o *NodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes o k response has a 4xx status code
func (o *NodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes o k response has a 5xx status code
func (o *NodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes o k response a status code equal to that given
func (o *NodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *NodesOK) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesOK  %+v", 200, o.Payload)
}

func (o *NodesOK) String() string {
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

// NewNodesUnauthorized creates a NodesUnauthorized with default headers values
func NewNodesUnauthorized() *NodesUnauthorized {
	return &NodesUnauthorized{}
}

/*
NodesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type NodesUnauthorized struct {
}

// IsSuccess returns true when this nodes unauthorized response has a 2xx status code
func (o *NodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes unauthorized response has a 3xx status code
func (o *NodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes unauthorized response has a 4xx status code
func (o *NodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes unauthorized response has a 5xx status code
func (o *NodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes unauthorized response a status code equal to that given
func (o *NodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesUnauthorized ", 401)
}

func (o *NodesUnauthorized) String() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesUnauthorized ", 401)
}

func (o *NodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNodesForbidden creates a NodesForbidden with default headers values
func NewNodesForbidden() *NodesForbidden {
	return &NodesForbidden{}
}

/*
NodesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type NodesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this nodes forbidden response has a 2xx status code
func (o *NodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes forbidden response has a 3xx status code
func (o *NodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes forbidden response has a 4xx status code
func (o *NodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes forbidden response has a 5xx status code
func (o *NodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes forbidden response a status code equal to that given
func (o *NodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NodesForbidden) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesForbidden  %+v", 403, o.Payload)
}

func (o *NodesForbidden) String() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesForbidden  %+v", 403, o.Payload)
}

func (o *NodesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *NodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesNotFound creates a NodesNotFound with default headers values
func NewNodesNotFound() *NodesNotFound {
	return &NodesNotFound{}
}

/*
NodesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type NodesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this nodes not found response has a 2xx status code
func (o *NodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes not found response has a 3xx status code
func (o *NodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes not found response has a 4xx status code
func (o *NodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes not found response has a 5xx status code
func (o *NodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes not found response a status code equal to that given
func (o *NodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NodesNotFound) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesNotFound  %+v", 404, o.Payload)
}

func (o *NodesNotFound) String() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesNotFound  %+v", 404, o.Payload)
}

func (o *NodesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *NodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesInternalServerError creates a NodesInternalServerError with default headers values
func NewNodesInternalServerError() *NodesInternalServerError {
	return &NodesInternalServerError{}
}

/*
NodesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type NodesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this nodes internal server error response has a 2xx status code
func (o *NodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes internal server error response has a 3xx status code
func (o *NodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes internal server error response has a 4xx status code
func (o *NodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes internal server error response has a 5xx status code
func (o *NodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this nodes internal server error response a status code equal to that given
func (o *NodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /search/nodes][%d] nodesInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *NodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
