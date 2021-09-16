// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// FrontBootConfReader is a Reader for the FrontBootConf structure.
type FrontBootConfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontBootConfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrontBootConfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFrontBootConfOK creates a FrontBootConfOK with default headers values
func NewFrontBootConfOK() *FrontBootConfOK {
	return &FrontBootConfOK{}
}

/* FrontBootConfOK describes a response with status code 200, with default header values.

FrontBootConfOK front boot conf o k
*/
type FrontBootConfOK struct {
	Payload *models.RestFrontBootConfResponse
}

func (o *FrontBootConfOK) Error() string {
	return fmt.Sprintf("[GET /frontend/bootconf][%d] frontBootConfOK  %+v", 200, o.Payload)
}
func (o *FrontBootConfOK) GetPayload() *models.RestFrontBootConfResponse {
	return o.Payload
}

func (o *FrontBootConfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestFrontBootConfResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
