// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// SearchWorkspacesReader is a Reader for the SearchWorkspaces structure.
type SearchWorkspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchWorkspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchWorkspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchWorkspacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchWorkspacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchWorkspacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchWorkspacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workspace] SearchWorkspaces", response, response.Code())
	}
}

// NewSearchWorkspacesOK creates a SearchWorkspacesOK with default headers values
func NewSearchWorkspacesOK() *SearchWorkspacesOK {
	return &SearchWorkspacesOK{}
}

/*
SearchWorkspacesOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchWorkspacesOK struct {
	Payload *models.RestWorkspaceCollection
}

// IsSuccess returns true when this search workspaces o k response has a 2xx status code
func (o *SearchWorkspacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search workspaces o k response has a 3xx status code
func (o *SearchWorkspacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search workspaces o k response has a 4xx status code
func (o *SearchWorkspacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search workspaces o k response has a 5xx status code
func (o *SearchWorkspacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search workspaces o k response a status code equal to that given
func (o *SearchWorkspacesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search workspaces o k response
func (o *SearchWorkspacesOK) Code() int {
	return 200
}

func (o *SearchWorkspacesOK) Error() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesOK  %+v", 200, o.Payload)
}

func (o *SearchWorkspacesOK) String() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesOK  %+v", 200, o.Payload)
}

func (o *SearchWorkspacesOK) GetPayload() *models.RestWorkspaceCollection {
	return o.Payload
}

func (o *SearchWorkspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestWorkspaceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchWorkspacesUnauthorized creates a SearchWorkspacesUnauthorized with default headers values
func NewSearchWorkspacesUnauthorized() *SearchWorkspacesUnauthorized {
	return &SearchWorkspacesUnauthorized{}
}

/*
SearchWorkspacesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SearchWorkspacesUnauthorized struct {
}

// IsSuccess returns true when this search workspaces unauthorized response has a 2xx status code
func (o *SearchWorkspacesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search workspaces unauthorized response has a 3xx status code
func (o *SearchWorkspacesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search workspaces unauthorized response has a 4xx status code
func (o *SearchWorkspacesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search workspaces unauthorized response has a 5xx status code
func (o *SearchWorkspacesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search workspaces unauthorized response a status code equal to that given
func (o *SearchWorkspacesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the search workspaces unauthorized response
func (o *SearchWorkspacesUnauthorized) Code() int {
	return 401
}

func (o *SearchWorkspacesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesUnauthorized ", 401)
}

func (o *SearchWorkspacesUnauthorized) String() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesUnauthorized ", 401)
}

func (o *SearchWorkspacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchWorkspacesForbidden creates a SearchWorkspacesForbidden with default headers values
func NewSearchWorkspacesForbidden() *SearchWorkspacesForbidden {
	return &SearchWorkspacesForbidden{}
}

/*
SearchWorkspacesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SearchWorkspacesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search workspaces forbidden response has a 2xx status code
func (o *SearchWorkspacesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search workspaces forbidden response has a 3xx status code
func (o *SearchWorkspacesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search workspaces forbidden response has a 4xx status code
func (o *SearchWorkspacesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search workspaces forbidden response has a 5xx status code
func (o *SearchWorkspacesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search workspaces forbidden response a status code equal to that given
func (o *SearchWorkspacesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search workspaces forbidden response
func (o *SearchWorkspacesForbidden) Code() int {
	return 403
}

func (o *SearchWorkspacesForbidden) Error() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesForbidden  %+v", 403, o.Payload)
}

func (o *SearchWorkspacesForbidden) String() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesForbidden  %+v", 403, o.Payload)
}

func (o *SearchWorkspacesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchWorkspacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchWorkspacesNotFound creates a SearchWorkspacesNotFound with default headers values
func NewSearchWorkspacesNotFound() *SearchWorkspacesNotFound {
	return &SearchWorkspacesNotFound{}
}

/*
SearchWorkspacesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SearchWorkspacesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search workspaces not found response has a 2xx status code
func (o *SearchWorkspacesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search workspaces not found response has a 3xx status code
func (o *SearchWorkspacesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search workspaces not found response has a 4xx status code
func (o *SearchWorkspacesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search workspaces not found response has a 5xx status code
func (o *SearchWorkspacesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search workspaces not found response a status code equal to that given
func (o *SearchWorkspacesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search workspaces not found response
func (o *SearchWorkspacesNotFound) Code() int {
	return 404
}

func (o *SearchWorkspacesNotFound) Error() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesNotFound  %+v", 404, o.Payload)
}

func (o *SearchWorkspacesNotFound) String() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesNotFound  %+v", 404, o.Payload)
}

func (o *SearchWorkspacesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchWorkspacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchWorkspacesInternalServerError creates a SearchWorkspacesInternalServerError with default headers values
func NewSearchWorkspacesInternalServerError() *SearchWorkspacesInternalServerError {
	return &SearchWorkspacesInternalServerError{}
}

/*
SearchWorkspacesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SearchWorkspacesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search workspaces internal server error response has a 2xx status code
func (o *SearchWorkspacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search workspaces internal server error response has a 3xx status code
func (o *SearchWorkspacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search workspaces internal server error response has a 4xx status code
func (o *SearchWorkspacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search workspaces internal server error response has a 5xx status code
func (o *SearchWorkspacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search workspaces internal server error response a status code equal to that given
func (o *SearchWorkspacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search workspaces internal server error response
func (o *SearchWorkspacesInternalServerError) Code() int {
	return 500
}

func (o *SearchWorkspacesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchWorkspacesInternalServerError) String() string {
	return fmt.Sprintf("[POST /workspace][%d] searchWorkspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchWorkspacesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchWorkspacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
