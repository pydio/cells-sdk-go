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

// ListPeersAddressesReader is a Reader for the ListPeersAddresses structure.
type ListPeersAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPeersAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPeersAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPeersAddressesOK creates a ListPeersAddressesOK with default headers values
func NewListPeersAddressesOK() *ListPeersAddressesOK {
	return &ListPeersAddressesOK{}
}

/*ListPeersAddressesOK handles this case with default header values.

ListPeersAddressesOK list peers addresses o k
*/
type ListPeersAddressesOK struct {
	Payload *models.RestListPeersAddressesResponse
}

func (o *ListPeersAddressesOK) Error() string {
	return fmt.Sprintf("[GET /config/peers][%d] listPeersAddressesOK  %+v", 200, o.Payload)
}

func (o *ListPeersAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListPeersAddressesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
