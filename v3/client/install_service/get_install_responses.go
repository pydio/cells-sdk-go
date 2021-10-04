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

// GetInstallReader is a Reader for the GetInstall structure.
type GetInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetInstallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstallOK creates a GetInstallOK with default headers values
func NewGetInstallOK() *GetInstallOK {
	return &GetInstallOK{}
}

/* GetInstallOK describes a response with status code 200, with default header values.

GetInstallOK get install o k
*/
type GetInstallOK struct {
	Payload *models.InstallGetDefaultsResponse
}

func (o *GetInstallOK) Error() string {
	return fmt.Sprintf("[GET /install][%d] getInstallOK  %+v", 200, o.Payload)
}
func (o *GetInstallOK) GetPayload() *models.InstallGetDefaultsResponse {
	return o.Payload
}

func (o *GetInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallGetDefaultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallUnauthorized creates a GetInstallUnauthorized with default headers values
func NewGetInstallUnauthorized() *GetInstallUnauthorized {
	return &GetInstallUnauthorized{}
}

/* GetInstallUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetInstallUnauthorized struct {
	Payload *models.RestError
}

func (o *GetInstallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /install][%d] getInstallUnauthorized  %+v", 401, o.Payload)
}
func (o *GetInstallUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetInstallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallForbidden creates a GetInstallForbidden with default headers values
func NewGetInstallForbidden() *GetInstallForbidden {
	return &GetInstallForbidden{}
}

/* GetInstallForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetInstallForbidden struct {
	Payload *models.RestError
}

func (o *GetInstallForbidden) Error() string {
	return fmt.Sprintf("[GET /install][%d] getInstallForbidden  %+v", 403, o.Payload)
}
func (o *GetInstallForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetInstallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallNotFound creates a GetInstallNotFound with default headers values
func NewGetInstallNotFound() *GetInstallNotFound {
	return &GetInstallNotFound{}
}

/* GetInstallNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetInstallNotFound struct {
	Payload *models.RestError
}

func (o *GetInstallNotFound) Error() string {
	return fmt.Sprintf("[GET /install][%d] getInstallNotFound  %+v", 404, o.Payload)
}
func (o *GetInstallNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetInstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallInternalServerError creates a GetInstallInternalServerError with default headers values
func NewGetInstallInternalServerError() *GetInstallInternalServerError {
	return &GetInstallInternalServerError{}
}

/* GetInstallInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetInstallInternalServerError struct {
	Payload *models.RestError
}

func (o *GetInstallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /install][%d] getInstallInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInstallInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetInstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
