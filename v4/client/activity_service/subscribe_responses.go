// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// SubscribeReader is a Reader for the Subscribe structure.
type SubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscribeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscribeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscribeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscribeOK creates a SubscribeOK with default headers values
func NewSubscribeOK() *SubscribeOK {
	return &SubscribeOK{}
}

/*
SubscribeOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscribeOK struct {
	Payload *models.ActivitySubscription
}

// IsSuccess returns true when this subscribe o k response has a 2xx status code
func (o *SubscribeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscribe o k response has a 3xx status code
func (o *SubscribeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscribe o k response has a 4xx status code
func (o *SubscribeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscribe o k response has a 5xx status code
func (o *SubscribeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscribe o k response a status code equal to that given
func (o *SubscribeOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscribeOK) Error() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeOK  %+v", 200, o.Payload)
}

func (o *SubscribeOK) String() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeOK  %+v", 200, o.Payload)
}

func (o *SubscribeOK) GetPayload() *models.ActivitySubscription {
	return o.Payload
}

func (o *SubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivitySubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscribeUnauthorized creates a SubscribeUnauthorized with default headers values
func NewSubscribeUnauthorized() *SubscribeUnauthorized {
	return &SubscribeUnauthorized{}
}

/*
SubscribeUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SubscribeUnauthorized struct {
}

// IsSuccess returns true when this subscribe unauthorized response has a 2xx status code
func (o *SubscribeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscribe unauthorized response has a 3xx status code
func (o *SubscribeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscribe unauthorized response has a 4xx status code
func (o *SubscribeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscribe unauthorized response has a 5xx status code
func (o *SubscribeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscribe unauthorized response a status code equal to that given
func (o *SubscribeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeUnauthorized ", 401)
}

func (o *SubscribeUnauthorized) String() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeUnauthorized ", 401)
}

func (o *SubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubscribeForbidden creates a SubscribeForbidden with default headers values
func NewSubscribeForbidden() *SubscribeForbidden {
	return &SubscribeForbidden{}
}

/*
SubscribeForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SubscribeForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this subscribe forbidden response has a 2xx status code
func (o *SubscribeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscribe forbidden response has a 3xx status code
func (o *SubscribeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscribe forbidden response has a 4xx status code
func (o *SubscribeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscribe forbidden response has a 5xx status code
func (o *SubscribeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscribe forbidden response a status code equal to that given
func (o *SubscribeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubscribeForbidden) Error() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeForbidden  %+v", 403, o.Payload)
}

func (o *SubscribeForbidden) String() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeForbidden  %+v", 403, o.Payload)
}

func (o *SubscribeForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SubscribeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscribeNotFound creates a SubscribeNotFound with default headers values
func NewSubscribeNotFound() *SubscribeNotFound {
	return &SubscribeNotFound{}
}

/*
SubscribeNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SubscribeNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this subscribe not found response has a 2xx status code
func (o *SubscribeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscribe not found response has a 3xx status code
func (o *SubscribeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscribe not found response has a 4xx status code
func (o *SubscribeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscribe not found response has a 5xx status code
func (o *SubscribeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscribe not found response a status code equal to that given
func (o *SubscribeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubscribeNotFound) Error() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeNotFound  %+v", 404, o.Payload)
}

func (o *SubscribeNotFound) String() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeNotFound  %+v", 404, o.Payload)
}

func (o *SubscribeNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SubscribeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscribeInternalServerError creates a SubscribeInternalServerError with default headers values
func NewSubscribeInternalServerError() *SubscribeInternalServerError {
	return &SubscribeInternalServerError{}
}

/*
SubscribeInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SubscribeInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this subscribe internal server error response has a 2xx status code
func (o *SubscribeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscribe internal server error response has a 3xx status code
func (o *SubscribeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscribe internal server error response has a 4xx status code
func (o *SubscribeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscribe internal server error response has a 5xx status code
func (o *SubscribeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscribe internal server error response a status code equal to that given
func (o *SubscribeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubscribeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeInternalServerError  %+v", 500, o.Payload)
}

func (o *SubscribeInternalServerError) String() string {
	return fmt.Sprintf("[POST /activity/subscribe][%d] subscribeInternalServerError  %+v", 500, o.Payload)
}

func (o *SubscribeInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SubscribeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
