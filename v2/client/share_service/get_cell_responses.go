// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// GetCellReader is a Reader for the GetCell structure.
type GetCellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCellOK creates a GetCellOK with default headers values
func NewGetCellOK() *GetCellOK {
	return &GetCellOK{}
}

/*GetCellOK handles this case with default header values.

GetCellOK get cell o k
*/
type GetCellOK struct {
	Payload *models.RestCell
}

func (o *GetCellOK) Error() string {
	return fmt.Sprintf("[GET /share/cell/{Uuid}][%d] getCellOK  %+v", 200, o.Payload)
}

func (o *GetCellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCell)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
