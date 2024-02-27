// Code generated by go-swagger; DO NOT EDIT.

package acl_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
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
		return nil, runtime.NewAPIError("[POST /acl/bulk/delete] DeleteAcl", response, response.Code())
	}
}

// NewDeleteACLOK creates a DeleteACLOK with default headers values
func NewDeleteACLOK() *DeleteACLOK {
	return &DeleteACLOK{}
}

/*
DeleteACLOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteACLOK struct {
	Payload *models.RestDeleteResponse
}

// IsSuccess returns true when this delete Acl o k response has a 2xx status code
func (o *DeleteACLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Acl o k response has a 3xx status code
func (o *DeleteACLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl o k response has a 4xx status code
func (o *DeleteACLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Acl o k response has a 5xx status code
func (o *DeleteACLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl o k response a status code equal to that given
func (o *DeleteACLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete Acl o k response
func (o *DeleteACLOK) Code() int {
	return 200
}

func (o *DeleteACLOK) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclOK  %+v", 200, o.Payload)
}

func (o *DeleteACLOK) String() string {
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

/*
DeleteACLUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteACLUnauthorized struct {
}

// IsSuccess returns true when this delete Acl unauthorized response has a 2xx status code
func (o *DeleteACLUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Acl unauthorized response has a 3xx status code
func (o *DeleteACLUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl unauthorized response has a 4xx status code
func (o *DeleteACLUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Acl unauthorized response has a 5xx status code
func (o *DeleteACLUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl unauthorized response a status code equal to that given
func (o *DeleteACLUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete Acl unauthorized response
func (o *DeleteACLUnauthorized) Code() int {
	return 401
}

func (o *DeleteACLUnauthorized) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclUnauthorized ", 401)
}

func (o *DeleteACLUnauthorized) String() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclUnauthorized ", 401)
}

func (o *DeleteACLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteACLForbidden creates a DeleteACLForbidden with default headers values
func NewDeleteACLForbidden() *DeleteACLForbidden {
	return &DeleteACLForbidden{}
}

/*
DeleteACLForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteACLForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete Acl forbidden response has a 2xx status code
func (o *DeleteACLForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Acl forbidden response has a 3xx status code
func (o *DeleteACLForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl forbidden response has a 4xx status code
func (o *DeleteACLForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Acl forbidden response has a 5xx status code
func (o *DeleteACLForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl forbidden response a status code equal to that given
func (o *DeleteACLForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete Acl forbidden response
func (o *DeleteACLForbidden) Code() int {
	return 403
}

func (o *DeleteACLForbidden) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclForbidden  %+v", 403, o.Payload)
}

func (o *DeleteACLForbidden) String() string {
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

/*
DeleteACLNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteACLNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete Acl not found response has a 2xx status code
func (o *DeleteACLNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Acl not found response has a 3xx status code
func (o *DeleteACLNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl not found response has a 4xx status code
func (o *DeleteACLNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Acl not found response has a 5xx status code
func (o *DeleteACLNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl not found response a status code equal to that given
func (o *DeleteACLNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Acl not found response
func (o *DeleteACLNotFound) Code() int {
	return 404
}

func (o *DeleteACLNotFound) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclNotFound  %+v", 404, o.Payload)
}

func (o *DeleteACLNotFound) String() string {
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

/*
DeleteACLInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteACLInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete Acl internal server error response has a 2xx status code
func (o *DeleteACLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Acl internal server error response has a 3xx status code
func (o *DeleteACLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl internal server error response has a 4xx status code
func (o *DeleteACLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Acl internal server error response has a 5xx status code
func (o *DeleteACLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Acl internal server error response a status code equal to that given
func (o *DeleteACLInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Acl internal server error response
func (o *DeleteACLInternalServerError) Code() int {
	return 500
}

func (o *DeleteACLInternalServerError) Error() string {
	return fmt.Sprintf("[POST /acl/bulk/delete][%d] deleteAclInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteACLInternalServerError) String() string {
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
