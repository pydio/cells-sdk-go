// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// ExportEncryptionKeyReader is a Reader for the ExportEncryptionKey structure.
type ExportEncryptionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportEncryptionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportEncryptionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportEncryptionKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportEncryptionKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportEncryptionKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExportEncryptionKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /config/encryption/export] ExportEncryptionKey", response, response.Code())
	}
}

// NewExportEncryptionKeyOK creates a ExportEncryptionKeyOK with default headers values
func NewExportEncryptionKeyOK() *ExportEncryptionKeyOK {
	return &ExportEncryptionKeyOK{}
}

/*
ExportEncryptionKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExportEncryptionKeyOK struct {
	Payload *models.EncryptionAdminExportKeyResponse
}

// IsSuccess returns true when this export encryption key o k response has a 2xx status code
func (o *ExportEncryptionKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export encryption key o k response has a 3xx status code
func (o *ExportEncryptionKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export encryption key o k response has a 4xx status code
func (o *ExportEncryptionKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export encryption key o k response has a 5xx status code
func (o *ExportEncryptionKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export encryption key o k response a status code equal to that given
func (o *ExportEncryptionKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export encryption key o k response
func (o *ExportEncryptionKeyOK) Code() int {
	return 200
}

func (o *ExportEncryptionKeyOK) Error() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyOK  %+v", 200, o.Payload)
}

func (o *ExportEncryptionKeyOK) String() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyOK  %+v", 200, o.Payload)
}

func (o *ExportEncryptionKeyOK) GetPayload() *models.EncryptionAdminExportKeyResponse {
	return o.Payload
}

func (o *ExportEncryptionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionAdminExportKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionKeyUnauthorized creates a ExportEncryptionKeyUnauthorized with default headers values
func NewExportEncryptionKeyUnauthorized() *ExportEncryptionKeyUnauthorized {
	return &ExportEncryptionKeyUnauthorized{}
}

/*
ExportEncryptionKeyUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ExportEncryptionKeyUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this export encryption key unauthorized response has a 2xx status code
func (o *ExportEncryptionKeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export encryption key unauthorized response has a 3xx status code
func (o *ExportEncryptionKeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export encryption key unauthorized response has a 4xx status code
func (o *ExportEncryptionKeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export encryption key unauthorized response has a 5xx status code
func (o *ExportEncryptionKeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export encryption key unauthorized response a status code equal to that given
func (o *ExportEncryptionKeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the export encryption key unauthorized response
func (o *ExportEncryptionKeyUnauthorized) Code() int {
	return 401
}

func (o *ExportEncryptionKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *ExportEncryptionKeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *ExportEncryptionKeyUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ExportEncryptionKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionKeyForbidden creates a ExportEncryptionKeyForbidden with default headers values
func NewExportEncryptionKeyForbidden() *ExportEncryptionKeyForbidden {
	return &ExportEncryptionKeyForbidden{}
}

/*
ExportEncryptionKeyForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type ExportEncryptionKeyForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this export encryption key forbidden response has a 2xx status code
func (o *ExportEncryptionKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export encryption key forbidden response has a 3xx status code
func (o *ExportEncryptionKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export encryption key forbidden response has a 4xx status code
func (o *ExportEncryptionKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this export encryption key forbidden response has a 5xx status code
func (o *ExportEncryptionKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this export encryption key forbidden response a status code equal to that given
func (o *ExportEncryptionKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the export encryption key forbidden response
func (o *ExportEncryptionKeyForbidden) Code() int {
	return 403
}

func (o *ExportEncryptionKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyForbidden  %+v", 403, o.Payload)
}

func (o *ExportEncryptionKeyForbidden) String() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyForbidden  %+v", 403, o.Payload)
}

func (o *ExportEncryptionKeyForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ExportEncryptionKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionKeyNotFound creates a ExportEncryptionKeyNotFound with default headers values
func NewExportEncryptionKeyNotFound() *ExportEncryptionKeyNotFound {
	return &ExportEncryptionKeyNotFound{}
}

/*
ExportEncryptionKeyNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ExportEncryptionKeyNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this export encryption key not found response has a 2xx status code
func (o *ExportEncryptionKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export encryption key not found response has a 3xx status code
func (o *ExportEncryptionKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export encryption key not found response has a 4xx status code
func (o *ExportEncryptionKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export encryption key not found response has a 5xx status code
func (o *ExportEncryptionKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export encryption key not found response a status code equal to that given
func (o *ExportEncryptionKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the export encryption key not found response
func (o *ExportEncryptionKeyNotFound) Code() int {
	return 404
}

func (o *ExportEncryptionKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyNotFound  %+v", 404, o.Payload)
}

func (o *ExportEncryptionKeyNotFound) String() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyNotFound  %+v", 404, o.Payload)
}

func (o *ExportEncryptionKeyNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ExportEncryptionKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionKeyInternalServerError creates a ExportEncryptionKeyInternalServerError with default headers values
func NewExportEncryptionKeyInternalServerError() *ExportEncryptionKeyInternalServerError {
	return &ExportEncryptionKeyInternalServerError{}
}

/*
ExportEncryptionKeyInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ExportEncryptionKeyInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this export encryption key internal server error response has a 2xx status code
func (o *ExportEncryptionKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export encryption key internal server error response has a 3xx status code
func (o *ExportEncryptionKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export encryption key internal server error response has a 4xx status code
func (o *ExportEncryptionKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this export encryption key internal server error response has a 5xx status code
func (o *ExportEncryptionKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this export encryption key internal server error response a status code equal to that given
func (o *ExportEncryptionKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the export encryption key internal server error response
func (o *ExportEncryptionKeyInternalServerError) Code() int {
	return 500
}

func (o *ExportEncryptionKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *ExportEncryptionKeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /config/encryption/export][%d] exportEncryptionKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *ExportEncryptionKeyInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ExportEncryptionKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
