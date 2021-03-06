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

// PostInstallReader is a Reader for the PostInstall structure.
type PostInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInstallOK creates a PostInstallOK with default headers values
func NewPostInstallOK() *PostInstallOK {
	return &PostInstallOK{}
}

/*PostInstallOK handles this case with default header values.

PostInstallOK post install o k
*/
type PostInstallOK struct {
	Payload *models.InstallInstallResponse
}

func (o *PostInstallOK) Error() string {
	return fmt.Sprintf("[POST /install][%d] postInstallOK  %+v", 200, o.Payload)
}

func (o *PostInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallInstallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
