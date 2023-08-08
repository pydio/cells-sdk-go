// Code generated by go-swagger; DO NOT EDIT.

package policy_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// ListPoliciesReader is a Reader for the ListPolicies structure.
type ListPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /policy] ListPolicies", response, response.Code())
	}
}

// NewListPoliciesOK creates a ListPoliciesOK with default headers values
func NewListPoliciesOK() *ListPoliciesOK {
	return &ListPoliciesOK{}
}

/*
ListPoliciesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPoliciesOK struct {
	Payload *models.IdmListPolicyGroupsResponse
}

// IsSuccess returns true when this list policies o k response has a 2xx status code
func (o *ListPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list policies o k response has a 3xx status code
func (o *ListPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies o k response has a 4xx status code
func (o *ListPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list policies o k response has a 5xx status code
func (o *ListPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies o k response a status code equal to that given
func (o *ListPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list policies o k response
func (o *ListPoliciesOK) Code() int {
	return 200
}

func (o *ListPoliciesOK) Error() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListPoliciesOK) String() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListPoliciesOK) GetPayload() *models.IdmListPolicyGroupsResponse {
	return o.Payload
}

func (o *ListPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmListPolicyGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesUnauthorized creates a ListPoliciesUnauthorized with default headers values
func NewListPoliciesUnauthorized() *ListPoliciesUnauthorized {
	return &ListPoliciesUnauthorized{}
}

/*
ListPoliciesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListPoliciesUnauthorized struct {
}

// IsSuccess returns true when this list policies unauthorized response has a 2xx status code
func (o *ListPoliciesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies unauthorized response has a 3xx status code
func (o *ListPoliciesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies unauthorized response has a 4xx status code
func (o *ListPoliciesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies unauthorized response has a 5xx status code
func (o *ListPoliciesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies unauthorized response a status code equal to that given
func (o *ListPoliciesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list policies unauthorized response
func (o *ListPoliciesUnauthorized) Code() int {
	return 401
}

func (o *ListPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesUnauthorized ", 401)
}

func (o *ListPoliciesUnauthorized) String() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesUnauthorized ", 401)
}

func (o *ListPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPoliciesForbidden creates a ListPoliciesForbidden with default headers values
func NewListPoliciesForbidden() *ListPoliciesForbidden {
	return &ListPoliciesForbidden{}
}

/*
ListPoliciesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListPoliciesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list policies forbidden response has a 2xx status code
func (o *ListPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies forbidden response has a 3xx status code
func (o *ListPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies forbidden response has a 4xx status code
func (o *ListPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies forbidden response has a 5xx status code
func (o *ListPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies forbidden response a status code equal to that given
func (o *ListPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list policies forbidden response
func (o *ListPoliciesForbidden) Code() int {
	return 403
}

func (o *ListPoliciesForbidden) Error() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListPoliciesForbidden) String() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListPoliciesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesNotFound creates a ListPoliciesNotFound with default headers values
func NewListPoliciesNotFound() *ListPoliciesNotFound {
	return &ListPoliciesNotFound{}
}

/*
ListPoliciesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListPoliciesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list policies not found response has a 2xx status code
func (o *ListPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies not found response has a 3xx status code
func (o *ListPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies not found response has a 4xx status code
func (o *ListPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies not found response has a 5xx status code
func (o *ListPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies not found response a status code equal to that given
func (o *ListPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list policies not found response
func (o *ListPoliciesNotFound) Code() int {
	return 404
}

func (o *ListPoliciesNotFound) Error() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *ListPoliciesNotFound) String() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *ListPoliciesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesInternalServerError creates a ListPoliciesInternalServerError with default headers values
func NewListPoliciesInternalServerError() *ListPoliciesInternalServerError {
	return &ListPoliciesInternalServerError{}
}

/*
ListPoliciesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListPoliciesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list policies internal server error response has a 2xx status code
func (o *ListPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies internal server error response has a 3xx status code
func (o *ListPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies internal server error response has a 4xx status code
func (o *ListPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list policies internal server error response has a 5xx status code
func (o *ListPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list policies internal server error response a status code equal to that given
func (o *ListPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list policies internal server error response
func (o *ListPoliciesInternalServerError) Code() int {
	return 500
}

func (o *ListPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy][%d] listPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPoliciesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
