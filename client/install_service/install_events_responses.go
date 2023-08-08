// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// InstallEventsReader is a Reader for the InstallEvents structure.
type InstallEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewInstallEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstallEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstallEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInstallEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /install/events] InstallEvents", response, response.Code())
	}
}

// NewInstallEventsOK creates a InstallEventsOK with default headers values
func NewInstallEventsOK() *InstallEventsOK {
	return &InstallEventsOK{}
}

/*
InstallEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type InstallEventsOK struct {
	Payload *models.InstallInstallEventsResponse
}

// IsSuccess returns true when this install events o k response has a 2xx status code
func (o *InstallEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this install events o k response has a 3xx status code
func (o *InstallEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install events o k response has a 4xx status code
func (o *InstallEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this install events o k response has a 5xx status code
func (o *InstallEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this install events o k response a status code equal to that given
func (o *InstallEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the install events o k response
func (o *InstallEventsOK) Code() int {
	return 200
}

func (o *InstallEventsOK) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsOK  %+v", 200, o.Payload)
}

func (o *InstallEventsOK) String() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsOK  %+v", 200, o.Payload)
}

func (o *InstallEventsOK) GetPayload() *models.InstallInstallEventsResponse {
	return o.Payload
}

func (o *InstallEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallInstallEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallEventsUnauthorized creates a InstallEventsUnauthorized with default headers values
func NewInstallEventsUnauthorized() *InstallEventsUnauthorized {
	return &InstallEventsUnauthorized{}
}

/*
InstallEventsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type InstallEventsUnauthorized struct {
}

// IsSuccess returns true when this install events unauthorized response has a 2xx status code
func (o *InstallEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this install events unauthorized response has a 3xx status code
func (o *InstallEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install events unauthorized response has a 4xx status code
func (o *InstallEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this install events unauthorized response has a 5xx status code
func (o *InstallEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this install events unauthorized response a status code equal to that given
func (o *InstallEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the install events unauthorized response
func (o *InstallEventsUnauthorized) Code() int {
	return 401
}

func (o *InstallEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsUnauthorized ", 401)
}

func (o *InstallEventsUnauthorized) String() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsUnauthorized ", 401)
}

func (o *InstallEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstallEventsForbidden creates a InstallEventsForbidden with default headers values
func NewInstallEventsForbidden() *InstallEventsForbidden {
	return &InstallEventsForbidden{}
}

/*
InstallEventsForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type InstallEventsForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this install events forbidden response has a 2xx status code
func (o *InstallEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this install events forbidden response has a 3xx status code
func (o *InstallEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install events forbidden response has a 4xx status code
func (o *InstallEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this install events forbidden response has a 5xx status code
func (o *InstallEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this install events forbidden response a status code equal to that given
func (o *InstallEventsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the install events forbidden response
func (o *InstallEventsForbidden) Code() int {
	return 403
}

func (o *InstallEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsForbidden  %+v", 403, o.Payload)
}

func (o *InstallEventsForbidden) String() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsForbidden  %+v", 403, o.Payload)
}

func (o *InstallEventsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *InstallEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallEventsNotFound creates a InstallEventsNotFound with default headers values
func NewInstallEventsNotFound() *InstallEventsNotFound {
	return &InstallEventsNotFound{}
}

/*
InstallEventsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type InstallEventsNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this install events not found response has a 2xx status code
func (o *InstallEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this install events not found response has a 3xx status code
func (o *InstallEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install events not found response has a 4xx status code
func (o *InstallEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this install events not found response has a 5xx status code
func (o *InstallEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this install events not found response a status code equal to that given
func (o *InstallEventsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the install events not found response
func (o *InstallEventsNotFound) Code() int {
	return 404
}

func (o *InstallEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsNotFound  %+v", 404, o.Payload)
}

func (o *InstallEventsNotFound) String() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsNotFound  %+v", 404, o.Payload)
}

func (o *InstallEventsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *InstallEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallEventsInternalServerError creates a InstallEventsInternalServerError with default headers values
func NewInstallEventsInternalServerError() *InstallEventsInternalServerError {
	return &InstallEventsInternalServerError{}
}

/*
InstallEventsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type InstallEventsInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this install events internal server error response has a 2xx status code
func (o *InstallEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this install events internal server error response has a 3xx status code
func (o *InstallEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install events internal server error response has a 4xx status code
func (o *InstallEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this install events internal server error response has a 5xx status code
func (o *InstallEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this install events internal server error response a status code equal to that given
func (o *InstallEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the install events internal server error response
func (o *InstallEventsInternalServerError) Code() int {
	return 500
}

func (o *InstallEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *InstallEventsInternalServerError) String() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *InstallEventsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *InstallEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
