// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// PutDataSourceReader is a Reader for the PutDataSource structure.
type PutDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDataSourceOK creates a PutDataSourceOK with default headers values
func NewPutDataSourceOK() *PutDataSourceOK {
	return &PutDataSourceOK{}
}

/*PutDataSourceOK handles this case with default header values.

PutDataSourceOK put data source o k
*/
type PutDataSourceOK struct {
	Payload *models.ObjectDataSource
}

func (o *PutDataSourceOK) Error() string {
	return fmt.Sprintf("[POST /config/datasource/{Name}][%d] putDataSourceOK  %+v", 200, o.Payload)
}

func (o *PutDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectDataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
