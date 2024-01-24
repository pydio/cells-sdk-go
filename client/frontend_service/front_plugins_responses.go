// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// FrontPluginsReader is a Reader for the FrontPlugins structure.
type FrontPluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontPluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrontPluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFrontPluginsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFrontPluginsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFrontPluginsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFrontPluginsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /frontend/plugins/{Lang}] FrontPlugins", response, response.Code())
	}
}

// NewFrontPluginsOK creates a FrontPluginsOK with default headers values
func NewFrontPluginsOK() *FrontPluginsOK {
	return &FrontPluginsOK{}
}

/*
FrontPluginsOK describes a response with status code 200, with default header values.

A successful response.
*/
type FrontPluginsOK struct {
	Payload *models.RestFrontPluginsResponse
}

// IsSuccess returns true when this front plugins o k response has a 2xx status code
func (o *FrontPluginsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this front plugins o k response has a 3xx status code
func (o *FrontPluginsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front plugins o k response has a 4xx status code
func (o *FrontPluginsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this front plugins o k response has a 5xx status code
func (o *FrontPluginsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this front plugins o k response a status code equal to that given
func (o *FrontPluginsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the front plugins o k response
func (o *FrontPluginsOK) Code() int {
	return 200
}

func (o *FrontPluginsOK) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsOK  %+v", 200, o.Payload)
}

func (o *FrontPluginsOK) String() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsOK  %+v", 200, o.Payload)
}

func (o *FrontPluginsOK) GetPayload() *models.RestFrontPluginsResponse {
	return o.Payload
}

func (o *FrontPluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestFrontPluginsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPluginsUnauthorized creates a FrontPluginsUnauthorized with default headers values
func NewFrontPluginsUnauthorized() *FrontPluginsUnauthorized {
	return &FrontPluginsUnauthorized{}
}

/*
FrontPluginsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type FrontPluginsUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front plugins unauthorized response has a 2xx status code
func (o *FrontPluginsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front plugins unauthorized response has a 3xx status code
func (o *FrontPluginsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front plugins unauthorized response has a 4xx status code
func (o *FrontPluginsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this front plugins unauthorized response has a 5xx status code
func (o *FrontPluginsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this front plugins unauthorized response a status code equal to that given
func (o *FrontPluginsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the front plugins unauthorized response
func (o *FrontPluginsUnauthorized) Code() int {
	return 401
}

func (o *FrontPluginsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsUnauthorized  %+v", 401, o.Payload)
}

func (o *FrontPluginsUnauthorized) String() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsUnauthorized  %+v", 401, o.Payload)
}

func (o *FrontPluginsUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPluginsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPluginsForbidden creates a FrontPluginsForbidden with default headers values
func NewFrontPluginsForbidden() *FrontPluginsForbidden {
	return &FrontPluginsForbidden{}
}

/*
FrontPluginsForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type FrontPluginsForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front plugins forbidden response has a 2xx status code
func (o *FrontPluginsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front plugins forbidden response has a 3xx status code
func (o *FrontPluginsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front plugins forbidden response has a 4xx status code
func (o *FrontPluginsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this front plugins forbidden response has a 5xx status code
func (o *FrontPluginsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this front plugins forbidden response a status code equal to that given
func (o *FrontPluginsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the front plugins forbidden response
func (o *FrontPluginsForbidden) Code() int {
	return 403
}

func (o *FrontPluginsForbidden) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsForbidden  %+v", 403, o.Payload)
}

func (o *FrontPluginsForbidden) String() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsForbidden  %+v", 403, o.Payload)
}

func (o *FrontPluginsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPluginsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPluginsNotFound creates a FrontPluginsNotFound with default headers values
func NewFrontPluginsNotFound() *FrontPluginsNotFound {
	return &FrontPluginsNotFound{}
}

/*
FrontPluginsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type FrontPluginsNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front plugins not found response has a 2xx status code
func (o *FrontPluginsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front plugins not found response has a 3xx status code
func (o *FrontPluginsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front plugins not found response has a 4xx status code
func (o *FrontPluginsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this front plugins not found response has a 5xx status code
func (o *FrontPluginsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this front plugins not found response a status code equal to that given
func (o *FrontPluginsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the front plugins not found response
func (o *FrontPluginsNotFound) Code() int {
	return 404
}

func (o *FrontPluginsNotFound) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsNotFound  %+v", 404, o.Payload)
}

func (o *FrontPluginsNotFound) String() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsNotFound  %+v", 404, o.Payload)
}

func (o *FrontPluginsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPluginsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontPluginsInternalServerError creates a FrontPluginsInternalServerError with default headers values
func NewFrontPluginsInternalServerError() *FrontPluginsInternalServerError {
	return &FrontPluginsInternalServerError{}
}

/*
FrontPluginsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type FrontPluginsInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front plugins internal server error response has a 2xx status code
func (o *FrontPluginsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front plugins internal server error response has a 3xx status code
func (o *FrontPluginsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front plugins internal server error response has a 4xx status code
func (o *FrontPluginsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this front plugins internal server error response has a 5xx status code
func (o *FrontPluginsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this front plugins internal server error response a status code equal to that given
func (o *FrontPluginsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the front plugins internal server error response
func (o *FrontPluginsInternalServerError) Code() int {
	return 500
}

func (o *FrontPluginsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontPluginsInternalServerError) String() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontPluginsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontPluginsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
