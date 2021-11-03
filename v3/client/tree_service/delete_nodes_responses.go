// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNodesOK creates a DeleteNodesOK with default headers values
func NewDeleteNodesOK() *DeleteNodesOK {
	return &DeleteNodesOK{}
}

/* DeleteNodesOK describes a response with status code 200, with default header values.

DeleteNodesOK delete nodes o k
*/
type DeleteNodesOK struct {
	Payload *models.RestDeleteNodesResponse
}

func (o *DeleteNodesOK) Error() string {
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

/* DeleteNodesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteNodesUnauthorized struct {
}

func (o *DeleteNodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tree/delete][%d] deleteNodesUnauthorized ", 401)
}

func (o *DeleteNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodesForbidden creates a DeleteNodesForbidden with default headers values
func NewDeleteNodesForbidden() *DeleteNodesForbidden {
	return &DeleteNodesForbidden{}
}

/* DeleteNodesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteNodesForbidden struct {
	Payload *models.RestError
}

func (o *DeleteNodesForbidden) Error() string {
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

/* DeleteNodesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteNodesNotFound struct {
	Payload *models.RestError
}

func (o *DeleteNodesNotFound) Error() string {
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

/* DeleteNodesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteNodesInternalServerError struct {
	Payload *models.RestError
}

func (o *DeleteNodesInternalServerError) Error() string {
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
