// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// PerformInstallCheckReader is a Reader for the PerformInstallCheck structure.
type PerformInstallCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformInstallCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPerformInstallCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPerformInstallCheckOK creates a PerformInstallCheckOK with default headers values
func NewPerformInstallCheckOK() *PerformInstallCheckOK {
	return &PerformInstallCheckOK{}
}

/*PerformInstallCheckOK handles this case with default header values.

PerformInstallCheckOK perform install check o k
*/
type PerformInstallCheckOK struct {
	Payload *models.InstallPerformCheckResponse
}

func (o *PerformInstallCheckOK) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckOK  %+v", 200, o.Payload)
}

func (o *PerformInstallCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallPerformCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}