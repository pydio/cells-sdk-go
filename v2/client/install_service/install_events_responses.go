// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// InstallEventsReader is a Reader for the InstallEvents structure.
type InstallEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInstallEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstallEventsOK creates a InstallEventsOK with default headers values
func NewInstallEventsOK() *InstallEventsOK {
	return &InstallEventsOK{}
}

/*InstallEventsOK handles this case with default header values.

InstallEventsOK install events o k
*/
type InstallEventsOK struct {
	Payload models.InstallInstallEventsResponse
}

func (o *InstallEventsOK) Error() string {
	return fmt.Sprintf("[GET /install/events][%d] installEventsOK  %+v", 200, o.Payload)
}

func (o *InstallEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}