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

// DeleteShareLinkReader is a Reader for the DeleteShareLink structure.
type DeleteShareLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteShareLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteShareLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteShareLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteShareLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteShareLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteShareLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteShareLinkOK creates a DeleteShareLinkOK with default headers values
func NewDeleteShareLinkOK() *DeleteShareLinkOK {
	return &DeleteShareLinkOK{}
}

/*
DeleteShareLinkOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteShareLinkOK struct {
	Payload *models.RestDeleteShareLinkResponse
}

// IsSuccess returns true when this delete share link o k response has a 2xx status code
func (o *DeleteShareLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete share link o k response has a 3xx status code
func (o *DeleteShareLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete share link o k response has a 4xx status code
func (o *DeleteShareLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete share link o k response has a 5xx status code
func (o *DeleteShareLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete share link o k response a status code equal to that given
func (o *DeleteShareLinkOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteShareLinkOK) Error() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkOK  %+v", 200, o.Payload)
}

func (o *DeleteShareLinkOK) String() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkOK  %+v", 200, o.Payload)
}

func (o *DeleteShareLinkOK) GetPayload() *models.RestDeleteShareLinkResponse {
	return o.Payload
}

func (o *DeleteShareLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteShareLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteShareLinkUnauthorized creates a DeleteShareLinkUnauthorized with default headers values
func NewDeleteShareLinkUnauthorized() *DeleteShareLinkUnauthorized {
	return &DeleteShareLinkUnauthorized{}
}

/*
DeleteShareLinkUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteShareLinkUnauthorized struct {
}

// IsSuccess returns true when this delete share link unauthorized response has a 2xx status code
func (o *DeleteShareLinkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete share link unauthorized response has a 3xx status code
func (o *DeleteShareLinkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete share link unauthorized response has a 4xx status code
func (o *DeleteShareLinkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete share link unauthorized response has a 5xx status code
func (o *DeleteShareLinkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete share link unauthorized response a status code equal to that given
func (o *DeleteShareLinkUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteShareLinkUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkUnauthorized ", 401)
}

func (o *DeleteShareLinkUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkUnauthorized ", 401)
}

func (o *DeleteShareLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteShareLinkForbidden creates a DeleteShareLinkForbidden with default headers values
func NewDeleteShareLinkForbidden() *DeleteShareLinkForbidden {
	return &DeleteShareLinkForbidden{}
}

/*
DeleteShareLinkForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteShareLinkForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete share link forbidden response has a 2xx status code
func (o *DeleteShareLinkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete share link forbidden response has a 3xx status code
func (o *DeleteShareLinkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete share link forbidden response has a 4xx status code
func (o *DeleteShareLinkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete share link forbidden response has a 5xx status code
func (o *DeleteShareLinkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete share link forbidden response a status code equal to that given
func (o *DeleteShareLinkForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteShareLinkForbidden) Error() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkForbidden  %+v", 403, o.Payload)
}

func (o *DeleteShareLinkForbidden) String() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkForbidden  %+v", 403, o.Payload)
}

func (o *DeleteShareLinkForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteShareLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteShareLinkNotFound creates a DeleteShareLinkNotFound with default headers values
func NewDeleteShareLinkNotFound() *DeleteShareLinkNotFound {
	return &DeleteShareLinkNotFound{}
}

/*
DeleteShareLinkNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteShareLinkNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete share link not found response has a 2xx status code
func (o *DeleteShareLinkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete share link not found response has a 3xx status code
func (o *DeleteShareLinkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete share link not found response has a 4xx status code
func (o *DeleteShareLinkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete share link not found response has a 5xx status code
func (o *DeleteShareLinkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete share link not found response a status code equal to that given
func (o *DeleteShareLinkNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteShareLinkNotFound) Error() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkNotFound  %+v", 404, o.Payload)
}

func (o *DeleteShareLinkNotFound) String() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkNotFound  %+v", 404, o.Payload)
}

func (o *DeleteShareLinkNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteShareLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteShareLinkInternalServerError creates a DeleteShareLinkInternalServerError with default headers values
func NewDeleteShareLinkInternalServerError() *DeleteShareLinkInternalServerError {
	return &DeleteShareLinkInternalServerError{}
}

/*
DeleteShareLinkInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteShareLinkInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete share link internal server error response has a 2xx status code
func (o *DeleteShareLinkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete share link internal server error response has a 3xx status code
func (o *DeleteShareLinkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete share link internal server error response has a 4xx status code
func (o *DeleteShareLinkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete share link internal server error response has a 5xx status code
func (o *DeleteShareLinkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete share link internal server error response a status code equal to that given
func (o *DeleteShareLinkInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteShareLinkInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteShareLinkInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /share/link/{Uuid}][%d] deleteShareLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteShareLinkInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteShareLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}