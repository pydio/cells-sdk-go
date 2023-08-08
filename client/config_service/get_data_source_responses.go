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

// GetDataSourceReader is a Reader for the GetDataSource structure.
type GetDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDataSourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataSourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataSourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config/datasource/{Name}] GetDataSource", response, response.Code())
	}
}

// NewGetDataSourceOK creates a GetDataSourceOK with default headers values
func NewGetDataSourceOK() *GetDataSourceOK {
	return &GetDataSourceOK{}
}

/*
GetDataSourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetDataSourceOK struct {
	Payload *models.ObjectDataSource
}

// IsSuccess returns true when this get data source o k response has a 2xx status code
func (o *GetDataSourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data source o k response has a 3xx status code
func (o *GetDataSourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source o k response has a 4xx status code
func (o *GetDataSourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data source o k response has a 5xx status code
func (o *GetDataSourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source o k response a status code equal to that given
func (o *GetDataSourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data source o k response
func (o *GetDataSourceOK) Code() int {
	return 200
}

func (o *GetDataSourceOK) Error() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceOK  %+v", 200, o.Payload)
}

func (o *GetDataSourceOK) String() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceOK  %+v", 200, o.Payload)
}

func (o *GetDataSourceOK) GetPayload() *models.ObjectDataSource {
	return o.Payload
}

func (o *GetDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectDataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceUnauthorized creates a GetDataSourceUnauthorized with default headers values
func NewGetDataSourceUnauthorized() *GetDataSourceUnauthorized {
	return &GetDataSourceUnauthorized{}
}

/*
GetDataSourceUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetDataSourceUnauthorized struct {
}

// IsSuccess returns true when this get data source unauthorized response has a 2xx status code
func (o *GetDataSourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source unauthorized response has a 3xx status code
func (o *GetDataSourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source unauthorized response has a 4xx status code
func (o *GetDataSourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source unauthorized response has a 5xx status code
func (o *GetDataSourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source unauthorized response a status code equal to that given
func (o *GetDataSourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get data source unauthorized response
func (o *GetDataSourceUnauthorized) Code() int {
	return 401
}

func (o *GetDataSourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceUnauthorized ", 401)
}

func (o *GetDataSourceUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceUnauthorized ", 401)
}

func (o *GetDataSourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDataSourceForbidden creates a GetDataSourceForbidden with default headers values
func NewGetDataSourceForbidden() *GetDataSourceForbidden {
	return &GetDataSourceForbidden{}
}

/*
GetDataSourceForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetDataSourceForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get data source forbidden response has a 2xx status code
func (o *GetDataSourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source forbidden response has a 3xx status code
func (o *GetDataSourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source forbidden response has a 4xx status code
func (o *GetDataSourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source forbidden response has a 5xx status code
func (o *GetDataSourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source forbidden response a status code equal to that given
func (o *GetDataSourceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get data source forbidden response
func (o *GetDataSourceForbidden) Code() int {
	return 403
}

func (o *GetDataSourceForbidden) Error() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourceForbidden) String() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourceForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetDataSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceNotFound creates a GetDataSourceNotFound with default headers values
func NewGetDataSourceNotFound() *GetDataSourceNotFound {
	return &GetDataSourceNotFound{}
}

/*
GetDataSourceNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetDataSourceNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get data source not found response has a 2xx status code
func (o *GetDataSourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source not found response has a 3xx status code
func (o *GetDataSourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source not found response has a 4xx status code
func (o *GetDataSourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source not found response has a 5xx status code
func (o *GetDataSourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source not found response a status code equal to that given
func (o *GetDataSourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get data source not found response
func (o *GetDataSourceNotFound) Code() int {
	return 404
}

func (o *GetDataSourceNotFound) Error() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceNotFound  %+v", 404, o.Payload)
}

func (o *GetDataSourceNotFound) String() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceNotFound  %+v", 404, o.Payload)
}

func (o *GetDataSourceNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetDataSourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceInternalServerError creates a GetDataSourceInternalServerError with default headers values
func NewGetDataSourceInternalServerError() *GetDataSourceInternalServerError {
	return &GetDataSourceInternalServerError{}
}

/*
GetDataSourceInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetDataSourceInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get data source internal server error response has a 2xx status code
func (o *GetDataSourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source internal server error response has a 3xx status code
func (o *GetDataSourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source internal server error response has a 4xx status code
func (o *GetDataSourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data source internal server error response has a 5xx status code
func (o *GetDataSourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get data source internal server error response a status code equal to that given
func (o *GetDataSourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get data source internal server error response
func (o *GetDataSourceInternalServerError) Code() int {
	return 500
}

func (o *GetDataSourceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourceInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/datasource/{Name}][%d] getDataSourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourceInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetDataSourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}