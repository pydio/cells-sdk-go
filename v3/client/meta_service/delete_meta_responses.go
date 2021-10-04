// Code generated by go-swagger; DO NOT EDIT.

package meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// DeleteMetaReader is a Reader for the DeleteMeta structure.
type DeleteMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMetaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMetaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMetaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMetaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMetaOK creates a DeleteMetaOK with default headers values
func NewDeleteMetaOK() *DeleteMetaOK {
	return &DeleteMetaOK{}
}

/* DeleteMetaOK describes a response with status code 200, with default header values.

DeleteMetaOK delete meta o k
*/
type DeleteMetaOK struct {
	Payload *models.TreeNode
}

func (o *DeleteMetaOK) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaOK  %+v", 200, o.Payload)
}
func (o *DeleteMetaOK) GetPayload() *models.TreeNode {
	return o.Payload
}

func (o *DeleteMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreeNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetaUnauthorized creates a DeleteMetaUnauthorized with default headers values
func NewDeleteMetaUnauthorized() *DeleteMetaUnauthorized {
	return &DeleteMetaUnauthorized{}
}

/* DeleteMetaUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteMetaUnauthorized struct {
	Payload *models.RestError
}

func (o *DeleteMetaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteMetaUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteMetaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetaForbidden creates a DeleteMetaForbidden with default headers values
func NewDeleteMetaForbidden() *DeleteMetaForbidden {
	return &DeleteMetaForbidden{}
}

/* DeleteMetaForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteMetaForbidden struct {
	Payload *models.RestError
}

func (o *DeleteMetaForbidden) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaForbidden  %+v", 403, o.Payload)
}
func (o *DeleteMetaForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteMetaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetaNotFound creates a DeleteMetaNotFound with default headers values
func NewDeleteMetaNotFound() *DeleteMetaNotFound {
	return &DeleteMetaNotFound{}
}

/* DeleteMetaNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteMetaNotFound struct {
	Payload *models.RestError
}

func (o *DeleteMetaNotFound) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaNotFound  %+v", 404, o.Payload)
}
func (o *DeleteMetaNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteMetaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetaInternalServerError creates a DeleteMetaInternalServerError with default headers values
func NewDeleteMetaInternalServerError() *DeleteMetaInternalServerError {
	return &DeleteMetaInternalServerError{}
}

/* DeleteMetaInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteMetaInternalServerError struct {
	Payload *models.RestError
}

func (o *DeleteMetaInternalServerError) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteMetaInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteMetaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
