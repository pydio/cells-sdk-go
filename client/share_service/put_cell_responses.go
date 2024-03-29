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

// PutCellReader is a Reader for the PutCell structure.
type PutCellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutCellUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCellForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCellNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCellInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /share/cell] PutCell", response, response.Code())
	}
}

// NewPutCellOK creates a PutCellOK with default headers values
func NewPutCellOK() *PutCellOK {
	return &PutCellOK{}
}

/*
PutCellOK describes a response with status code 200, with default header values.

A successful response.
*/
type PutCellOK struct {
	Payload *models.RestCell
}

// IsSuccess returns true when this put cell o k response has a 2xx status code
func (o *PutCellOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put cell o k response has a 3xx status code
func (o *PutCellOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cell o k response has a 4xx status code
func (o *PutCellOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put cell o k response has a 5xx status code
func (o *PutCellOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put cell o k response a status code equal to that given
func (o *PutCellOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put cell o k response
func (o *PutCellOK) Code() int {
	return 200
}

func (o *PutCellOK) Error() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellOK  %+v", 200, o.Payload)
}

func (o *PutCellOK) String() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellOK  %+v", 200, o.Payload)
}

func (o *PutCellOK) GetPayload() *models.RestCell {
	return o.Payload
}

func (o *PutCellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCell)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCellUnauthorized creates a PutCellUnauthorized with default headers values
func NewPutCellUnauthorized() *PutCellUnauthorized {
	return &PutCellUnauthorized{}
}

/*
PutCellUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutCellUnauthorized struct {
}

// IsSuccess returns true when this put cell unauthorized response has a 2xx status code
func (o *PutCellUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cell unauthorized response has a 3xx status code
func (o *PutCellUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cell unauthorized response has a 4xx status code
func (o *PutCellUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put cell unauthorized response has a 5xx status code
func (o *PutCellUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put cell unauthorized response a status code equal to that given
func (o *PutCellUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put cell unauthorized response
func (o *PutCellUnauthorized) Code() int {
	return 401
}

func (o *PutCellUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellUnauthorized ", 401)
}

func (o *PutCellUnauthorized) String() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellUnauthorized ", 401)
}

func (o *PutCellUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCellForbidden creates a PutCellForbidden with default headers values
func NewPutCellForbidden() *PutCellForbidden {
	return &PutCellForbidden{}
}

/*
PutCellForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutCellForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put cell forbidden response has a 2xx status code
func (o *PutCellForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cell forbidden response has a 3xx status code
func (o *PutCellForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cell forbidden response has a 4xx status code
func (o *PutCellForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put cell forbidden response has a 5xx status code
func (o *PutCellForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put cell forbidden response a status code equal to that given
func (o *PutCellForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put cell forbidden response
func (o *PutCellForbidden) Code() int {
	return 403
}

func (o *PutCellForbidden) Error() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellForbidden  %+v", 403, o.Payload)
}

func (o *PutCellForbidden) String() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellForbidden  %+v", 403, o.Payload)
}

func (o *PutCellForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutCellForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCellNotFound creates a PutCellNotFound with default headers values
func NewPutCellNotFound() *PutCellNotFound {
	return &PutCellNotFound{}
}

/*
PutCellNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutCellNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put cell not found response has a 2xx status code
func (o *PutCellNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cell not found response has a 3xx status code
func (o *PutCellNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cell not found response has a 4xx status code
func (o *PutCellNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put cell not found response has a 5xx status code
func (o *PutCellNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put cell not found response a status code equal to that given
func (o *PutCellNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put cell not found response
func (o *PutCellNotFound) Code() int {
	return 404
}

func (o *PutCellNotFound) Error() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellNotFound  %+v", 404, o.Payload)
}

func (o *PutCellNotFound) String() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellNotFound  %+v", 404, o.Payload)
}

func (o *PutCellNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutCellNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCellInternalServerError creates a PutCellInternalServerError with default headers values
func NewPutCellInternalServerError() *PutCellInternalServerError {
	return &PutCellInternalServerError{}
}

/*
PutCellInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutCellInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put cell internal server error response has a 2xx status code
func (o *PutCellInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cell internal server error response has a 3xx status code
func (o *PutCellInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cell internal server error response has a 4xx status code
func (o *PutCellInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put cell internal server error response has a 5xx status code
func (o *PutCellInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put cell internal server error response a status code equal to that given
func (o *PutCellInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put cell internal server error response
func (o *PutCellInternalServerError) Code() int {
	return 500
}

func (o *PutCellInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCellInternalServerError) String() string {
	return fmt.Sprintf("[PUT /share/cell][%d] putCellInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCellInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutCellInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
