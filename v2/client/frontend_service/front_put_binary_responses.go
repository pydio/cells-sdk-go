// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// FrontPutBinaryReader is a Reader for the FrontPutBinary structure.
type FrontPutBinaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontPutBinaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFrontPutBinaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFrontPutBinaryOK creates a FrontPutBinaryOK with default headers values
func NewFrontPutBinaryOK() *FrontPutBinaryOK {
	return &FrontPutBinaryOK{}
}

/*FrontPutBinaryOK handles this case with default header values.

FrontPutBinaryOK front put binary o k
*/
type FrontPutBinaryOK struct {
	Payload models.RestFrontBinaryResponse
}

func (o *FrontPutBinaryOK) Error() string {
	return fmt.Sprintf("[POST /frontend/binaries/{BinaryType}/{Uuid}][%d] frontPutBinaryOK  %+v", 200, o.Payload)
}

func (o *FrontPutBinaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}