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

// PutConfigReader is a Reader for the PutConfig structure.
type PutConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /config/{FullPath}] PutConfig", response, response.Code())
	}
}

// NewPutConfigOK creates a PutConfigOK with default headers values
func NewPutConfigOK() *PutConfigOK {
	return &PutConfigOK{}
}

/*
PutConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type PutConfigOK struct {
	Payload *models.RestConfiguration
}

// IsSuccess returns true when this put config o k response has a 2xx status code
func (o *PutConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put config o k response has a 3xx status code
func (o *PutConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config o k response has a 4xx status code
func (o *PutConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config o k response has a 5xx status code
func (o *PutConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put config o k response a status code equal to that given
func (o *PutConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put config o k response
func (o *PutConfigOK) Code() int {
	return 200
}

func (o *PutConfigOK) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigOK  %+v", 200, o.Payload)
}

func (o *PutConfigOK) String() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigOK  %+v", 200, o.Payload)
}

func (o *PutConfigOK) GetPayload() *models.RestConfiguration {
	return o.Payload
}

func (o *PutConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigUnauthorized creates a PutConfigUnauthorized with default headers values
func NewPutConfigUnauthorized() *PutConfigUnauthorized {
	return &PutConfigUnauthorized{}
}

/*
PutConfigUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutConfigUnauthorized struct {
}

// IsSuccess returns true when this put config unauthorized response has a 2xx status code
func (o *PutConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config unauthorized response has a 3xx status code
func (o *PutConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config unauthorized response has a 4xx status code
func (o *PutConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put config unauthorized response has a 5xx status code
func (o *PutConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put config unauthorized response a status code equal to that given
func (o *PutConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put config unauthorized response
func (o *PutConfigUnauthorized) Code() int {
	return 401
}

func (o *PutConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigUnauthorized ", 401)
}

func (o *PutConfigUnauthorized) String() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigUnauthorized ", 401)
}

func (o *PutConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutConfigForbidden creates a PutConfigForbidden with default headers values
func NewPutConfigForbidden() *PutConfigForbidden {
	return &PutConfigForbidden{}
}

/*
PutConfigForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PutConfigForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put config forbidden response has a 2xx status code
func (o *PutConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config forbidden response has a 3xx status code
func (o *PutConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config forbidden response has a 4xx status code
func (o *PutConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put config forbidden response has a 5xx status code
func (o *PutConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put config forbidden response a status code equal to that given
func (o *PutConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put config forbidden response
func (o *PutConfigForbidden) Code() int {
	return 403
}

func (o *PutConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigForbidden  %+v", 403, o.Payload)
}

func (o *PutConfigForbidden) String() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigForbidden  %+v", 403, o.Payload)
}

func (o *PutConfigForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigNotFound creates a PutConfigNotFound with default headers values
func NewPutConfigNotFound() *PutConfigNotFound {
	return &PutConfigNotFound{}
}

/*
PutConfigNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutConfigNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put config not found response has a 2xx status code
func (o *PutConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config not found response has a 3xx status code
func (o *PutConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config not found response has a 4xx status code
func (o *PutConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put config not found response has a 5xx status code
func (o *PutConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put config not found response a status code equal to that given
func (o *PutConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put config not found response
func (o *PutConfigNotFound) Code() int {
	return 404
}

func (o *PutConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigNotFound  %+v", 404, o.Payload)
}

func (o *PutConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigNotFound  %+v", 404, o.Payload)
}

func (o *PutConfigNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigInternalServerError creates a PutConfigInternalServerError with default headers values
func NewPutConfigInternalServerError() *PutConfigInternalServerError {
	return &PutConfigInternalServerError{}
}

/*
PutConfigInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutConfigInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put config internal server error response has a 2xx status code
func (o *PutConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config internal server error response has a 3xx status code
func (o *PutConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config internal server error response has a 4xx status code
func (o *PutConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config internal server error response has a 5xx status code
func (o *PutConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put config internal server error response a status code equal to that given
func (o *PutConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put config internal server error response
func (o *PutConfigInternalServerError) Code() int {
	return 500
}

func (o *PutConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutConfigBody Configuration message. Data is an Json representation of any value
swagger:model PutConfigBody
*/
type PutConfigBody struct {

	// JSON-encoded data to store
	Data string `json:"Data,omitempty"`
}

// Validate validates this put config body
func (o *PutConfigBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put config body based on context it is used
func (o *PutConfigBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutConfigBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutConfigBody) UnmarshalBinary(b []byte) error {
	var res PutConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
