// Code generated by go-swagger; DO NOT EDIT.

package role_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// SearchRolesReader is a Reader for the SearchRoles structure.
type SearchRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRolesOK creates a SearchRolesOK with default headers values
func NewSearchRolesOK() *SearchRolesOK {
	return &SearchRolesOK{}
}

/* SearchRolesOK describes a response with status code 200, with default header values.

SearchRolesOK search roles o k
*/
type SearchRolesOK struct {
	Payload *models.RestRolesCollection
}

func (o *SearchRolesOK) Error() string {
	return fmt.Sprintf("[POST /role][%d] searchRolesOK  %+v", 200, o.Payload)
}
func (o *SearchRolesOK) GetPayload() *models.RestRolesCollection {
	return o.Payload
}

func (o *SearchRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestRolesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRolesUnauthorized creates a SearchRolesUnauthorized with default headers values
func NewSearchRolesUnauthorized() *SearchRolesUnauthorized {
	return &SearchRolesUnauthorized{}
}

/* SearchRolesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SearchRolesUnauthorized struct {
}

func (o *SearchRolesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /role][%d] searchRolesUnauthorized ", 401)
}

func (o *SearchRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchRolesForbidden creates a SearchRolesForbidden with default headers values
func NewSearchRolesForbidden() *SearchRolesForbidden {
	return &SearchRolesForbidden{}
}

/* SearchRolesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SearchRolesForbidden struct {
	Payload *models.RestError
}

func (o *SearchRolesForbidden) Error() string {
	return fmt.Sprintf("[POST /role][%d] searchRolesForbidden  %+v", 403, o.Payload)
}
func (o *SearchRolesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRolesNotFound creates a SearchRolesNotFound with default headers values
func NewSearchRolesNotFound() *SearchRolesNotFound {
	return &SearchRolesNotFound{}
}

/* SearchRolesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SearchRolesNotFound struct {
	Payload *models.RestError
}

func (o *SearchRolesNotFound) Error() string {
	return fmt.Sprintf("[POST /role][%d] searchRolesNotFound  %+v", 404, o.Payload)
}
func (o *SearchRolesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRolesInternalServerError creates a SearchRolesInternalServerError with default headers values
func NewSearchRolesInternalServerError() *SearchRolesInternalServerError {
	return &SearchRolesInternalServerError{}
}

/* SearchRolesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SearchRolesInternalServerError struct {
	Payload *models.RestError
}

func (o *SearchRolesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /role][%d] searchRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchRolesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
