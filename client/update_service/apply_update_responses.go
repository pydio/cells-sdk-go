// Code generated by go-swagger; DO NOT EDIT.

package update_service

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

// ApplyUpdateReader is a Reader for the ApplyUpdate structure.
type ApplyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewApplyUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewApplyUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewApplyUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplyUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /update/{TargetVersion}] ApplyUpdate", response, response.Code())
	}
}

// NewApplyUpdateOK creates a ApplyUpdateOK with default headers values
func NewApplyUpdateOK() *ApplyUpdateOK {
	return &ApplyUpdateOK{}
}

/*
ApplyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplyUpdateOK struct {
	Payload *models.UpdateApplyUpdateResponse
}

// IsSuccess returns true when this apply update o k response has a 2xx status code
func (o *ApplyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this apply update o k response has a 3xx status code
func (o *ApplyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply update o k response has a 4xx status code
func (o *ApplyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this apply update o k response has a 5xx status code
func (o *ApplyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this apply update o k response a status code equal to that given
func (o *ApplyUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the apply update o k response
func (o *ApplyUpdateOK) Code() int {
	return 200
}

func (o *ApplyUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateOK  %+v", 200, o.Payload)
}

func (o *ApplyUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateOK  %+v", 200, o.Payload)
}

func (o *ApplyUpdateOK) GetPayload() *models.UpdateApplyUpdateResponse {
	return o.Payload
}

func (o *ApplyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateApplyUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyUpdateUnauthorized creates a ApplyUpdateUnauthorized with default headers values
func NewApplyUpdateUnauthorized() *ApplyUpdateUnauthorized {
	return &ApplyUpdateUnauthorized{}
}

/*
ApplyUpdateUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ApplyUpdateUnauthorized struct {
}

// IsSuccess returns true when this apply update unauthorized response has a 2xx status code
func (o *ApplyUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply update unauthorized response has a 3xx status code
func (o *ApplyUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply update unauthorized response has a 4xx status code
func (o *ApplyUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this apply update unauthorized response has a 5xx status code
func (o *ApplyUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this apply update unauthorized response a status code equal to that given
func (o *ApplyUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the apply update unauthorized response
func (o *ApplyUpdateUnauthorized) Code() int {
	return 401
}

func (o *ApplyUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateUnauthorized ", 401)
}

func (o *ApplyUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateUnauthorized ", 401)
}

func (o *ApplyUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyUpdateForbidden creates a ApplyUpdateForbidden with default headers values
func NewApplyUpdateForbidden() *ApplyUpdateForbidden {
	return &ApplyUpdateForbidden{}
}

/*
ApplyUpdateForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ApplyUpdateForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this apply update forbidden response has a 2xx status code
func (o *ApplyUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply update forbidden response has a 3xx status code
func (o *ApplyUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply update forbidden response has a 4xx status code
func (o *ApplyUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this apply update forbidden response has a 5xx status code
func (o *ApplyUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this apply update forbidden response a status code equal to that given
func (o *ApplyUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the apply update forbidden response
func (o *ApplyUpdateForbidden) Code() int {
	return 403
}

func (o *ApplyUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ApplyUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ApplyUpdateForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ApplyUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyUpdateNotFound creates a ApplyUpdateNotFound with default headers values
func NewApplyUpdateNotFound() *ApplyUpdateNotFound {
	return &ApplyUpdateNotFound{}
}

/*
ApplyUpdateNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ApplyUpdateNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this apply update not found response has a 2xx status code
func (o *ApplyUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply update not found response has a 3xx status code
func (o *ApplyUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply update not found response has a 4xx status code
func (o *ApplyUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this apply update not found response has a 5xx status code
func (o *ApplyUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this apply update not found response a status code equal to that given
func (o *ApplyUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the apply update not found response
func (o *ApplyUpdateNotFound) Code() int {
	return 404
}

func (o *ApplyUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ApplyUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ApplyUpdateNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ApplyUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyUpdateInternalServerError creates a ApplyUpdateInternalServerError with default headers values
func NewApplyUpdateInternalServerError() *ApplyUpdateInternalServerError {
	return &ApplyUpdateInternalServerError{}
}

/*
ApplyUpdateInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ApplyUpdateInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this apply update internal server error response has a 2xx status code
func (o *ApplyUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply update internal server error response has a 3xx status code
func (o *ApplyUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply update internal server error response has a 4xx status code
func (o *ApplyUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this apply update internal server error response has a 5xx status code
func (o *ApplyUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this apply update internal server error response a status code equal to that given
func (o *ApplyUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the apply update internal server error response
func (o *ApplyUpdateInternalServerError) Code() int {
	return 500
}

func (o *ApplyUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplyUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /update/{TargetVersion}][%d] applyUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplyUpdateInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ApplyUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ApplyUpdateBody UpdateApplyUpdateRequest
swagger:model ApplyUpdateBody
*/
type ApplyUpdateBody struct {

	// Name of the package if it's not the same as the current binary
	PackageName string `json:"PackageName,omitempty"`
}

// Validate validates this apply update body
func (o *ApplyUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apply update body based on context it is used
func (o *ApplyUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ApplyUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ApplyUpdateBody) UnmarshalBinary(b []byte) error {
	var res ApplyUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
