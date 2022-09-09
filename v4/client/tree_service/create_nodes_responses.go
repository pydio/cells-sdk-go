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

// CreateNodesReader is a Reader for the CreateNodes structure.
type CreateNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateNodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNodesOK creates a CreateNodesOK with default headers values
func NewCreateNodesOK() *CreateNodesOK {
	return &CreateNodesOK{}
}

/*
CreateNodesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateNodesOK struct {
	Payload *models.RestNodesCollection
}

// IsSuccess returns true when this create nodes o k response has a 2xx status code
func (o *CreateNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create nodes o k response has a 3xx status code
func (o *CreateNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nodes o k response has a 4xx status code
func (o *CreateNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create nodes o k response has a 5xx status code
func (o *CreateNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create nodes o k response a status code equal to that given
func (o *CreateNodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateNodesOK) Error() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesOK  %+v", 200, o.Payload)
}

func (o *CreateNodesOK) String() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesOK  %+v", 200, o.Payload)
}

func (o *CreateNodesOK) GetPayload() *models.RestNodesCollection {
	return o.Payload
}

func (o *CreateNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestNodesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodesUnauthorized creates a CreateNodesUnauthorized with default headers values
func NewCreateNodesUnauthorized() *CreateNodesUnauthorized {
	return &CreateNodesUnauthorized{}
}

/*
CreateNodesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type CreateNodesUnauthorized struct {
}

// IsSuccess returns true when this create nodes unauthorized response has a 2xx status code
func (o *CreateNodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nodes unauthorized response has a 3xx status code
func (o *CreateNodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nodes unauthorized response has a 4xx status code
func (o *CreateNodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create nodes unauthorized response has a 5xx status code
func (o *CreateNodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create nodes unauthorized response a status code equal to that given
func (o *CreateNodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateNodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesUnauthorized ", 401)
}

func (o *CreateNodesUnauthorized) String() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesUnauthorized ", 401)
}

func (o *CreateNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodesForbidden creates a CreateNodesForbidden with default headers values
func NewCreateNodesForbidden() *CreateNodesForbidden {
	return &CreateNodesForbidden{}
}

/*
CreateNodesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type CreateNodesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create nodes forbidden response has a 2xx status code
func (o *CreateNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nodes forbidden response has a 3xx status code
func (o *CreateNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nodes forbidden response has a 4xx status code
func (o *CreateNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create nodes forbidden response has a 5xx status code
func (o *CreateNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create nodes forbidden response a status code equal to that given
func (o *CreateNodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesForbidden  %+v", 403, o.Payload)
}

func (o *CreateNodesForbidden) String() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesForbidden  %+v", 403, o.Payload)
}

func (o *CreateNodesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodesNotFound creates a CreateNodesNotFound with default headers values
func NewCreateNodesNotFound() *CreateNodesNotFound {
	return &CreateNodesNotFound{}
}

/*
CreateNodesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type CreateNodesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create nodes not found response has a 2xx status code
func (o *CreateNodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nodes not found response has a 3xx status code
func (o *CreateNodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nodes not found response has a 4xx status code
func (o *CreateNodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create nodes not found response has a 5xx status code
func (o *CreateNodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create nodes not found response a status code equal to that given
func (o *CreateNodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateNodesNotFound) Error() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesNotFound  %+v", 404, o.Payload)
}

func (o *CreateNodesNotFound) String() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesNotFound  %+v", 404, o.Payload)
}

func (o *CreateNodesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateNodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodesInternalServerError creates a CreateNodesInternalServerError with default headers values
func NewCreateNodesInternalServerError() *CreateNodesInternalServerError {
	return &CreateNodesInternalServerError{}
}

/*
CreateNodesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type CreateNodesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create nodes internal server error response has a 2xx status code
func (o *CreateNodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nodes internal server error response has a 3xx status code
func (o *CreateNodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nodes internal server error response has a 4xx status code
func (o *CreateNodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create nodes internal server error response has a 5xx status code
func (o *CreateNodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create nodes internal server error response a status code equal to that given
func (o *CreateNodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateNodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /tree/create][%d] createNodesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNodesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
