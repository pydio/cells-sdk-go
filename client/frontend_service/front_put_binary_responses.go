// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// FrontPutBinaryReader is a Reader for the FrontPutBinary structure.
type FrontPutBinaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontPutBinaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrontPutBinaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFrontPutBinaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFrontPutBinaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFrontPutBinaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFrontPutBinaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /frontend/binaries/{BinaryType}/{Uuid}] FrontPutBinary", response, response.Code())
	}
}

// NewFrontPutBinaryOK creates a FrontPutBinaryOK with default headers values
func NewFrontPutBinaryOK() *FrontPutBinaryOK {
	return &FrontPutBinaryOK{}
}

/*
FrontPutBinaryOK describes a response with status code 200, with default header values.

A successful response.
*/
type FrontPutBinaryOK struct {
	Payload *models.RestFrontBinaryResponse
}

// IsSuccess returns true when this front put binary o k response has a 2xx status code
func (o *FrontPutBinaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this front put binary o k response has a 3xx status code
func (o *FrontPutBinaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front put binary o k response has a 4xx status code
func (o *FrontPutBinaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this front put binary o k response has a 5xx status code
func (o *FrontPutBinaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this front put binary o k response a status code equal to that given
func (o *FrontPutBinaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the front put binary o k response
func (o *FrontPutBinaryOK) Code() int {
	return 200
}

func (o *FrontPutBinaryOK) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryOK  %+v", 200, o.Payload)
}

func (o *FrontPutBinaryOK) String() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryOK  %+v", 200, o.Payload)
}

func (o *FrontPutBinaryOK) GetPayload() *models.RestFrontBinaryResponse {
	return o.Payload
}

func (o *FrontPutBinaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestFrontBinaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPutBinaryUnauthorized creates a FrontPutBinaryUnauthorized with default headers values
func NewFrontPutBinaryUnauthorized() *FrontPutBinaryUnauthorized {
	return &FrontPutBinaryUnauthorized{}
}

/*
FrontPutBinaryUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type FrontPutBinaryUnauthorized struct {
}

// IsSuccess returns true when this front put binary unauthorized response has a 2xx status code
func (o *FrontPutBinaryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front put binary unauthorized response has a 3xx status code
func (o *FrontPutBinaryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front put binary unauthorized response has a 4xx status code
func (o *FrontPutBinaryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this front put binary unauthorized response has a 5xx status code
func (o *FrontPutBinaryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this front put binary unauthorized response a status code equal to that given
func (o *FrontPutBinaryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the front put binary unauthorized response
func (o *FrontPutBinaryUnauthorized) Code() int {
	return 401
}

func (o *FrontPutBinaryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryUnauthorized ", 401)
}

func (o *FrontPutBinaryUnauthorized) String() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryUnauthorized ", 401)
}

func (o *FrontPutBinaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFrontPutBinaryForbidden creates a FrontPutBinaryForbidden with default headers values
func NewFrontPutBinaryForbidden() *FrontPutBinaryForbidden {
	return &FrontPutBinaryForbidden{}
}

/*
FrontPutBinaryForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type FrontPutBinaryForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front put binary forbidden response has a 2xx status code
func (o *FrontPutBinaryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front put binary forbidden response has a 3xx status code
func (o *FrontPutBinaryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front put binary forbidden response has a 4xx status code
func (o *FrontPutBinaryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this front put binary forbidden response has a 5xx status code
func (o *FrontPutBinaryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this front put binary forbidden response a status code equal to that given
func (o *FrontPutBinaryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the front put binary forbidden response
func (o *FrontPutBinaryForbidden) Code() int {
	return 403
}

func (o *FrontPutBinaryForbidden) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryForbidden  %+v", 403, o.Payload)
}

func (o *FrontPutBinaryForbidden) String() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryForbidden  %+v", 403, o.Payload)
}

func (o *FrontPutBinaryForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPutBinaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPutBinaryNotFound creates a FrontPutBinaryNotFound with default headers values
func NewFrontPutBinaryNotFound() *FrontPutBinaryNotFound {
	return &FrontPutBinaryNotFound{}
}

/*
FrontPutBinaryNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type FrontPutBinaryNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front put binary not found response has a 2xx status code
func (o *FrontPutBinaryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front put binary not found response has a 3xx status code
func (o *FrontPutBinaryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front put binary not found response has a 4xx status code
func (o *FrontPutBinaryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this front put binary not found response has a 5xx status code
func (o *FrontPutBinaryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this front put binary not found response a status code equal to that given
func (o *FrontPutBinaryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the front put binary not found response
func (o *FrontPutBinaryNotFound) Code() int {
	return 404
}

func (o *FrontPutBinaryNotFound) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryNotFound  %+v", 404, o.Payload)
}

func (o *FrontPutBinaryNotFound) String() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryNotFound  %+v", 404, o.Payload)
}

func (o *FrontPutBinaryNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPutBinaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPutBinaryInternalServerError creates a FrontPutBinaryInternalServerError with default headers values
func NewFrontPutBinaryInternalServerError() *FrontPutBinaryInternalServerError {
	return &FrontPutBinaryInternalServerError{}
}

/*
FrontPutBinaryInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type FrontPutBinaryInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front put binary internal server error response has a 2xx status code
func (o *FrontPutBinaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front put binary internal server error response has a 3xx status code
func (o *FrontPutBinaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front put binary internal server error response has a 4xx status code
func (o *FrontPutBinaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this front put binary internal server error response has a 5xx status code
func (o *FrontPutBinaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this front put binary internal server error response a status code equal to that given
func (o *FrontPutBinaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the front put binary internal server error response
func (o *FrontPutBinaryInternalServerError) Code() int {
	return 500
}

func (o *FrontPutBinaryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontPutBinaryInternalServerError) String() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontPutBinaryInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPutBinaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
