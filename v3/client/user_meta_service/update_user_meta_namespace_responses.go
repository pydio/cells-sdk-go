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

// UpdateUserMetaNamespaceReader is a Reader for the UpdateUserMetaNamespace structure.
type UpdateUserMetaNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserMetaNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserMetaNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserMetaNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserMetaNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserMetaNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserMetaNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserMetaNamespaceOK creates a UpdateUserMetaNamespaceOK with default headers values
func NewUpdateUserMetaNamespaceOK() *UpdateUserMetaNamespaceOK {
	return &UpdateUserMetaNamespaceOK{}
}

/* UpdateUserMetaNamespaceOK describes a response with status code 200, with default header values.

UpdateUserMetaNamespaceOK update user meta namespace o k
*/
type UpdateUserMetaNamespaceOK struct {
	Payload *models.IdmUpdateUserMetaNamespaceResponse
}

func (o *UpdateUserMetaNamespaceOK) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceOK  %+v", 200, o.Payload)
}
func (o *UpdateUserMetaNamespaceOK) GetPayload() *models.IdmUpdateUserMetaNamespaceResponse {
	return o.Payload
}

func (o *UpdateUserMetaNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmUpdateUserMetaNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserMetaNamespaceUnauthorized creates a UpdateUserMetaNamespaceUnauthorized with default headers values
func NewUpdateUserMetaNamespaceUnauthorized() *UpdateUserMetaNamespaceUnauthorized {
	return &UpdateUserMetaNamespaceUnauthorized{}
}

/* UpdateUserMetaNamespaceUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type UpdateUserMetaNamespaceUnauthorized struct {
	Payload *models.RestError
}

func (o *UpdateUserMetaNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateUserMetaNamespaceUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateUserMetaNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserMetaNamespaceForbidden creates a UpdateUserMetaNamespaceForbidden with default headers values
func NewUpdateUserMetaNamespaceForbidden() *UpdateUserMetaNamespaceForbidden {
	return &UpdateUserMetaNamespaceForbidden{}
}

/* UpdateUserMetaNamespaceForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type UpdateUserMetaNamespaceForbidden struct {
	Payload *models.RestError
}

func (o *UpdateUserMetaNamespaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserMetaNamespaceForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateUserMetaNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserMetaNamespaceNotFound creates a UpdateUserMetaNamespaceNotFound with default headers values
func NewUpdateUserMetaNamespaceNotFound() *UpdateUserMetaNamespaceNotFound {
	return &UpdateUserMetaNamespaceNotFound{}
}

/* UpdateUserMetaNamespaceNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type UpdateUserMetaNamespaceNotFound struct {
	Payload *models.RestError
}

func (o *UpdateUserMetaNamespaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUserMetaNamespaceNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateUserMetaNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserMetaNamespaceInternalServerError creates a UpdateUserMetaNamespaceInternalServerError with default headers values
func NewUpdateUserMetaNamespaceInternalServerError() *UpdateUserMetaNamespaceInternalServerError {
	return &UpdateUserMetaNamespaceInternalServerError{}
}

/* UpdateUserMetaNamespaceInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type UpdateUserMetaNamespaceInternalServerError struct {
	Payload *models.RestError
}

func (o *UpdateUserMetaNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user-meta/namespace][%d] updateUserMetaNamespaceInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateUserMetaNamespaceInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateUserMetaNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
