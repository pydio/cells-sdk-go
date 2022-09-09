// Code generated by go-swagger; DO NOT EDIT.

package update_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// UpdateRequiredReader is a Reader for the UpdateRequired structure.
type UpdateRequiredReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRequiredReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRequiredOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRequiredUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRequiredForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRequiredNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRequiredInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRequiredOK creates a UpdateRequiredOK with default headers values
func NewUpdateRequiredOK() *UpdateRequiredOK {
	return &UpdateRequiredOK{}
}

/*
UpdateRequiredOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRequiredOK struct {
	Payload *models.UpdateUpdateResponse
}

// IsSuccess returns true when this update required o k response has a 2xx status code
func (o *UpdateRequiredOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update required o k response has a 3xx status code
func (o *UpdateRequiredOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update required o k response has a 4xx status code
func (o *UpdateRequiredOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update required o k response has a 5xx status code
func (o *UpdateRequiredOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update required o k response a status code equal to that given
func (o *UpdateRequiredOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRequiredOK) Error() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredOK  %+v", 200, o.Payload)
}

func (o *UpdateRequiredOK) String() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredOK  %+v", 200, o.Payload)
}

func (o *UpdateRequiredOK) GetPayload() *models.UpdateUpdateResponse {
	return o.Payload
}

func (o *UpdateRequiredOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRequiredUnauthorized creates a UpdateRequiredUnauthorized with default headers values
func NewUpdateRequiredUnauthorized() *UpdateRequiredUnauthorized {
	return &UpdateRequiredUnauthorized{}
}

/*
UpdateRequiredUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type UpdateRequiredUnauthorized struct {
}

// IsSuccess returns true when this update required unauthorized response has a 2xx status code
func (o *UpdateRequiredUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update required unauthorized response has a 3xx status code
func (o *UpdateRequiredUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update required unauthorized response has a 4xx status code
func (o *UpdateRequiredUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update required unauthorized response has a 5xx status code
func (o *UpdateRequiredUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update required unauthorized response a status code equal to that given
func (o *UpdateRequiredUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRequiredUnauthorized) Error() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredUnauthorized ", 401)
}

func (o *UpdateRequiredUnauthorized) String() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredUnauthorized ", 401)
}

func (o *UpdateRequiredUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRequiredForbidden creates a UpdateRequiredForbidden with default headers values
func NewUpdateRequiredForbidden() *UpdateRequiredForbidden {
	return &UpdateRequiredForbidden{}
}

/*
UpdateRequiredForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type UpdateRequiredForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update required forbidden response has a 2xx status code
func (o *UpdateRequiredForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update required forbidden response has a 3xx status code
func (o *UpdateRequiredForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update required forbidden response has a 4xx status code
func (o *UpdateRequiredForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update required forbidden response has a 5xx status code
func (o *UpdateRequiredForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update required forbidden response a status code equal to that given
func (o *UpdateRequiredForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRequiredForbidden) Error() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRequiredForbidden) String() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRequiredForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateRequiredForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRequiredNotFound creates a UpdateRequiredNotFound with default headers values
func NewUpdateRequiredNotFound() *UpdateRequiredNotFound {
	return &UpdateRequiredNotFound{}
}

/*
UpdateRequiredNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type UpdateRequiredNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update required not found response has a 2xx status code
func (o *UpdateRequiredNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update required not found response has a 3xx status code
func (o *UpdateRequiredNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update required not found response has a 4xx status code
func (o *UpdateRequiredNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update required not found response has a 5xx status code
func (o *UpdateRequiredNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update required not found response a status code equal to that given
func (o *UpdateRequiredNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRequiredNotFound) Error() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRequiredNotFound) String() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRequiredNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateRequiredNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRequiredInternalServerError creates a UpdateRequiredInternalServerError with default headers values
func NewUpdateRequiredInternalServerError() *UpdateRequiredInternalServerError {
	return &UpdateRequiredInternalServerError{}
}

/*
UpdateRequiredInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type UpdateRequiredInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update required internal server error response has a 2xx status code
func (o *UpdateRequiredInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update required internal server error response has a 3xx status code
func (o *UpdateRequiredInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update required internal server error response has a 4xx status code
func (o *UpdateRequiredInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update required internal server error response has a 5xx status code
func (o *UpdateRequiredInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update required internal server error response a status code equal to that given
func (o *UpdateRequiredInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateRequiredInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRequiredInternalServerError) String() string {
	return fmt.Sprintf("[POST /update][%d] updateRequiredInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRequiredInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateRequiredInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
