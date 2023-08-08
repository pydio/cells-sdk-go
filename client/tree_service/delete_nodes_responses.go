// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// DeleteNodesReader is a Reader for the DeleteNodes structure.
type DeleteNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /tree/delete] DeleteNodes", response, response.Code())
	}
}

// NewDeleteNodesOK creates a DeleteNodesOK with default headers values
func NewDeleteNodesOK() *DeleteNodesOK {
	return &DeleteNodesOK{}
}

/*
DeleteNodesOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteNodesOK struct {
	Payload *models.RestDeleteNodesResponse
}

// IsSuccess returns true when this delete nodes o k response has a 2xx status code
func (o *DeleteNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete nodes o k response has a 3xx status code
func (o *DeleteNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete nodes o k response has a 4xx status code
func (o *DeleteNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete nodes o k response has a 5xx status code
func (o *DeleteNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete nodes o k response a status code equal to that given
func (o *DeleteNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete nodes o k response
func (o *DeleteNodesOK) Code() int {
	return 200
}

func (o *DeleteNodesOK) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesOK  %+v", 200, o.Payload)
}

func (o *DeleteNodesOK) String() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesOK  %+v", 200, o.Payload)
}

func (o *DeleteNodesOK) GetPayload() *models.RestDeleteNodesResponse {
	return o.Payload
}

func (o *DeleteNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteNodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodesUnauthorized creates a DeleteNodesUnauthorized with default headers values
func NewDeleteNodesUnauthorized() *DeleteNodesUnauthorized {
	return &DeleteNodesUnauthorized{}
}

/*
DeleteNodesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteNodesUnauthorized struct {
}

// IsSuccess returns true when this delete nodes unauthorized response has a 2xx status code
func (o *DeleteNodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete nodes unauthorized response has a 3xx status code
func (o *DeleteNodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete nodes unauthorized response has a 4xx status code
func (o *DeleteNodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete nodes unauthorized response has a 5xx status code
func (o *DeleteNodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete nodes unauthorized response a status code equal to that given
func (o *DeleteNodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete nodes unauthorized response
func (o *DeleteNodesUnauthorized) Code() int {
	return 401
}

func (o *DeleteNodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesUnauthorized ", 401)
}

func (o *DeleteNodesUnauthorized) String() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesUnauthorized ", 401)
}

func (o *DeleteNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodesForbidden creates a DeleteNodesForbidden with default headers values
func NewDeleteNodesForbidden() *DeleteNodesForbidden {
	return &DeleteNodesForbidden{}
}

/*
DeleteNodesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteNodesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete nodes forbidden response has a 2xx status code
func (o *DeleteNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete nodes forbidden response has a 3xx status code
func (o *DeleteNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete nodes forbidden response has a 4xx status code
func (o *DeleteNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete nodes forbidden response has a 5xx status code
func (o *DeleteNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete nodes forbidden response a status code equal to that given
func (o *DeleteNodesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete nodes forbidden response
func (o *DeleteNodesForbidden) Code() int {
	return 403
}

func (o *DeleteNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNodesForbidden) String() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNodesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodesNotFound creates a DeleteNodesNotFound with default headers values
func NewDeleteNodesNotFound() *DeleteNodesNotFound {
	return &DeleteNodesNotFound{}
}

/*
DeleteNodesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteNodesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete nodes not found response has a 2xx status code
func (o *DeleteNodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete nodes not found response has a 3xx status code
func (o *DeleteNodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete nodes not found response has a 4xx status code
func (o *DeleteNodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete nodes not found response has a 5xx status code
func (o *DeleteNodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete nodes not found response a status code equal to that given
func (o *DeleteNodesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete nodes not found response
func (o *DeleteNodesNotFound) Code() int {
	return 404
}

func (o *DeleteNodesNotFound) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodesNotFound) String() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteNodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodesInternalServerError creates a DeleteNodesInternalServerError with default headers values
func NewDeleteNodesInternalServerError() *DeleteNodesInternalServerError {
	return &DeleteNodesInternalServerError{}
}

/*
DeleteNodesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteNodesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete nodes internal server error response has a 2xx status code
func (o *DeleteNodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete nodes internal server error response has a 3xx status code
func (o *DeleteNodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete nodes internal server error response has a 4xx status code
func (o *DeleteNodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete nodes internal server error response has a 5xx status code
func (o *DeleteNodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete nodes internal server error response a status code equal to that given
func (o *DeleteNodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete nodes internal server error response
func (o *DeleteNodesInternalServerError) Code() int {
	return 500
}

func (o *DeleteNodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNodesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}