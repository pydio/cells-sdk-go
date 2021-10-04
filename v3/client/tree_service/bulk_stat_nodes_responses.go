// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// BulkStatNodesReader is a Reader for the BulkStatNodes structure.
type BulkStatNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkStatNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkStatNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBulkStatNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkStatNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkStatNodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkStatNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkStatNodesOK creates a BulkStatNodesOK with default headers values
func NewBulkStatNodesOK() *BulkStatNodesOK {
	return &BulkStatNodesOK{}
}

/* BulkStatNodesOK describes a response with status code 200, with default header values.

BulkStatNodesOK bulk stat nodes o k
*/
type BulkStatNodesOK struct {
	Payload *models.RestBulkMetaResponse
}

func (o *BulkStatNodesOK) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesOK  %+v", 200, o.Payload)
}
func (o *BulkStatNodesOK) GetPayload() *models.RestBulkMetaResponse {
	return o.Payload
}

func (o *BulkStatNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestBulkMetaResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkStatNodesUnauthorized creates a BulkStatNodesUnauthorized with default headers values
func NewBulkStatNodesUnauthorized() *BulkStatNodesUnauthorized {
	return &BulkStatNodesUnauthorized{}
}

/* BulkStatNodesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type BulkStatNodesUnauthorized struct {
	Payload *models.RestError
}

func (o *BulkStatNodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkStatNodesUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *BulkStatNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkStatNodesForbidden creates a BulkStatNodesForbidden with default headers values
func NewBulkStatNodesForbidden() *BulkStatNodesForbidden {
	return &BulkStatNodesForbidden{}
}

/* BulkStatNodesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type BulkStatNodesForbidden struct {
	Payload *models.RestError
}

func (o *BulkStatNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesForbidden  %+v", 403, o.Payload)
}
func (o *BulkStatNodesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *BulkStatNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkStatNodesNotFound creates a BulkStatNodesNotFound with default headers values
func NewBulkStatNodesNotFound() *BulkStatNodesNotFound {
	return &BulkStatNodesNotFound{}
}

/* BulkStatNodesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type BulkStatNodesNotFound struct {
	Payload *models.RestError
}

func (o *BulkStatNodesNotFound) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesNotFound  %+v", 404, o.Payload)
}
func (o *BulkStatNodesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *BulkStatNodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkStatNodesInternalServerError creates a BulkStatNodesInternalServerError with default headers values
func NewBulkStatNodesInternalServerError() *BulkStatNodesInternalServerError {
	return &BulkStatNodesInternalServerError{}
}

/* BulkStatNodesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type BulkStatNodesInternalServerError struct {
	Payload *models.RestError
}

func (o *BulkStatNodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkStatNodesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *BulkStatNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
