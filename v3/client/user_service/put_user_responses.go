// Code generated by go-swagger; DO NOT EDIT.

package user_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PutUserReader is a Reader for the PutUser structure.
type PutUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserOK creates a PutUserOK with default headers values
func NewPutUserOK() *PutUserOK {
	return &PutUserOK{}
}

/* PutUserOK describes a response with status code 200, with default header values.

PutUserOK put user o k
*/
type PutUserOK struct {
	Payload *models.IdmUser
}

func (o *PutUserOK) Error() string {
	return fmt.Sprintf("[PUT /user/{Login}][%d] putUserOK  %+v", 200, o.Payload)
}
func (o *PutUserOK) GetPayload() *models.IdmUser {
	return o.Payload
}

func (o *PutUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserUnauthorized creates a PutUserUnauthorized with default headers values
func NewPutUserUnauthorized() *PutUserUnauthorized {
	return &PutUserUnauthorized{}
}

/* PutUserUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutUserUnauthorized struct {
	Payload *models.RestError
}

func (o *PutUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user/{Login}][%d] putUserUnauthorized  %+v", 401, o.Payload)
}
func (o *PutUserUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserForbidden creates a PutUserForbidden with default headers values
func NewPutUserForbidden() *PutUserForbidden {
	return &PutUserForbidden{}
}

/* PutUserForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutUserForbidden struct {
	Payload *models.RestError
}

func (o *PutUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /user/{Login}][%d] putUserForbidden  %+v", 403, o.Payload)
}
func (o *PutUserForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserNotFound creates a PutUserNotFound with default headers values
func NewPutUserNotFound() *PutUserNotFound {
	return &PutUserNotFound{}
}

/* PutUserNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutUserNotFound struct {
	Payload *models.RestError
}

func (o *PutUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /user/{Login}][%d] putUserNotFound  %+v", 404, o.Payload)
}
func (o *PutUserNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserInternalServerError creates a PutUserInternalServerError with default headers values
func NewPutUserInternalServerError() *PutUserInternalServerError {
	return &PutUserInternalServerError{}
}

/* PutUserInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutUserInternalServerError struct {
	Payload *models.RestError
}

func (o *PutUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user/{Login}][%d] putUserInternalServerError  %+v", 500, o.Payload)
}
func (o *PutUserInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
