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

// DeleteCellReader is a Reader for the DeleteCell structure.
type DeleteCellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCellUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCellForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCellNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCellInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCellOK creates a DeleteCellOK with default headers values
func NewDeleteCellOK() *DeleteCellOK {
	return &DeleteCellOK{}
}

/*
DeleteCellOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteCellOK struct {
	Payload *models.RestDeleteCellResponse
}

// IsSuccess returns true when this delete cell o k response has a 2xx status code
func (o *DeleteCellOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cell o k response has a 3xx status code
func (o *DeleteCellOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cell o k response has a 4xx status code
func (o *DeleteCellOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cell o k response has a 5xx status code
func (o *DeleteCellOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cell o k response a status code equal to that given
func (o *DeleteCellOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteCellOK) Error() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellOK  %+v", 200, o.Payload)
}

func (o *DeleteCellOK) String() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellOK  %+v", 200, o.Payload)
}

func (o *DeleteCellOK) GetPayload() *models.RestDeleteCellResponse {
	return o.Payload
}

func (o *DeleteCellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteCellResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCellUnauthorized creates a DeleteCellUnauthorized with default headers values
func NewDeleteCellUnauthorized() *DeleteCellUnauthorized {
	return &DeleteCellUnauthorized{}
}

/*
DeleteCellUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteCellUnauthorized struct {
}

// IsSuccess returns true when this delete cell unauthorized response has a 2xx status code
func (o *DeleteCellUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cell unauthorized response has a 3xx status code
func (o *DeleteCellUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cell unauthorized response has a 4xx status code
func (o *DeleteCellUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cell unauthorized response has a 5xx status code
func (o *DeleteCellUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cell unauthorized response a status code equal to that given
func (o *DeleteCellUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteCellUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellUnauthorized ", 401)
}

func (o *DeleteCellUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellUnauthorized ", 401)
}

func (o *DeleteCellUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCellForbidden creates a DeleteCellForbidden with default headers values
func NewDeleteCellForbidden() *DeleteCellForbidden {
	return &DeleteCellForbidden{}
}

/*
DeleteCellForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteCellForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete cell forbidden response has a 2xx status code
func (o *DeleteCellForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cell forbidden response has a 3xx status code
func (o *DeleteCellForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cell forbidden response has a 4xx status code
func (o *DeleteCellForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cell forbidden response has a 5xx status code
func (o *DeleteCellForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cell forbidden response a status code equal to that given
func (o *DeleteCellForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteCellForbidden) Error() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCellForbidden) String() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCellForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteCellForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCellNotFound creates a DeleteCellNotFound with default headers values
func NewDeleteCellNotFound() *DeleteCellNotFound {
	return &DeleteCellNotFound{}
}

/*
DeleteCellNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteCellNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete cell not found response has a 2xx status code
func (o *DeleteCellNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cell not found response has a 3xx status code
func (o *DeleteCellNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cell not found response has a 4xx status code
func (o *DeleteCellNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cell not found response has a 5xx status code
func (o *DeleteCellNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cell not found response a status code equal to that given
func (o *DeleteCellNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteCellNotFound) Error() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCellNotFound) String() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCellNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteCellNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCellInternalServerError creates a DeleteCellInternalServerError with default headers values
func NewDeleteCellInternalServerError() *DeleteCellInternalServerError {
	return &DeleteCellInternalServerError{}
}

/*
DeleteCellInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteCellInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete cell internal server error response has a 2xx status code
func (o *DeleteCellInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cell internal server error response has a 3xx status code
func (o *DeleteCellInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cell internal server error response has a 4xx status code
func (o *DeleteCellInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cell internal server error response has a 5xx status code
func (o *DeleteCellInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete cell internal server error response a status code equal to that given
func (o *DeleteCellInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteCellInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCellInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /share/cell/{Uuid}][%d] deleteCellInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCellInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteCellInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
