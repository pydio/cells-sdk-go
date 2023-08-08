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

// ListDataSourcesReader is a Reader for the ListDataSources structure.
type ListDataSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDataSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDataSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListDataSourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListDataSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListDataSourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListDataSourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config/datasource] ListDataSources", response, response.Code())
	}
}

// NewListDataSourcesOK creates a ListDataSourcesOK with default headers values
func NewListDataSourcesOK() *ListDataSourcesOK {
	return &ListDataSourcesOK{}
}

/*
ListDataSourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListDataSourcesOK struct {
	Payload *models.RestDataSourceCollection
}

// IsSuccess returns true when this list data sources o k response has a 2xx status code
func (o *ListDataSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list data sources o k response has a 3xx status code
func (o *ListDataSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list data sources o k response has a 4xx status code
func (o *ListDataSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list data sources o k response has a 5xx status code
func (o *ListDataSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list data sources o k response a status code equal to that given
func (o *ListDataSourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list data sources o k response
func (o *ListDataSourcesOK) Code() int {
	return 200
}

func (o *ListDataSourcesOK) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesOK  %+v", 200, o.Payload)
}

func (o *ListDataSourcesOK) String() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesOK  %+v", 200, o.Payload)
}

func (o *ListDataSourcesOK) GetPayload() *models.RestDataSourceCollection {
	return o.Payload
}

func (o *ListDataSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDataSourceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDataSourcesUnauthorized creates a ListDataSourcesUnauthorized with default headers values
func NewListDataSourcesUnauthorized() *ListDataSourcesUnauthorized {
	return &ListDataSourcesUnauthorized{}
}

/*
ListDataSourcesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListDataSourcesUnauthorized struct {
}

// IsSuccess returns true when this list data sources unauthorized response has a 2xx status code
func (o *ListDataSourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list data sources unauthorized response has a 3xx status code
func (o *ListDataSourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list data sources unauthorized response has a 4xx status code
func (o *ListDataSourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list data sources unauthorized response has a 5xx status code
func (o *ListDataSourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list data sources unauthorized response a status code equal to that given
func (o *ListDataSourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list data sources unauthorized response
func (o *ListDataSourcesUnauthorized) Code() int {
	return 401
}

func (o *ListDataSourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesUnauthorized ", 401)
}

func (o *ListDataSourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesUnauthorized ", 401)
}

func (o *ListDataSourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListDataSourcesForbidden creates a ListDataSourcesForbidden with default headers values
func NewListDataSourcesForbidden() *ListDataSourcesForbidden {
	return &ListDataSourcesForbidden{}
}

/*
ListDataSourcesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListDataSourcesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list data sources forbidden response has a 2xx status code
func (o *ListDataSourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list data sources forbidden response has a 3xx status code
func (o *ListDataSourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list data sources forbidden response has a 4xx status code
func (o *ListDataSourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list data sources forbidden response has a 5xx status code
func (o *ListDataSourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list data sources forbidden response a status code equal to that given
func (o *ListDataSourcesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list data sources forbidden response
func (o *ListDataSourcesForbidden) Code() int {
	return 403
}

func (o *ListDataSourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesForbidden  %+v", 403, o.Payload)
}

func (o *ListDataSourcesForbidden) String() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesForbidden  %+v", 403, o.Payload)
}

func (o *ListDataSourcesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListDataSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDataSourcesNotFound creates a ListDataSourcesNotFound with default headers values
func NewListDataSourcesNotFound() *ListDataSourcesNotFound {
	return &ListDataSourcesNotFound{}
}

/*
ListDataSourcesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListDataSourcesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list data sources not found response has a 2xx status code
func (o *ListDataSourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list data sources not found response has a 3xx status code
func (o *ListDataSourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list data sources not found response has a 4xx status code
func (o *ListDataSourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list data sources not found response has a 5xx status code
func (o *ListDataSourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list data sources not found response a status code equal to that given
func (o *ListDataSourcesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list data sources not found response
func (o *ListDataSourcesNotFound) Code() int {
	return 404
}

func (o *ListDataSourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesNotFound  %+v", 404, o.Payload)
}

func (o *ListDataSourcesNotFound) String() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesNotFound  %+v", 404, o.Payload)
}

func (o *ListDataSourcesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListDataSourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDataSourcesInternalServerError creates a ListDataSourcesInternalServerError with default headers values
func NewListDataSourcesInternalServerError() *ListDataSourcesInternalServerError {
	return &ListDataSourcesInternalServerError{}
}

/*
ListDataSourcesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListDataSourcesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list data sources internal server error response has a 2xx status code
func (o *ListDataSourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list data sources internal server error response has a 3xx status code
func (o *ListDataSourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list data sources internal server error response has a 4xx status code
func (o *ListDataSourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list data sources internal server error response has a 5xx status code
func (o *ListDataSourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list data sources internal server error response a status code equal to that given
func (o *ListDataSourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list data sources internal server error response
func (o *ListDataSourcesInternalServerError) Code() int {
	return 500
}

func (o *ListDataSourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDataSourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDataSourcesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListDataSourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
