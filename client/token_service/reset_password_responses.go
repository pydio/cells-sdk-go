// Code generated by go-swagger; DO NOT EDIT.

package token_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// ResetPasswordReader is a Reader for the ResetPassword structure.
type ResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResetPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResetPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResetPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/reset-password] ResetPassword", response, response.Code())
	}
}

// NewResetPasswordOK creates a ResetPasswordOK with default headers values
func NewResetPasswordOK() *ResetPasswordOK {
	return &ResetPasswordOK{}
}

/*
ResetPasswordOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResetPasswordOK struct {
	Payload *models.RestResetPasswordResponse
}

// IsSuccess returns true when this reset password o k response has a 2xx status code
func (o *ResetPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reset password o k response has a 3xx status code
func (o *ResetPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password o k response has a 4xx status code
func (o *ResetPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset password o k response has a 5xx status code
func (o *ResetPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password o k response a status code equal to that given
func (o *ResetPasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reset password o k response
func (o *ResetPasswordOK) Code() int {
	return 200
}

func (o *ResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordOK  %+v", 200, o.Payload)
}

func (o *ResetPasswordOK) String() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordOK  %+v", 200, o.Payload)
}

func (o *ResetPasswordOK) GetPayload() *models.RestResetPasswordResponse {
	return o.Payload
}

func (o *ResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestResetPasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordUnauthorized creates a ResetPasswordUnauthorized with default headers values
func NewResetPasswordUnauthorized() *ResetPasswordUnauthorized {
	return &ResetPasswordUnauthorized{}
}

/*
ResetPasswordUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ResetPasswordUnauthorized struct {
}

// IsSuccess returns true when this reset password unauthorized response has a 2xx status code
func (o *ResetPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password unauthorized response has a 3xx status code
func (o *ResetPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password unauthorized response has a 4xx status code
func (o *ResetPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset password unauthorized response has a 5xx status code
func (o *ResetPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password unauthorized response a status code equal to that given
func (o *ResetPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the reset password unauthorized response
func (o *ResetPasswordUnauthorized) Code() int {
	return 401
}

func (o *ResetPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordUnauthorized ", 401)
}

func (o *ResetPasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordUnauthorized ", 401)
}

func (o *ResetPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetPasswordForbidden creates a ResetPasswordForbidden with default headers values
func NewResetPasswordForbidden() *ResetPasswordForbidden {
	return &ResetPasswordForbidden{}
}

/*
ResetPasswordForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ResetPasswordForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this reset password forbidden response has a 2xx status code
func (o *ResetPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password forbidden response has a 3xx status code
func (o *ResetPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password forbidden response has a 4xx status code
func (o *ResetPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset password forbidden response has a 5xx status code
func (o *ResetPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password forbidden response a status code equal to that given
func (o *ResetPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the reset password forbidden response
func (o *ResetPasswordForbidden) Code() int {
	return 403
}

func (o *ResetPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *ResetPasswordForbidden) String() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *ResetPasswordForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ResetPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordNotFound creates a ResetPasswordNotFound with default headers values
func NewResetPasswordNotFound() *ResetPasswordNotFound {
	return &ResetPasswordNotFound{}
}

/*
ResetPasswordNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ResetPasswordNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this reset password not found response has a 2xx status code
func (o *ResetPasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password not found response has a 3xx status code
func (o *ResetPasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password not found response has a 4xx status code
func (o *ResetPasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset password not found response has a 5xx status code
func (o *ResetPasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password not found response a status code equal to that given
func (o *ResetPasswordNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reset password not found response
func (o *ResetPasswordNotFound) Code() int {
	return 404
}

func (o *ResetPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *ResetPasswordNotFound) String() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *ResetPasswordNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ResetPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordInternalServerError creates a ResetPasswordInternalServerError with default headers values
func NewResetPasswordInternalServerError() *ResetPasswordInternalServerError {
	return &ResetPasswordInternalServerError{}
}

/*
ResetPasswordInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ResetPasswordInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this reset password internal server error response has a 2xx status code
func (o *ResetPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password internal server error response has a 3xx status code
func (o *ResetPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password internal server error response has a 4xx status code
func (o *ResetPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset password internal server error response has a 5xx status code
func (o *ResetPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this reset password internal server error response a status code equal to that given
func (o *ResetPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the reset password internal server error response
func (o *ResetPasswordInternalServerError) Code() int {
	return 500
}

func (o *ResetPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *ResetPasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *ResetPasswordInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ResetPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
