// Code generated by go-swagger; DO NOT EDIT.

package graph_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// RelationReader is a Reader for the Relation structure.
type RelationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRelationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRelationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRelationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRelationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelationOK creates a RelationOK with default headers values
func NewRelationOK() *RelationOK {
	return &RelationOK{}
}

/* RelationOK describes a response with status code 200, with default header values.

RelationOK relation o k
*/
type RelationOK struct {
	Payload *models.RestRelationResponse
}

func (o *RelationOK) Error() string {
	return fmt.Sprintf("[GET /graph/relation/{UserId}][%d] relationOK  %+v", 200, o.Payload)
}
func (o *RelationOK) GetPayload() *models.RestRelationResponse {
	return o.Payload
}

func (o *RelationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestRelationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationUnauthorized creates a RelationUnauthorized with default headers values
func NewRelationUnauthorized() *RelationUnauthorized {
	return &RelationUnauthorized{}
}

/* RelationUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type RelationUnauthorized struct {
	Payload *models.RestError
}

func (o *RelationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /graph/relation/{UserId}][%d] relationUnauthorized  %+v", 401, o.Payload)
}
func (o *RelationUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *RelationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationForbidden creates a RelationForbidden with default headers values
func NewRelationForbidden() *RelationForbidden {
	return &RelationForbidden{}
}

/* RelationForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type RelationForbidden struct {
	Payload *models.RestError
}

func (o *RelationForbidden) Error() string {
	return fmt.Sprintf("[GET /graph/relation/{UserId}][%d] relationForbidden  %+v", 403, o.Payload)
}
func (o *RelationForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *RelationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationNotFound creates a RelationNotFound with default headers values
func NewRelationNotFound() *RelationNotFound {
	return &RelationNotFound{}
}

/* RelationNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type RelationNotFound struct {
	Payload *models.RestError
}

func (o *RelationNotFound) Error() string {
	return fmt.Sprintf("[GET /graph/relation/{UserId}][%d] relationNotFound  %+v", 404, o.Payload)
}
func (o *RelationNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *RelationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationInternalServerError creates a RelationInternalServerError with default headers values
func NewRelationInternalServerError() *RelationInternalServerError {
	return &RelationInternalServerError{}
}

/* RelationInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type RelationInternalServerError struct {
	Payload *models.RestError
}

func (o *RelationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /graph/relation/{UserId}][%d] relationInternalServerError  %+v", 500, o.Payload)
}
func (o *RelationInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *RelationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
