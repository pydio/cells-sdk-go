// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// UserControlJobReader is a Reader for the UserControlJob structure.
type UserControlJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserControlJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserControlJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserControlJobOK creates a UserControlJobOK with default headers values
func NewUserControlJobOK() *UserControlJobOK {
	return &UserControlJobOK{}
}

/*UserControlJobOK handles this case with default header values.

UserControlJobOK user control job o k
*/
type UserControlJobOK struct {
	Payload *models.JobsCtrlCommandResponse
}

func (o *UserControlJobOK) Error() string {
	return fmt.Sprintf("[PUT /jobs/user][%d] userControlJobOK  %+v", 200, o.Payload)
}

func (o *UserControlJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobsCtrlCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}