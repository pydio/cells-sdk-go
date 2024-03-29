// Code generated by go-swagger; DO NOT EDIT.

package token_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// GenerateDocumentAccessTokenReader is a Reader for the GenerateDocumentAccessToken structure.
type GenerateDocumentAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateDocumentAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateDocumentAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGenerateDocumentAccessTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenerateDocumentAccessTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGenerateDocumentAccessTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateDocumentAccessTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/token/document] GenerateDocumentAccessToken", response, response.Code())
	}
}

// NewGenerateDocumentAccessTokenOK creates a GenerateDocumentAccessTokenOK with default headers values
func NewGenerateDocumentAccessTokenOK() *GenerateDocumentAccessTokenOK {
	return &GenerateDocumentAccessTokenOK{}
}

/*
GenerateDocumentAccessTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type GenerateDocumentAccessTokenOK struct {
	Payload *models.RestDocumentAccessTokenResponse
}

// IsSuccess returns true when this generate document access token o k response has a 2xx status code
func (o *GenerateDocumentAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate document access token o k response has a 3xx status code
func (o *GenerateDocumentAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate document access token o k response has a 4xx status code
func (o *GenerateDocumentAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate document access token o k response has a 5xx status code
func (o *GenerateDocumentAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate document access token o k response a status code equal to that given
func (o *GenerateDocumentAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate document access token o k response
func (o *GenerateDocumentAccessTokenOK) Code() int {
	return 200
}

func (o *GenerateDocumentAccessTokenOK) Error() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenOK  %+v", 200, o.Payload)
}

func (o *GenerateDocumentAccessTokenOK) String() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenOK  %+v", 200, o.Payload)
}

func (o *GenerateDocumentAccessTokenOK) GetPayload() *models.RestDocumentAccessTokenResponse {
	return o.Payload
}

func (o *GenerateDocumentAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDocumentAccessTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateDocumentAccessTokenUnauthorized creates a GenerateDocumentAccessTokenUnauthorized with default headers values
func NewGenerateDocumentAccessTokenUnauthorized() *GenerateDocumentAccessTokenUnauthorized {
	return &GenerateDocumentAccessTokenUnauthorized{}
}

/*
GenerateDocumentAccessTokenUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GenerateDocumentAccessTokenUnauthorized struct {
}

// IsSuccess returns true when this generate document access token unauthorized response has a 2xx status code
func (o *GenerateDocumentAccessTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate document access token unauthorized response has a 3xx status code
func (o *GenerateDocumentAccessTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate document access token unauthorized response has a 4xx status code
func (o *GenerateDocumentAccessTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate document access token unauthorized response has a 5xx status code
func (o *GenerateDocumentAccessTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this generate document access token unauthorized response a status code equal to that given
func (o *GenerateDocumentAccessTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the generate document access token unauthorized response
func (o *GenerateDocumentAccessTokenUnauthorized) Code() int {
	return 401
}

func (o *GenerateDocumentAccessTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenUnauthorized ", 401)
}

func (o *GenerateDocumentAccessTokenUnauthorized) String() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenUnauthorized ", 401)
}

func (o *GenerateDocumentAccessTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateDocumentAccessTokenForbidden creates a GenerateDocumentAccessTokenForbidden with default headers values
func NewGenerateDocumentAccessTokenForbidden() *GenerateDocumentAccessTokenForbidden {
	return &GenerateDocumentAccessTokenForbidden{}
}

/*
GenerateDocumentAccessTokenForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GenerateDocumentAccessTokenForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this generate document access token forbidden response has a 2xx status code
func (o *GenerateDocumentAccessTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate document access token forbidden response has a 3xx status code
func (o *GenerateDocumentAccessTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate document access token forbidden response has a 4xx status code
func (o *GenerateDocumentAccessTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate document access token forbidden response has a 5xx status code
func (o *GenerateDocumentAccessTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this generate document access token forbidden response a status code equal to that given
func (o *GenerateDocumentAccessTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the generate document access token forbidden response
func (o *GenerateDocumentAccessTokenForbidden) Code() int {
	return 403
}

func (o *GenerateDocumentAccessTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenForbidden  %+v", 403, o.Payload)
}

func (o *GenerateDocumentAccessTokenForbidden) String() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenForbidden  %+v", 403, o.Payload)
}

func (o *GenerateDocumentAccessTokenForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GenerateDocumentAccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateDocumentAccessTokenNotFound creates a GenerateDocumentAccessTokenNotFound with default headers values
func NewGenerateDocumentAccessTokenNotFound() *GenerateDocumentAccessTokenNotFound {
	return &GenerateDocumentAccessTokenNotFound{}
}

/*
GenerateDocumentAccessTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GenerateDocumentAccessTokenNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this generate document access token not found response has a 2xx status code
func (o *GenerateDocumentAccessTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate document access token not found response has a 3xx status code
func (o *GenerateDocumentAccessTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate document access token not found response has a 4xx status code
func (o *GenerateDocumentAccessTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate document access token not found response has a 5xx status code
func (o *GenerateDocumentAccessTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this generate document access token not found response a status code equal to that given
func (o *GenerateDocumentAccessTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the generate document access token not found response
func (o *GenerateDocumentAccessTokenNotFound) Code() int {
	return 404
}

func (o *GenerateDocumentAccessTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenNotFound  %+v", 404, o.Payload)
}

func (o *GenerateDocumentAccessTokenNotFound) String() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenNotFound  %+v", 404, o.Payload)
}

func (o *GenerateDocumentAccessTokenNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GenerateDocumentAccessTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateDocumentAccessTokenInternalServerError creates a GenerateDocumentAccessTokenInternalServerError with default headers values
func NewGenerateDocumentAccessTokenInternalServerError() *GenerateDocumentAccessTokenInternalServerError {
	return &GenerateDocumentAccessTokenInternalServerError{}
}

/*
GenerateDocumentAccessTokenInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GenerateDocumentAccessTokenInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this generate document access token internal server error response has a 2xx status code
func (o *GenerateDocumentAccessTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate document access token internal server error response has a 3xx status code
func (o *GenerateDocumentAccessTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate document access token internal server error response has a 4xx status code
func (o *GenerateDocumentAccessTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate document access token internal server error response has a 5xx status code
func (o *GenerateDocumentAccessTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate document access token internal server error response a status code equal to that given
func (o *GenerateDocumentAccessTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the generate document access token internal server error response
func (o *GenerateDocumentAccessTokenInternalServerError) Code() int {
	return 500
}

func (o *GenerateDocumentAccessTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateDocumentAccessTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/token/document][%d] generateDocumentAccessTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateDocumentAccessTokenInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GenerateDocumentAccessTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
