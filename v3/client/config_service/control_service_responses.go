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

// ControlServiceReader is a Reader for the ControlService structure.
type ControlServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControlServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewControlServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewControlServiceOK creates a ControlServiceOK with default headers values
func NewControlServiceOK() *ControlServiceOK {
	return &ControlServiceOK{}
}

/* ControlServiceOK describes a response with status code 200, with default header values.

ControlServiceOK control service o k
*/
type ControlServiceOK struct {
	Payload *models.CtlService
}

func (o *ControlServiceOK) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceOK  %+v", 200, o.Payload)
}
func (o *ControlServiceOK) GetPayload() *models.CtlService {
	return o.Payload
}

func (o *ControlServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CtlService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
