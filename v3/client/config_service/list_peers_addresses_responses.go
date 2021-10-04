// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// ListPeersAddressesReader is a Reader for the ListPeersAddresses structure.
type ListPeersAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPeersAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPeersAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPeersAddressesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPeersAddressesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPeersAddressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPeersAddressesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPeersAddressesOK creates a ListPeersAddressesOK with default headers values
func NewListPeersAddressesOK() *ListPeersAddressesOK {
	return &ListPeersAddressesOK{}
}

/* ListPeersAddressesOK describes a response with status code 200, with default header values.

ListPeersAddressesOK list peers addresses o k
*/
type ListPeersAddressesOK struct {
	Payload *models.RestListPeersAddressesResponse
}

func (o *ListPeersAddressesOK) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesOK  %+v", 200, o.Payload)
}
func (o *ListPeersAddressesOK) GetPayload() *models.RestListPeersAddressesResponse {
	return o.Payload
}

func (o *ListPeersAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListPeersAddressesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeersAddressesUnauthorized creates a ListPeersAddressesUnauthorized with default headers values
func NewListPeersAddressesUnauthorized() *ListPeersAddressesUnauthorized {
	return &ListPeersAddressesUnauthorized{}
}

/* ListPeersAddressesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListPeersAddressesUnauthorized struct {
	Payload *models.RestError
}

func (o *ListPeersAddressesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListPeersAddressesUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeersAddressesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeersAddressesForbidden creates a ListPeersAddressesForbidden with default headers values
func NewListPeersAddressesForbidden() *ListPeersAddressesForbidden {
	return &ListPeersAddressesForbidden{}
}

/* ListPeersAddressesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListPeersAddressesForbidden struct {
	Payload *models.RestError
}

func (o *ListPeersAddressesForbidden) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesForbidden  %+v", 403, o.Payload)
}
func (o *ListPeersAddressesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeersAddressesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeersAddressesNotFound creates a ListPeersAddressesNotFound with default headers values
func NewListPeersAddressesNotFound() *ListPeersAddressesNotFound {
	return &ListPeersAddressesNotFound{}
}

/* ListPeersAddressesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListPeersAddressesNotFound struct {
	Payload *models.RestError
}

func (o *ListPeersAddressesNotFound) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesNotFound  %+v", 404, o.Payload)
}
func (o *ListPeersAddressesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeersAddressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeersAddressesInternalServerError creates a ListPeersAddressesInternalServerError with default headers values
func NewListPeersAddressesInternalServerError() *ListPeersAddressesInternalServerError {
	return &ListPeersAddressesInternalServerError{}
}

/* ListPeersAddressesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListPeersAddressesInternalServerError struct {
	Payload *models.RestError
}

func (o *ListPeersAddressesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListPeersAddressesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeersAddressesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
