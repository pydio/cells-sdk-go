// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// PutShareLinkReader is a Reader for the PutShareLink structure.
type PutShareLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutShareLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutShareLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutShareLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutShareLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutShareLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutShareLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /share/link] PutShareLink", response, response.Code())
	}
}

// NewPutShareLinkOK creates a PutShareLinkOK with default headers values
func NewPutShareLinkOK() *PutShareLinkOK {
	return &PutShareLinkOK{}
}

/*
PutShareLinkOK describes a response with status code 200, with default header values.

A successful response.
*/
type PutShareLinkOK struct {
	Payload *models.RestShareLink
}

// IsSuccess returns true when this put share link o k response has a 2xx status code
func (o *PutShareLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put share link o k response has a 3xx status code
func (o *PutShareLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put share link o k response has a 4xx status code
func (o *PutShareLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put share link o k response has a 5xx status code
func (o *PutShareLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put share link o k response a status code equal to that given
func (o *PutShareLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put share link o k response
func (o *PutShareLinkOK) Code() int {
	return 200
}

func (o *PutShareLinkOK) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkOK  %+v", 200, o.Payload)
}

func (o *PutShareLinkOK) String() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkOK  %+v", 200, o.Payload)
}

func (o *PutShareLinkOK) GetPayload() *models.RestShareLink {
	return o.Payload
}

func (o *PutShareLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestShareLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkUnauthorized creates a PutShareLinkUnauthorized with default headers values
func NewPutShareLinkUnauthorized() *PutShareLinkUnauthorized {
	return &PutShareLinkUnauthorized{}
}

/*
PutShareLinkUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutShareLinkUnauthorized struct {
}

// IsSuccess returns true when this put share link unauthorized response has a 2xx status code
func (o *PutShareLinkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put share link unauthorized response has a 3xx status code
func (o *PutShareLinkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put share link unauthorized response has a 4xx status code
func (o *PutShareLinkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put share link unauthorized response has a 5xx status code
func (o *PutShareLinkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put share link unauthorized response a status code equal to that given
func (o *PutShareLinkUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put share link unauthorized response
func (o *PutShareLinkUnauthorized) Code() int {
	return 401
}

func (o *PutShareLinkUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkUnauthorized ", 401)
}

func (o *PutShareLinkUnauthorized) String() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkUnauthorized ", 401)
}

func (o *PutShareLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutShareLinkForbidden creates a PutShareLinkForbidden with default headers values
func NewPutShareLinkForbidden() *PutShareLinkForbidden {
	return &PutShareLinkForbidden{}
}

/*
PutShareLinkForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutShareLinkForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put share link forbidden response has a 2xx status code
func (o *PutShareLinkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put share link forbidden response has a 3xx status code
func (o *PutShareLinkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put share link forbidden response has a 4xx status code
func (o *PutShareLinkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put share link forbidden response has a 5xx status code
func (o *PutShareLinkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put share link forbidden response a status code equal to that given
func (o *PutShareLinkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put share link forbidden response
func (o *PutShareLinkForbidden) Code() int {
	return 403
}

func (o *PutShareLinkForbidden) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkForbidden  %+v", 403, o.Payload)
}

func (o *PutShareLinkForbidden) String() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkForbidden  %+v", 403, o.Payload)
}

func (o *PutShareLinkForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkNotFound creates a PutShareLinkNotFound with default headers values
func NewPutShareLinkNotFound() *PutShareLinkNotFound {
	return &PutShareLinkNotFound{}
}

/*
PutShareLinkNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutShareLinkNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put share link not found response has a 2xx status code
func (o *PutShareLinkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put share link not found response has a 3xx status code
func (o *PutShareLinkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put share link not found response has a 4xx status code
func (o *PutShareLinkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put share link not found response has a 5xx status code
func (o *PutShareLinkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put share link not found response a status code equal to that given
func (o *PutShareLinkNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put share link not found response
func (o *PutShareLinkNotFound) Code() int {
	return 404
}

func (o *PutShareLinkNotFound) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkNotFound  %+v", 404, o.Payload)
}

func (o *PutShareLinkNotFound) String() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkNotFound  %+v", 404, o.Payload)
}

func (o *PutShareLinkNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkInternalServerError creates a PutShareLinkInternalServerError with default headers values
func NewPutShareLinkInternalServerError() *PutShareLinkInternalServerError {
	return &PutShareLinkInternalServerError{}
}

/*
PutShareLinkInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutShareLinkInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put share link internal server error response has a 2xx status code
func (o *PutShareLinkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put share link internal server error response has a 3xx status code
func (o *PutShareLinkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put share link internal server error response has a 4xx status code
func (o *PutShareLinkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put share link internal server error response has a 5xx status code
func (o *PutShareLinkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put share link internal server error response a status code equal to that given
func (o *PutShareLinkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put share link internal server error response
func (o *PutShareLinkInternalServerError) Code() int {
	return 500
}

func (o *PutShareLinkInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *PutShareLinkInternalServerError) String() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *PutShareLinkInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
