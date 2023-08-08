// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// GetCellReader is a Reader for the GetCell structure.
type GetCellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCellUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCellForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCellNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCellInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /share/cell/{Uuid}] GetCell", response, response.Code())
	}
}

// NewGetCellOK creates a GetCellOK with default headers values
func NewGetCellOK() *GetCellOK {
	return &GetCellOK{}
}

/*
GetCellOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetCellOK struct {
	Payload *models.RestCell
}

// IsSuccess returns true when this get cell o k response has a 2xx status code
func (o *GetCellOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cell o k response has a 3xx status code
func (o *GetCellOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cell o k response has a 4xx status code
func (o *GetCellOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cell o k response has a 5xx status code
func (o *GetCellOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cell o k response a status code equal to that given
func (o *GetCellOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cell o k response
func (o *GetCellOK) Code() int {
	return 200
}

func (o *GetCellOK) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellOK  %+v", 200, o.Payload)
}

func (o *GetCellOK) String() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellOK  %+v", 200, o.Payload)
}

func (o *GetCellOK) GetPayload() *models.RestCell {
	return o.Payload
}

func (o *GetCellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCell)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCellUnauthorized creates a GetCellUnauthorized with default headers values
func NewGetCellUnauthorized() *GetCellUnauthorized {
	return &GetCellUnauthorized{}
}

/*
GetCellUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetCellUnauthorized struct {
}

// IsSuccess returns true when this get cell unauthorized response has a 2xx status code
func (o *GetCellUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cell unauthorized response has a 3xx status code
func (o *GetCellUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cell unauthorized response has a 4xx status code
func (o *GetCellUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cell unauthorized response has a 5xx status code
func (o *GetCellUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cell unauthorized response a status code equal to that given
func (o *GetCellUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cell unauthorized response
func (o *GetCellUnauthorized) Code() int {
	return 401
}

func (o *GetCellUnauthorized) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellUnauthorized ", 401)
}

func (o *GetCellUnauthorized) String() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellUnauthorized ", 401)
}

func (o *GetCellUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCellForbidden creates a GetCellForbidden with default headers values
func NewGetCellForbidden() *GetCellForbidden {
	return &GetCellForbidden{}
}

/*
GetCellForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetCellForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get cell forbidden response has a 2xx status code
func (o *GetCellForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cell forbidden response has a 3xx status code
func (o *GetCellForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cell forbidden response has a 4xx status code
func (o *GetCellForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cell forbidden response has a 5xx status code
func (o *GetCellForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cell forbidden response a status code equal to that given
func (o *GetCellForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cell forbidden response
func (o *GetCellForbidden) Code() int {
	return 403
}

func (o *GetCellForbidden) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellForbidden  %+v", 403, o.Payload)
}

func (o *GetCellForbidden) String() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellForbidden  %+v", 403, o.Payload)
}

func (o *GetCellForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetCellForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCellNotFound creates a GetCellNotFound with default headers values
func NewGetCellNotFound() *GetCellNotFound {
	return &GetCellNotFound{}
}

/*
GetCellNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetCellNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get cell not found response has a 2xx status code
func (o *GetCellNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cell not found response has a 3xx status code
func (o *GetCellNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cell not found response has a 4xx status code
func (o *GetCellNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cell not found response has a 5xx status code
func (o *GetCellNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cell not found response a status code equal to that given
func (o *GetCellNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cell not found response
func (o *GetCellNotFound) Code() int {
	return 404
}

func (o *GetCellNotFound) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellNotFound  %+v", 404, o.Payload)
}

func (o *GetCellNotFound) String() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellNotFound  %+v", 404, o.Payload)
}

func (o *GetCellNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetCellNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCellInternalServerError creates a GetCellInternalServerError with default headers values
func NewGetCellInternalServerError() *GetCellInternalServerError {
	return &GetCellInternalServerError{}
}

/*
GetCellInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetCellInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get cell internal server error response has a 2xx status code
func (o *GetCellInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cell internal server error response has a 3xx status code
func (o *GetCellInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cell internal server error response has a 4xx status code
func (o *GetCellInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cell internal server error response has a 5xx status code
func (o *GetCellInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cell internal server error response a status code equal to that given
func (o *GetCellInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cell internal server error response
func (o *GetCellInternalServerError) Code() int {
	return 500
}

func (o *GetCellInternalServerError) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCellInternalServerError) String() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCellInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetCellInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
