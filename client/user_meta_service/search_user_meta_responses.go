// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// SearchUserMetaReader is a Reader for the SearchUserMeta structure.
type SearchUserMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUserMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUserMetaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchUserMetaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUserMetaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUserMetaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user-meta/search] SearchUserMeta", response, response.Code())
	}
}

// NewSearchUserMetaOK creates a SearchUserMetaOK with default headers values
func NewSearchUserMetaOK() *SearchUserMetaOK {
	return &SearchUserMetaOK{}
}

/*
SearchUserMetaOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchUserMetaOK struct {
	Payload *models.RestUserMetaCollection
}

// IsSuccess returns true when this search user meta o k response has a 2xx status code
func (o *SearchUserMetaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search user meta o k response has a 3xx status code
func (o *SearchUserMetaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user meta o k response has a 4xx status code
func (o *SearchUserMetaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search user meta o k response has a 5xx status code
func (o *SearchUserMetaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search user meta o k response a status code equal to that given
func (o *SearchUserMetaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search user meta o k response
func (o *SearchUserMetaOK) Code() int {
	return 200
}

func (o *SearchUserMetaOK) Error() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaOK  %+v", 200, o.Payload)
}

func (o *SearchUserMetaOK) String() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaOK  %+v", 200, o.Payload)
}

func (o *SearchUserMetaOK) GetPayload() *models.RestUserMetaCollection {
	return o.Payload
}

func (o *SearchUserMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUserMetaCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserMetaUnauthorized creates a SearchUserMetaUnauthorized with default headers values
func NewSearchUserMetaUnauthorized() *SearchUserMetaUnauthorized {
	return &SearchUserMetaUnauthorized{}
}

/*
SearchUserMetaUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SearchUserMetaUnauthorized struct {
}

// IsSuccess returns true when this search user meta unauthorized response has a 2xx status code
func (o *SearchUserMetaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user meta unauthorized response has a 3xx status code
func (o *SearchUserMetaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user meta unauthorized response has a 4xx status code
func (o *SearchUserMetaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search user meta unauthorized response has a 5xx status code
func (o *SearchUserMetaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search user meta unauthorized response a status code equal to that given
func (o *SearchUserMetaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the search user meta unauthorized response
func (o *SearchUserMetaUnauthorized) Code() int {
	return 401
}

func (o *SearchUserMetaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaUnauthorized ", 401)
}

func (o *SearchUserMetaUnauthorized) String() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaUnauthorized ", 401)
}

func (o *SearchUserMetaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchUserMetaForbidden creates a SearchUserMetaForbidden with default headers values
func NewSearchUserMetaForbidden() *SearchUserMetaForbidden {
	return &SearchUserMetaForbidden{}
}

/*
SearchUserMetaForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SearchUserMetaForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search user meta forbidden response has a 2xx status code
func (o *SearchUserMetaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user meta forbidden response has a 3xx status code
func (o *SearchUserMetaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user meta forbidden response has a 4xx status code
func (o *SearchUserMetaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search user meta forbidden response has a 5xx status code
func (o *SearchUserMetaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search user meta forbidden response a status code equal to that given
func (o *SearchUserMetaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search user meta forbidden response
func (o *SearchUserMetaForbidden) Code() int {
	return 403
}

func (o *SearchUserMetaForbidden) Error() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaForbidden  %+v", 403, o.Payload)
}

func (o *SearchUserMetaForbidden) String() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaForbidden  %+v", 403, o.Payload)
}

func (o *SearchUserMetaForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchUserMetaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserMetaNotFound creates a SearchUserMetaNotFound with default headers values
func NewSearchUserMetaNotFound() *SearchUserMetaNotFound {
	return &SearchUserMetaNotFound{}
}

/*
SearchUserMetaNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SearchUserMetaNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search user meta not found response has a 2xx status code
func (o *SearchUserMetaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user meta not found response has a 3xx status code
func (o *SearchUserMetaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user meta not found response has a 4xx status code
func (o *SearchUserMetaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search user meta not found response has a 5xx status code
func (o *SearchUserMetaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search user meta not found response a status code equal to that given
func (o *SearchUserMetaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search user meta not found response
func (o *SearchUserMetaNotFound) Code() int {
	return 404
}

func (o *SearchUserMetaNotFound) Error() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaNotFound  %+v", 404, o.Payload)
}

func (o *SearchUserMetaNotFound) String() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaNotFound  %+v", 404, o.Payload)
}

func (o *SearchUserMetaNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchUserMetaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserMetaInternalServerError creates a SearchUserMetaInternalServerError with default headers values
func NewSearchUserMetaInternalServerError() *SearchUserMetaInternalServerError {
	return &SearchUserMetaInternalServerError{}
}

/*
SearchUserMetaInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SearchUserMetaInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this search user meta internal server error response has a 2xx status code
func (o *SearchUserMetaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user meta internal server error response has a 3xx status code
func (o *SearchUserMetaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user meta internal server error response has a 4xx status code
func (o *SearchUserMetaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search user meta internal server error response has a 5xx status code
func (o *SearchUserMetaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search user meta internal server error response a status code equal to that given
func (o *SearchUserMetaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search user meta internal server error response
func (o *SearchUserMetaInternalServerError) Code() int {
	return 500
}

func (o *SearchUserMetaInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUserMetaInternalServerError) String() string {
	return fmt.Sprintf("[POST /user-meta/search][%d] searchUserMetaInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUserMetaInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchUserMetaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
