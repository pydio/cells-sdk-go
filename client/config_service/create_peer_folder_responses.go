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

	"github.com/pydio/cells-sdk-go/v5/models"
)

// CreatePeerFolderReader is a Reader for the CreatePeerFolder structure.
type CreatePeerFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeerFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePeerFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePeerFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePeerFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePeerFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePeerFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /config/peers/{PeerAddress}] CreatePeerFolder", response, response.Code())
	}
}

// NewCreatePeerFolderOK creates a CreatePeerFolderOK with default headers values
func NewCreatePeerFolderOK() *CreatePeerFolderOK {
	return &CreatePeerFolderOK{}
}

/*
CreatePeerFolderOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreatePeerFolderOK struct {
	Payload *models.RestCreatePeerFolderResponse
}

// IsSuccess returns true when this create peer folder o k response has a 2xx status code
func (o *CreatePeerFolderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create peer folder o k response has a 3xx status code
func (o *CreatePeerFolderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer folder o k response has a 4xx status code
func (o *CreatePeerFolderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create peer folder o k response has a 5xx status code
func (o *CreatePeerFolderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer folder o k response a status code equal to that given
func (o *CreatePeerFolderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create peer folder o k response
func (o *CreatePeerFolderOK) Code() int {
	return 200
}

func (o *CreatePeerFolderOK) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderOK  %+v", 200, o.Payload)
}

func (o *CreatePeerFolderOK) String() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderOK  %+v", 200, o.Payload)
}

func (o *CreatePeerFolderOK) GetPayload() *models.RestCreatePeerFolderResponse {
	return o.Payload
}

func (o *CreatePeerFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCreatePeerFolderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerFolderUnauthorized creates a CreatePeerFolderUnauthorized with default headers values
func NewCreatePeerFolderUnauthorized() *CreatePeerFolderUnauthorized {
	return &CreatePeerFolderUnauthorized{}
}

/*
CreatePeerFolderUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type CreatePeerFolderUnauthorized struct {
}

// IsSuccess returns true when this create peer folder unauthorized response has a 2xx status code
func (o *CreatePeerFolderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer folder unauthorized response has a 3xx status code
func (o *CreatePeerFolderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer folder unauthorized response has a 4xx status code
func (o *CreatePeerFolderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create peer folder unauthorized response has a 5xx status code
func (o *CreatePeerFolderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer folder unauthorized response a status code equal to that given
func (o *CreatePeerFolderUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create peer folder unauthorized response
func (o *CreatePeerFolderUnauthorized) Code() int {
	return 401
}

func (o *CreatePeerFolderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderUnauthorized ", 401)
}

func (o *CreatePeerFolderUnauthorized) String() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderUnauthorized ", 401)
}

func (o *CreatePeerFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePeerFolderForbidden creates a CreatePeerFolderForbidden with default headers values
func NewCreatePeerFolderForbidden() *CreatePeerFolderForbidden {
	return &CreatePeerFolderForbidden{}
}

/*
CreatePeerFolderForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type CreatePeerFolderForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create peer folder forbidden response has a 2xx status code
func (o *CreatePeerFolderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer folder forbidden response has a 3xx status code
func (o *CreatePeerFolderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer folder forbidden response has a 4xx status code
func (o *CreatePeerFolderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create peer folder forbidden response has a 5xx status code
func (o *CreatePeerFolderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer folder forbidden response a status code equal to that given
func (o *CreatePeerFolderForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create peer folder forbidden response
func (o *CreatePeerFolderForbidden) Code() int {
	return 403
}

func (o *CreatePeerFolderForbidden) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderForbidden  %+v", 403, o.Payload)
}

func (o *CreatePeerFolderForbidden) String() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderForbidden  %+v", 403, o.Payload)
}

func (o *CreatePeerFolderForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreatePeerFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerFolderNotFound creates a CreatePeerFolderNotFound with default headers values
func NewCreatePeerFolderNotFound() *CreatePeerFolderNotFound {
	return &CreatePeerFolderNotFound{}
}

/*
CreatePeerFolderNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type CreatePeerFolderNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create peer folder not found response has a 2xx status code
func (o *CreatePeerFolderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer folder not found response has a 3xx status code
func (o *CreatePeerFolderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer folder not found response has a 4xx status code
func (o *CreatePeerFolderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create peer folder not found response has a 5xx status code
func (o *CreatePeerFolderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer folder not found response a status code equal to that given
func (o *CreatePeerFolderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create peer folder not found response
func (o *CreatePeerFolderNotFound) Code() int {
	return 404
}

func (o *CreatePeerFolderNotFound) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderNotFound  %+v", 404, o.Payload)
}

func (o *CreatePeerFolderNotFound) String() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderNotFound  %+v", 404, o.Payload)
}

func (o *CreatePeerFolderNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreatePeerFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerFolderInternalServerError creates a CreatePeerFolderInternalServerError with default headers values
func NewCreatePeerFolderInternalServerError() *CreatePeerFolderInternalServerError {
	return &CreatePeerFolderInternalServerError{}
}

/*
CreatePeerFolderInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type CreatePeerFolderInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this create peer folder internal server error response has a 2xx status code
func (o *CreatePeerFolderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer folder internal server error response has a 3xx status code
func (o *CreatePeerFolderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer folder internal server error response has a 4xx status code
func (o *CreatePeerFolderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create peer folder internal server error response has a 5xx status code
func (o *CreatePeerFolderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create peer folder internal server error response a status code equal to that given
func (o *CreatePeerFolderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create peer folder internal server error response
func (o *CreatePeerFolderInternalServerError) Code() int {
	return 500
}

func (o *CreatePeerFolderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePeerFolderInternalServerError) String() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePeerFolderInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *CreatePeerFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreatePeerFolderBody RestCreatePeerFolderRequest
swagger:model CreatePeerFolderBody
*/
type CreatePeerFolderBody struct {

	// Path to the folder to be created
	Path string `json:"Path,omitempty"`
}

// Validate validates this create peer folder body
func (o *CreatePeerFolderBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create peer folder body based on context it is used
func (o *CreatePeerFolderBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreatePeerFolderBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePeerFolderBody) UnmarshalBinary(b []byte) error {
	var res CreatePeerFolderBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
