// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// CreateEncryptionKeyReader is a Reader for the CreateEncryptionKey structure.
type CreateEncryptionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEncryptionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEncryptionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateEncryptionKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEncryptionKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEncryptionKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEncryptionKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /config/encryption/create] CreateEncryptionKey", response, response.Code())
	}
}

// NewCreateEncryptionKeyOK creates a CreateEncryptionKeyOK with default headers values
func NewCreateEncryptionKeyOK() *CreateEncryptionKeyOK {
	return &CreateEncryptionKeyOK{}
}

/*
CreateEncryptionKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEncryptionKeyOK struct {
	Payload *models.EncryptionAdminCreateKeyResponse
}

// IsSuccess returns true when this create encryption key o k response has a 2xx status code
func (o *CreateEncryptionKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create encryption key o k response has a 3xx status code
func (o *CreateEncryptionKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create encryption key o k response has a 4xx status code
func (o *CreateEncryptionKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create encryption key o k response has a 5xx status code
func (o *CreateEncryptionKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create encryption key o k response a status code equal to that given
func (o *CreateEncryptionKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create encryption key o k response
func (o *CreateEncryptionKeyOK) Code() int {
	return 200
}

func (o *CreateEncryptionKeyOK) Error() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyOK  %+v", 200, o.Payload)
}

func (o *CreateEncryptionKeyOK) String() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyOK  %+v", 200, o.Payload)
}

func (o *CreateEncryptionKeyOK) GetPayload() *models.EncryptionAdminCreateKeyResponse {
	return o.Payload
}

func (o *CreateEncryptionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionAdminCreateKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEncryptionKeyUnauthorized creates a CreateEncryptionKeyUnauthorized with default headers values
func NewCreateEncryptionKeyUnauthorized() *CreateEncryptionKeyUnauthorized {
	return &CreateEncryptionKeyUnauthorized{}
}

/*
CreateEncryptionKeyUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type CreateEncryptionKeyUnauthorized struct {
}

// IsSuccess returns true when this create encryption key unauthorized response has a 2xx status code
func (o *CreateEncryptionKeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create encryption key unauthorized response has a 3xx status code
func (o *CreateEncryptionKeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create encryption key unauthorized response has a 4xx status code
func (o *CreateEncryptionKeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create encryption key unauthorized response has a 5xx status code
func (o *CreateEncryptionKeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create encryption key unauthorized response a status code equal to that given
func (o *CreateEncryptionKeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create encryption key unauthorized response
func (o *CreateEncryptionKeyUnauthorized) Code() int {
	return 401
}

func (o *CreateEncryptionKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyUnauthorized ", 401)
}

func (o *CreateEncryptionKeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyUnauthorized ", 401)
}

func (o *CreateEncryptionKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEncryptionKeyForbidden creates a CreateEncryptionKeyForbidden with default headers values
func NewCreateEncryptionKeyForbidden() *CreateEncryptionKeyForbidden {
	return &CreateEncryptionKeyForbidden{}
}

/*
CreateEncryptionKeyForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type CreateEncryptionKeyForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create encryption key forbidden response has a 2xx status code
func (o *CreateEncryptionKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create encryption key forbidden response has a 3xx status code
func (o *CreateEncryptionKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create encryption key forbidden response has a 4xx status code
func (o *CreateEncryptionKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create encryption key forbidden response has a 5xx status code
func (o *CreateEncryptionKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create encryption key forbidden response a status code equal to that given
func (o *CreateEncryptionKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create encryption key forbidden response
func (o *CreateEncryptionKeyForbidden) Code() int {
	return 403
}

func (o *CreateEncryptionKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyForbidden  %+v", 403, o.Payload)
}

func (o *CreateEncryptionKeyForbidden) String() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyForbidden  %+v", 403, o.Payload)
}

func (o *CreateEncryptionKeyForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateEncryptionKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEncryptionKeyNotFound creates a CreateEncryptionKeyNotFound with default headers values
func NewCreateEncryptionKeyNotFound() *CreateEncryptionKeyNotFound {
	return &CreateEncryptionKeyNotFound{}
}

/*
CreateEncryptionKeyNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type CreateEncryptionKeyNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create encryption key not found response has a 2xx status code
func (o *CreateEncryptionKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create encryption key not found response has a 3xx status code
func (o *CreateEncryptionKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create encryption key not found response has a 4xx status code
func (o *CreateEncryptionKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create encryption key not found response has a 5xx status code
func (o *CreateEncryptionKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create encryption key not found response a status code equal to that given
func (o *CreateEncryptionKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create encryption key not found response
func (o *CreateEncryptionKeyNotFound) Code() int {
	return 404
}

func (o *CreateEncryptionKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyNotFound  %+v", 404, o.Payload)
}

func (o *CreateEncryptionKeyNotFound) String() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyNotFound  %+v", 404, o.Payload)
}

func (o *CreateEncryptionKeyNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateEncryptionKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEncryptionKeyInternalServerError creates a CreateEncryptionKeyInternalServerError with default headers values
func NewCreateEncryptionKeyInternalServerError() *CreateEncryptionKeyInternalServerError {
	return &CreateEncryptionKeyInternalServerError{}
}

/*
CreateEncryptionKeyInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type CreateEncryptionKeyInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create encryption key internal server error response has a 2xx status code
func (o *CreateEncryptionKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create encryption key internal server error response has a 3xx status code
func (o *CreateEncryptionKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create encryption key internal server error response has a 4xx status code
func (o *CreateEncryptionKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create encryption key internal server error response has a 5xx status code
func (o *CreateEncryptionKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create encryption key internal server error response a status code equal to that given
func (o *CreateEncryptionKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create encryption key internal server error response
func (o *CreateEncryptionKeyInternalServerError) Code() int {
	return 500
}

func (o *CreateEncryptionKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEncryptionKeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /config/encryption/create][%d] createEncryptionKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEncryptionKeyInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreateEncryptionKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
