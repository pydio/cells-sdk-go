// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pydio/cells-sdk-go/v4/models"
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
		return nil, runtime.NewAPIError("[POST /config/peers/{PeerAddress}] ListPeerFolders", response, response.Code())
	}
}

// NewListPeerFoldersOK creates a ListPeerFoldersOK with default headers values
func NewListPeerFoldersOK() *ListPeerFoldersOK {
	return &ListPeerFoldersOK{}
}

/*
ListPeerFoldersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPeerFoldersOK struct {
	Payload *models.RestNodesCollection
}

// IsSuccess returns true when this list peer folders o k response has a 2xx status code
func (o *ListPeerFoldersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list peer folders o k response has a 3xx status code
func (o *ListPeerFoldersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list peer folders o k response has a 4xx status code
func (o *ListPeerFoldersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list peer folders o k response has a 5xx status code
func (o *ListPeerFoldersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list peer folders o k response a status code equal to that given
func (o *ListPeerFoldersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list peer folders o k response
func (o *ListPeerFoldersOK) Code() int {
	return 200
}

func (o *ListPeerFoldersOK) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersOK  %+v", 200, o.Payload)
}

func (o *ListPeerFoldersOK) String() string {
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

/*
ListPeerFoldersUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListPeerFoldersUnauthorized struct {
}

// IsSuccess returns true when this list peer folders unauthorized response has a 2xx status code
func (o *ListPeerFoldersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list peer folders unauthorized response has a 3xx status code
func (o *ListPeerFoldersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list peer folders unauthorized response has a 4xx status code
func (o *ListPeerFoldersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list peer folders unauthorized response has a 5xx status code
func (o *ListPeerFoldersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list peer folders unauthorized response a status code equal to that given
func (o *ListPeerFoldersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list peer folders unauthorized response
func (o *ListPeerFoldersUnauthorized) Code() int {
	return 401
}

func (o *ListPeerFoldersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersUnauthorized ", 401)
}

func (o *ListPeerFoldersUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersUnauthorized ", 401)
}

func (o *ListPeerFoldersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPeerFoldersForbidden creates a ListPeerFoldersForbidden with default headers values
func NewListPeerFoldersForbidden() *ListPeerFoldersForbidden {
	return &ListPeerFoldersForbidden{}
}

/*
ListPeerFoldersForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListPeerFoldersForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list peer folders forbidden response has a 2xx status code
func (o *ListPeerFoldersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list peer folders forbidden response has a 3xx status code
func (o *ListPeerFoldersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list peer folders forbidden response has a 4xx status code
func (o *ListPeerFoldersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list peer folders forbidden response has a 5xx status code
func (o *ListPeerFoldersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list peer folders forbidden response a status code equal to that given
func (o *ListPeerFoldersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list peer folders forbidden response
func (o *ListPeerFoldersForbidden) Code() int {
	return 403
}

func (o *ListPeerFoldersForbidden) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersForbidden  %+v", 403, o.Payload)
}

func (o *ListPeerFoldersForbidden) String() string {
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

/*
ListPeerFoldersNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListPeerFoldersNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list peer folders not found response has a 2xx status code
func (o *ListPeerFoldersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list peer folders not found response has a 3xx status code
func (o *ListPeerFoldersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list peer folders not found response has a 4xx status code
func (o *ListPeerFoldersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list peer folders not found response has a 5xx status code
func (o *ListPeerFoldersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list peer folders not found response a status code equal to that given
func (o *ListPeerFoldersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list peer folders not found response
func (o *ListPeerFoldersNotFound) Code() int {
	return 404
}

func (o *ListPeerFoldersNotFound) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersNotFound  %+v", 404, o.Payload)
}

func (o *ListPeerFoldersNotFound) String() string {
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

/*
ListPeerFoldersInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListPeerFoldersInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list peer folders internal server error response has a 2xx status code
func (o *ListPeerFoldersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list peer folders internal server error response has a 3xx status code
func (o *ListPeerFoldersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list peer folders internal server error response has a 4xx status code
func (o *ListPeerFoldersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list peer folders internal server error response has a 5xx status code
func (o *ListPeerFoldersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list peer folders internal server error response a status code equal to that given
func (o *ListPeerFoldersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list peer folders internal server error response
func (o *ListPeerFoldersInternalServerError) Code() int {
	return 500
}

func (o *ListPeerFoldersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPeerFoldersInternalServerError) String() string {
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

/*
ListPeerFoldersBody RestListPeerFoldersRequest
swagger:model ListPeerFoldersBody
*/
type ListPeerFoldersBody struct {

	// Path to the parent folder for listing
	Path string `json:"Path,omitempty"`
}

// Validate validates this list peer folders body
func (o *ListPeerFoldersBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list peer folders body based on context it is used
func (o *ListPeerFoldersBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPeerFoldersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPeerFoldersBody) UnmarshalBinary(b []byte) error {
	var res ListPeerFoldersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
