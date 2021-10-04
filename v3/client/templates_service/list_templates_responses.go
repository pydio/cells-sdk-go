// Code generated by go-swagger; DO NOT EDIT.

package templates_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// ListTemplatesReader is a Reader for the ListTemplates structure.
type ListTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListTemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTemplatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTemplatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTemplatesOK creates a ListTemplatesOK with default headers values
func NewListTemplatesOK() *ListTemplatesOK {
	return &ListTemplatesOK{}
}

/* ListTemplatesOK describes a response with status code 200, with default header values.

ListTemplatesOK list templates o k
*/
type ListTemplatesOK struct {
	Payload *models.RestListTemplatesResponse
}

func (o *ListTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /templates][%d] listTemplatesOK  %+v", 200, o.Payload)
}
func (o *ListTemplatesOK) GetPayload() *models.RestListTemplatesResponse {
	return o.Payload
}

func (o *ListTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListTemplatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTemplatesUnauthorized creates a ListTemplatesUnauthorized with default headers values
func NewListTemplatesUnauthorized() *ListTemplatesUnauthorized {
	return &ListTemplatesUnauthorized{}
}

/* ListTemplatesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListTemplatesUnauthorized struct {
	Payload *models.RestError
}

func (o *ListTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /templates][%d] listTemplatesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListTemplatesUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTemplatesForbidden creates a ListTemplatesForbidden with default headers values
func NewListTemplatesForbidden() *ListTemplatesForbidden {
	return &ListTemplatesForbidden{}
}

/* ListTemplatesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListTemplatesForbidden struct {
	Payload *models.RestError
}

func (o *ListTemplatesForbidden) Error() string {
	return fmt.Sprintf("[GET /templates][%d] listTemplatesForbidden  %+v", 403, o.Payload)
}
func (o *ListTemplatesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTemplatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTemplatesNotFound creates a ListTemplatesNotFound with default headers values
func NewListTemplatesNotFound() *ListTemplatesNotFound {
	return &ListTemplatesNotFound{}
}

/* ListTemplatesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListTemplatesNotFound struct {
	Payload *models.RestError
}

func (o *ListTemplatesNotFound) Error() string {
	return fmt.Sprintf("[GET /templates][%d] listTemplatesNotFound  %+v", 404, o.Payload)
}
func (o *ListTemplatesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTemplatesInternalServerError creates a ListTemplatesInternalServerError with default headers values
func NewListTemplatesInternalServerError() *ListTemplatesInternalServerError {
	return &ListTemplatesInternalServerError{}
}

/* ListTemplatesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListTemplatesInternalServerError struct {
	Payload *models.RestError
}

func (o *ListTemplatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /templates][%d] listTemplatesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListTemplatesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTemplatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
