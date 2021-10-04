// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// DeleteEncryptionKeyReader is a Reader for the DeleteEncryptionKey structure.
type DeleteEncryptionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEncryptionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEncryptionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEncryptionKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEncryptionKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEncryptionKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteEncryptionKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEncryptionKeyOK creates a DeleteEncryptionKeyOK with default headers values
func NewDeleteEncryptionKeyOK() *DeleteEncryptionKeyOK {
	return &DeleteEncryptionKeyOK{}
}

/* DeleteEncryptionKeyOK describes a response with status code 200, with default header values.

DeleteEncryptionKeyOK delete encryption key o k
*/
type DeleteEncryptionKeyOK struct {
	Payload *models.EncryptionAdminDeleteKeyResponse
}

func (o *DeleteEncryptionKeyOK) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyOK  %+v", 200, o.Payload)
}
func (o *DeleteEncryptionKeyOK) GetPayload() *models.EncryptionAdminDeleteKeyResponse {
	return o.Payload
}

func (o *DeleteEncryptionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionAdminDeleteKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEncryptionKeyUnauthorized creates a DeleteEncryptionKeyUnauthorized with default headers values
func NewDeleteEncryptionKeyUnauthorized() *DeleteEncryptionKeyUnauthorized {
	return &DeleteEncryptionKeyUnauthorized{}
}

/* DeleteEncryptionKeyUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteEncryptionKeyUnauthorized struct {
	Payload *models.RestError
}

func (o *DeleteEncryptionKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteEncryptionKeyUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteEncryptionKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEncryptionKeyForbidden creates a DeleteEncryptionKeyForbidden with default headers values
func NewDeleteEncryptionKeyForbidden() *DeleteEncryptionKeyForbidden {
	return &DeleteEncryptionKeyForbidden{}
}

/* DeleteEncryptionKeyForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteEncryptionKeyForbidden struct {
	Payload *models.RestError
}

func (o *DeleteEncryptionKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyForbidden  %+v", 403, o.Payload)
}
func (o *DeleteEncryptionKeyForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteEncryptionKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEncryptionKeyNotFound creates a DeleteEncryptionKeyNotFound with default headers values
func NewDeleteEncryptionKeyNotFound() *DeleteEncryptionKeyNotFound {
	return &DeleteEncryptionKeyNotFound{}
}

/* DeleteEncryptionKeyNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteEncryptionKeyNotFound struct {
	Payload *models.RestError
}

func (o *DeleteEncryptionKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyNotFound  %+v", 404, o.Payload)
}
func (o *DeleteEncryptionKeyNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteEncryptionKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEncryptionKeyInternalServerError creates a DeleteEncryptionKeyInternalServerError with default headers values
func NewDeleteEncryptionKeyInternalServerError() *DeleteEncryptionKeyInternalServerError {
	return &DeleteEncryptionKeyInternalServerError{}
}

/* DeleteEncryptionKeyInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteEncryptionKeyInternalServerError struct {
	Payload *models.RestError
}

func (o *DeleteEncryptionKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteEncryptionKeyInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteEncryptionKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
