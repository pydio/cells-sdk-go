// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// UpdateSharePoliciesReader is a Reader for the UpdateSharePolicies structure.
type UpdateSharePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSharePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSharePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSharePoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSharePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSharePoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSharePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /share/policies] UpdateSharePolicies", response, response.Code())
	}
}

// NewUpdateSharePoliciesOK creates a UpdateSharePoliciesOK with default headers values
func NewUpdateSharePoliciesOK() *UpdateSharePoliciesOK {
	return &UpdateSharePoliciesOK{}
}

/*
UpdateSharePoliciesOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSharePoliciesOK struct {
	Payload *models.RestUpdateSharePoliciesResponse
}

// IsSuccess returns true when this update share policies o k response has a 2xx status code
func (o *UpdateSharePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update share policies o k response has a 3xx status code
func (o *UpdateSharePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update share policies o k response has a 4xx status code
func (o *UpdateSharePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update share policies o k response has a 5xx status code
func (o *UpdateSharePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update share policies o k response a status code equal to that given
func (o *UpdateSharePoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update share policies o k response
func (o *UpdateSharePoliciesOK) Code() int {
	return 200
}

func (o *UpdateSharePoliciesOK) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdateSharePoliciesOK) String() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdateSharePoliciesOK) GetPayload() *models.RestUpdateSharePoliciesResponse {
	return o.Payload
}

func (o *UpdateSharePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUpdateSharePoliciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSharePoliciesUnauthorized creates a UpdateSharePoliciesUnauthorized with default headers values
func NewUpdateSharePoliciesUnauthorized() *UpdateSharePoliciesUnauthorized {
	return &UpdateSharePoliciesUnauthorized{}
}

/*
UpdateSharePoliciesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type UpdateSharePoliciesUnauthorized struct {
}

// IsSuccess returns true when this update share policies unauthorized response has a 2xx status code
func (o *UpdateSharePoliciesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update share policies unauthorized response has a 3xx status code
func (o *UpdateSharePoliciesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update share policies unauthorized response has a 4xx status code
func (o *UpdateSharePoliciesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update share policies unauthorized response has a 5xx status code
func (o *UpdateSharePoliciesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update share policies unauthorized response a status code equal to that given
func (o *UpdateSharePoliciesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update share policies unauthorized response
func (o *UpdateSharePoliciesUnauthorized) Code() int {
	return 401
}

func (o *UpdateSharePoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesUnauthorized ", 401)
}

func (o *UpdateSharePoliciesUnauthorized) String() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesUnauthorized ", 401)
}

func (o *UpdateSharePoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSharePoliciesForbidden creates a UpdateSharePoliciesForbidden with default headers values
func NewUpdateSharePoliciesForbidden() *UpdateSharePoliciesForbidden {
	return &UpdateSharePoliciesForbidden{}
}

/*
UpdateSharePoliciesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type UpdateSharePoliciesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update share policies forbidden response has a 2xx status code
func (o *UpdateSharePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update share policies forbidden response has a 3xx status code
func (o *UpdateSharePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update share policies forbidden response has a 4xx status code
func (o *UpdateSharePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update share policies forbidden response has a 5xx status code
func (o *UpdateSharePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update share policies forbidden response a status code equal to that given
func (o *UpdateSharePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update share policies forbidden response
func (o *UpdateSharePoliciesForbidden) Code() int {
	return 403
}

func (o *UpdateSharePoliciesForbidden) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *UpdateSharePoliciesForbidden) String() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *UpdateSharePoliciesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateSharePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSharePoliciesNotFound creates a UpdateSharePoliciesNotFound with default headers values
func NewUpdateSharePoliciesNotFound() *UpdateSharePoliciesNotFound {
	return &UpdateSharePoliciesNotFound{}
}

/*
UpdateSharePoliciesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type UpdateSharePoliciesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update share policies not found response has a 2xx status code
func (o *UpdateSharePoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update share policies not found response has a 3xx status code
func (o *UpdateSharePoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update share policies not found response has a 4xx status code
func (o *UpdateSharePoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update share policies not found response has a 5xx status code
func (o *UpdateSharePoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update share policies not found response a status code equal to that given
func (o *UpdateSharePoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update share policies not found response
func (o *UpdateSharePoliciesNotFound) Code() int {
	return 404
}

func (o *UpdateSharePoliciesNotFound) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSharePoliciesNotFound) String() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSharePoliciesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateSharePoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSharePoliciesInternalServerError creates a UpdateSharePoliciesInternalServerError with default headers values
func NewUpdateSharePoliciesInternalServerError() *UpdateSharePoliciesInternalServerError {
	return &UpdateSharePoliciesInternalServerError{}
}

/*
UpdateSharePoliciesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type UpdateSharePoliciesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this update share policies internal server error response has a 2xx status code
func (o *UpdateSharePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update share policies internal server error response has a 3xx status code
func (o *UpdateSharePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update share policies internal server error response has a 4xx status code
func (o *UpdateSharePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update share policies internal server error response has a 5xx status code
func (o *UpdateSharePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update share policies internal server error response a status code equal to that given
func (o *UpdateSharePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update share policies internal server error response
func (o *UpdateSharePoliciesInternalServerError) Code() int {
	return 500
}

func (o *UpdateSharePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSharePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSharePoliciesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UpdateSharePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
