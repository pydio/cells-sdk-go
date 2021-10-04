// Code generated by go-swagger; DO NOT EDIT.

package token_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetPasswordOK creates a ResetPasswordOK with default headers values
func NewResetPasswordOK() *ResetPasswordOK {
	return &ResetPasswordOK{}
}

/* ResetPasswordOK describes a response with status code 200, with default header values.

ResetPasswordOK reset password o k
*/
type ResetPasswordOK struct {
	Payload *models.RestResetPasswordResponse
}

func (o *ResetPasswordOK) Error() string {
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

/* ResetPasswordUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ResetPasswordUnauthorized struct {
	Payload *models.RestError
}

func (o *ResetPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/reset-password][%d] resetPasswordUnauthorized  %+v", 401, o.Payload)
}
func (o *ResetPasswordUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ResetPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordForbidden creates a ResetPasswordForbidden with default headers values
func NewResetPasswordForbidden() *ResetPasswordForbidden {
	return &ResetPasswordForbidden{}
}

/* ResetPasswordForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ResetPasswordForbidden struct {
	Payload *models.RestError
}

func (o *ResetPasswordForbidden) Error() string {
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

/* ResetPasswordNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ResetPasswordNotFound struct {
	Payload *models.RestError
}

func (o *ResetPasswordNotFound) Error() string {
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

/* ResetPasswordInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ResetPasswordInternalServerError struct {
	Payload *models.RestError
}

func (o *ResetPasswordInternalServerError) Error() string {
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
