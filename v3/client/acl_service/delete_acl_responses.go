// Code generated by go-swagger; DO NOT EDIT.

package acl_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// DeleteACLReader is a Reader for the DeleteACL structure.
type DeleteACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteACLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteACLForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteACLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteACLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteACLOK creates a DeleteACLOK with default headers values
func NewDeleteACLOK() *DeleteACLOK {
	return &DeleteACLOK{}
}

/* DeleteACLOK describes a response with status code 200, with default header values.

DeleteACLOK delete Acl o k
*/
type DeleteACLOK struct {
	Payload *models.RestDeleteResponse
}

func (o *DeleteACLOK) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclOK  %+v", 200, o.Payload)
}
func (o *DeleteACLOK) GetPayload() *models.RestDeleteResponse {
	return o.Payload
}

func (o *DeleteACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteACLUnauthorized creates a DeleteACLUnauthorized with default headers values
func NewDeleteACLUnauthorized() *DeleteACLUnauthorized {
	return &DeleteACLUnauthorized{}
}

/* DeleteACLUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteACLUnauthorized struct {
	Payload *models.RestError
}

func (o *DeleteACLUnauthorized) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteACLUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteACLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteACLForbidden creates a DeleteACLForbidden with default headers values
func NewDeleteACLForbidden() *DeleteACLForbidden {
	return &DeleteACLForbidden{}
}

/* DeleteACLForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteACLForbidden struct {
	Payload *models.RestError
}

func (o *DeleteACLForbidden) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclForbidden  %+v", 403, o.Payload)
}
func (o *DeleteACLForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteACLForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteACLNotFound creates a DeleteACLNotFound with default headers values
func NewDeleteACLNotFound() *DeleteACLNotFound {
	return &DeleteACLNotFound{}
}

/* DeleteACLNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteACLNotFound struct {
	Payload *models.RestError
}

func (o *DeleteACLNotFound) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclNotFound  %+v", 404, o.Payload)
}
func (o *DeleteACLNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteACLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteACLInternalServerError creates a DeleteACLInternalServerError with default headers values
func NewDeleteACLInternalServerError() *DeleteACLInternalServerError {
	return &DeleteACLInternalServerError{}
}

/* DeleteACLInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteACLInternalServerError struct {
	Payload *models.RestError
}

func (o *DeleteACLInternalServerError) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteACLInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteACLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
