// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// StreamReader is a Reader for the Stream structure.
type StreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStreamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStreamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStreamOK creates a StreamOK with default headers values
func NewStreamOK() *StreamOK {
	return &StreamOK{}
}

/* StreamOK describes a response with status code 200, with default header values.

StreamOK stream o k
*/
type StreamOK struct {
	Payload *models.ActivityObject
}

func (o *StreamOK) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamOK  %+v", 200, o.Payload)
}
func (o *StreamOK) GetPayload() *models.ActivityObject {
	return o.Payload
}

func (o *StreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamUnauthorized creates a StreamUnauthorized with default headers values
func NewStreamUnauthorized() *StreamUnauthorized {
	return &StreamUnauthorized{}
}

/* StreamUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type StreamUnauthorized struct {
}

func (o *StreamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamUnauthorized ", 401)
}

func (o *StreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStreamForbidden creates a StreamForbidden with default headers values
func NewStreamForbidden() *StreamForbidden {
	return &StreamForbidden{}
}

/* StreamForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type StreamForbidden struct {
	Payload *models.RestError
}

func (o *StreamForbidden) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamForbidden  %+v", 403, o.Payload)
}
func (o *StreamForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *StreamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamNotFound creates a StreamNotFound with default headers values
func NewStreamNotFound() *StreamNotFound {
	return &StreamNotFound{}
}

/* StreamNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type StreamNotFound struct {
	Payload *models.RestError
}

func (o *StreamNotFound) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamNotFound  %+v", 404, o.Payload)
}
func (o *StreamNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *StreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamInternalServerError creates a StreamInternalServerError with default headers values
func NewStreamInternalServerError() *StreamInternalServerError {
	return &StreamInternalServerError{}
}

/* StreamInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type StreamInternalServerError struct {
	Payload *models.RestError
}

func (o *StreamInternalServerError) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamInternalServerError  %+v", 500, o.Payload)
}
func (o *StreamInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *StreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
