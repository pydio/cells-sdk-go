// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PerformInstallCheckReader is a Reader for the PerformInstallCheck structure.
type PerformInstallCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformInstallCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformInstallCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPerformInstallCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformInstallCheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformInstallCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformInstallCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPerformInstallCheckOK creates a PerformInstallCheckOK with default headers values
func NewPerformInstallCheckOK() *PerformInstallCheckOK {
	return &PerformInstallCheckOK{}
}

/* PerformInstallCheckOK describes a response with status code 200, with default header values.

PerformInstallCheckOK perform install check o k
*/
type PerformInstallCheckOK struct {
	Payload *models.InstallPerformCheckResponse
}

func (o *PerformInstallCheckOK) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckOK  %+v", 200, o.Payload)
}
func (o *PerformInstallCheckOK) GetPayload() *models.InstallPerformCheckResponse {
	return o.Payload
}

func (o *PerformInstallCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallPerformCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckUnauthorized creates a PerformInstallCheckUnauthorized with default headers values
func NewPerformInstallCheckUnauthorized() *PerformInstallCheckUnauthorized {
	return &PerformInstallCheckUnauthorized{}
}

/* PerformInstallCheckUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PerformInstallCheckUnauthorized struct {
}

func (o *PerformInstallCheckUnauthorized) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckUnauthorized ", 401)
}

func (o *PerformInstallCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPerformInstallCheckForbidden creates a PerformInstallCheckForbidden with default headers values
func NewPerformInstallCheckForbidden() *PerformInstallCheckForbidden {
	return &PerformInstallCheckForbidden{}
}

/* PerformInstallCheckForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PerformInstallCheckForbidden struct {
	Payload *models.RestError
}

func (o *PerformInstallCheckForbidden) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckForbidden  %+v", 403, o.Payload)
}
func (o *PerformInstallCheckForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckNotFound creates a PerformInstallCheckNotFound with default headers values
func NewPerformInstallCheckNotFound() *PerformInstallCheckNotFound {
	return &PerformInstallCheckNotFound{}
}

/* PerformInstallCheckNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PerformInstallCheckNotFound struct {
	Payload *models.RestError
}

func (o *PerformInstallCheckNotFound) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckNotFound  %+v", 404, o.Payload)
}
func (o *PerformInstallCheckNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckInternalServerError creates a PerformInstallCheckInternalServerError with default headers values
func NewPerformInstallCheckInternalServerError() *PerformInstallCheckInternalServerError {
	return &PerformInstallCheckInternalServerError{}
}

/* PerformInstallCheckInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PerformInstallCheckInternalServerError struct {
	Payload *models.RestError
}

func (o *PerformInstallCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckInternalServerError  %+v", 500, o.Payload)
}
func (o *PerformInstallCheckInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
