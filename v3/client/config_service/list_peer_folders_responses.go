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

// ListPeerFoldersReader is a Reader for the ListPeerFolders structure.
type ListPeerFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPeerFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPeerFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPeerFoldersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPeerFoldersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPeerFoldersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPeerFoldersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPeerFoldersOK creates a ListPeerFoldersOK with default headers values
func NewListPeerFoldersOK() *ListPeerFoldersOK {
	return &ListPeerFoldersOK{}
}

/* ListPeerFoldersOK describes a response with status code 200, with default header values.

ListPeerFoldersOK list peer folders o k
*/
type ListPeerFoldersOK struct {
	Payload *models.RestNodesCollection
}

func (o *ListPeerFoldersOK) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersOK  %+v", 200, o.Payload)
}
func (o *ListPeerFoldersOK) GetPayload() *models.RestNodesCollection {
	return o.Payload
}

func (o *ListPeerFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestNodesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeerFoldersUnauthorized creates a ListPeerFoldersUnauthorized with default headers values
func NewListPeerFoldersUnauthorized() *ListPeerFoldersUnauthorized {
	return &ListPeerFoldersUnauthorized{}
}

/* ListPeerFoldersUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListPeerFoldersUnauthorized struct {
}

func (o *ListPeerFoldersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersUnauthorized ", 401)
}

func (o *ListPeerFoldersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPeerFoldersForbidden creates a ListPeerFoldersForbidden with default headers values
func NewListPeerFoldersForbidden() *ListPeerFoldersForbidden {
	return &ListPeerFoldersForbidden{}
}

/* ListPeerFoldersForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListPeerFoldersForbidden struct {
	Payload *models.RestError
}

func (o *ListPeerFoldersForbidden) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersForbidden  %+v", 403, o.Payload)
}
func (o *ListPeerFoldersForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeerFoldersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeerFoldersNotFound creates a ListPeerFoldersNotFound with default headers values
func NewListPeerFoldersNotFound() *ListPeerFoldersNotFound {
	return &ListPeerFoldersNotFound{}
}

/* ListPeerFoldersNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListPeerFoldersNotFound struct {
	Payload *models.RestError
}

func (o *ListPeerFoldersNotFound) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersNotFound  %+v", 404, o.Payload)
}
func (o *ListPeerFoldersNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeerFoldersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPeerFoldersInternalServerError creates a ListPeerFoldersInternalServerError with default headers values
func NewListPeerFoldersInternalServerError() *ListPeerFoldersInternalServerError {
	return &ListPeerFoldersInternalServerError{}
}

/* ListPeerFoldersInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListPeerFoldersInternalServerError struct {
	Payload *models.RestError
}

func (o *ListPeerFoldersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersInternalServerError  %+v", 500, o.Payload)
}
func (o *ListPeerFoldersInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListPeerFoldersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
