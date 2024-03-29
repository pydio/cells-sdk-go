// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
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
		return nil, runtime.NewAPIError("[POST /activity/stream] Stream", response, response.Code())
	}
}

// NewStreamOK creates a StreamOK with default headers values
func NewStreamOK() *StreamOK {
	return &StreamOK{}
}

/*
StreamOK describes a response with status code 200, with default header values.

A successful response.
*/
type StreamOK struct {
	Payload *models.ActivityObject
}

// IsSuccess returns true when this stream o k response has a 2xx status code
func (o *StreamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stream o k response has a 3xx status code
func (o *StreamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream o k response has a 4xx status code
func (o *StreamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stream o k response has a 5xx status code
func (o *StreamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stream o k response a status code equal to that given
func (o *StreamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stream o k response
func (o *StreamOK) Code() int {
	return 200
}

func (o *StreamOK) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamOK  %+v", 200, o.Payload)
}

func (o *StreamOK) String() string {
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

/*
StreamUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type StreamUnauthorized struct {
}

// IsSuccess returns true when this stream unauthorized response has a 2xx status code
func (o *StreamUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream unauthorized response has a 3xx status code
func (o *StreamUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream unauthorized response has a 4xx status code
func (o *StreamUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream unauthorized response has a 5xx status code
func (o *StreamUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stream unauthorized response a status code equal to that given
func (o *StreamUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stream unauthorized response
func (o *StreamUnauthorized) Code() int {
	return 401
}

func (o *StreamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamUnauthorized ", 401)
}

func (o *StreamUnauthorized) String() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamUnauthorized ", 401)
}

func (o *StreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStreamForbidden creates a StreamForbidden with default headers values
func NewStreamForbidden() *StreamForbidden {
	return &StreamForbidden{}
}

/*
StreamForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type StreamForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this stream forbidden response has a 2xx status code
func (o *StreamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream forbidden response has a 3xx status code
func (o *StreamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream forbidden response has a 4xx status code
func (o *StreamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream forbidden response has a 5xx status code
func (o *StreamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stream forbidden response a status code equal to that given
func (o *StreamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stream forbidden response
func (o *StreamForbidden) Code() int {
	return 403
}

func (o *StreamForbidden) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamForbidden  %+v", 403, o.Payload)
}

func (o *StreamForbidden) String() string {
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

/*
StreamNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type StreamNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this stream not found response has a 2xx status code
func (o *StreamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream not found response has a 3xx status code
func (o *StreamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream not found response has a 4xx status code
func (o *StreamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream not found response has a 5xx status code
func (o *StreamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stream not found response a status code equal to that given
func (o *StreamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stream not found response
func (o *StreamNotFound) Code() int {
	return 404
}

func (o *StreamNotFound) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamNotFound  %+v", 404, o.Payload)
}

func (o *StreamNotFound) String() string {
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

/*
StreamInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type StreamInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this stream internal server error response has a 2xx status code
func (o *StreamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream internal server error response has a 3xx status code
func (o *StreamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream internal server error response has a 4xx status code
func (o *StreamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stream internal server error response has a 5xx status code
func (o *StreamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stream internal server error response a status code equal to that given
func (o *StreamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stream internal server error response
func (o *StreamInternalServerError) Code() int {
	return 500
}

func (o *StreamInternalServerError) Error() string {
	return fmt.Sprintf("[POST /activity/stream][%d] streamInternalServerError  %+v", 500, o.Payload)
}

func (o *StreamInternalServerError) String() string {
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
