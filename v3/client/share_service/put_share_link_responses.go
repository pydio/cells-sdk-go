// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PutShareLinkReader is a Reader for the PutShareLink structure.
type PutShareLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutShareLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutShareLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutShareLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutShareLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutShareLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutShareLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutShareLinkOK creates a PutShareLinkOK with default headers values
func NewPutShareLinkOK() *PutShareLinkOK {
	return &PutShareLinkOK{}
}

/* PutShareLinkOK describes a response with status code 200, with default header values.

PutShareLinkOK put share link o k
*/
type PutShareLinkOK struct {
	Payload *models.RestShareLink
}

func (o *PutShareLinkOK) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkOK  %+v", 200, o.Payload)
}
func (o *PutShareLinkOK) GetPayload() *models.RestShareLink {
	return o.Payload
}

func (o *PutShareLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestShareLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkUnauthorized creates a PutShareLinkUnauthorized with default headers values
func NewPutShareLinkUnauthorized() *PutShareLinkUnauthorized {
	return &PutShareLinkUnauthorized{}
}

/* PutShareLinkUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutShareLinkUnauthorized struct {
	Payload *models.RestError
}

func (o *PutShareLinkUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkUnauthorized  %+v", 401, o.Payload)
}
func (o *PutShareLinkUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkForbidden creates a PutShareLinkForbidden with default headers values
func NewPutShareLinkForbidden() *PutShareLinkForbidden {
	return &PutShareLinkForbidden{}
}

/* PutShareLinkForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutShareLinkForbidden struct {
	Payload *models.RestError
}

func (o *PutShareLinkForbidden) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkForbidden  %+v", 403, o.Payload)
}
func (o *PutShareLinkForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkNotFound creates a PutShareLinkNotFound with default headers values
func NewPutShareLinkNotFound() *PutShareLinkNotFound {
	return &PutShareLinkNotFound{}
}

/* PutShareLinkNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutShareLinkNotFound struct {
	Payload *models.RestError
}

func (o *PutShareLinkNotFound) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkNotFound  %+v", 404, o.Payload)
}
func (o *PutShareLinkNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutShareLinkInternalServerError creates a PutShareLinkInternalServerError with default headers values
func NewPutShareLinkInternalServerError() *PutShareLinkInternalServerError {
	return &PutShareLinkInternalServerError{}
}

/* PutShareLinkInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutShareLinkInternalServerError struct {
	Payload *models.RestError
}

func (o *PutShareLinkInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /share/link][%d] putShareLinkInternalServerError  %+v", 500, o.Payload)
}
func (o *PutShareLinkInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutShareLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
