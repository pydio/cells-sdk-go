// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// PutUserMetaTagReader is a Reader for the PutUserMetaTag structure.
type PutUserMetaTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserMetaTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserMetaTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutUserMetaTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserMetaTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserMetaTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserMetaTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserMetaTagOK creates a PutUserMetaTagOK with default headers values
func NewPutUserMetaTagOK() *PutUserMetaTagOK {
	return &PutUserMetaTagOK{}
}

/* PutUserMetaTagOK describes a response with status code 200, with default header values.

PutUserMetaTagOK put user meta tag o k
*/
type PutUserMetaTagOK struct {
	Payload *models.RestPutUserMetaTagResponse
}

func (o *PutUserMetaTagOK) Error() string {
	return fmt.Sprintf("[POST /user-meta/tags/{Namespace}][%d] putUserMetaTagOK  %+v", 200, o.Payload)
}
func (o *PutUserMetaTagOK) GetPayload() *models.RestPutUserMetaTagResponse {
	return o.Payload
}

func (o *PutUserMetaTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestPutUserMetaTagResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserMetaTagUnauthorized creates a PutUserMetaTagUnauthorized with default headers values
func NewPutUserMetaTagUnauthorized() *PutUserMetaTagUnauthorized {
	return &PutUserMetaTagUnauthorized{}
}

/* PutUserMetaTagUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutUserMetaTagUnauthorized struct {
}

func (o *PutUserMetaTagUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user-meta/tags/{Namespace}][%d] putUserMetaTagUnauthorized ", 401)
}

func (o *PutUserMetaTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUserMetaTagForbidden creates a PutUserMetaTagForbidden with default headers values
func NewPutUserMetaTagForbidden() *PutUserMetaTagForbidden {
	return &PutUserMetaTagForbidden{}
}

/* PutUserMetaTagForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutUserMetaTagForbidden struct {
	Payload *models.RestError
}

func (o *PutUserMetaTagForbidden) Error() string {
	return fmt.Sprintf("[POST /user-meta/tags/{Namespace}][%d] putUserMetaTagForbidden  %+v", 403, o.Payload)
}
func (o *PutUserMetaTagForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserMetaTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserMetaTagNotFound creates a PutUserMetaTagNotFound with default headers values
func NewPutUserMetaTagNotFound() *PutUserMetaTagNotFound {
	return &PutUserMetaTagNotFound{}
}

/* PutUserMetaTagNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutUserMetaTagNotFound struct {
	Payload *models.RestError
}

func (o *PutUserMetaTagNotFound) Error() string {
	return fmt.Sprintf("[POST /user-meta/tags/{Namespace}][%d] putUserMetaTagNotFound  %+v", 404, o.Payload)
}
func (o *PutUserMetaTagNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserMetaTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserMetaTagInternalServerError creates a PutUserMetaTagInternalServerError with default headers values
func NewPutUserMetaTagInternalServerError() *PutUserMetaTagInternalServerError {
	return &PutUserMetaTagInternalServerError{}
}

/* PutUserMetaTagInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutUserMetaTagInternalServerError struct {
	Payload *models.RestError
}

func (o *PutUserMetaTagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user-meta/tags/{Namespace}][%d] putUserMetaTagInternalServerError  %+v", 500, o.Payload)
}
func (o *PutUserMetaTagInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutUserMetaTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
