// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// UpdateSharePoliciesReader is a Reader for the UpdateSharePolicies structure.
type UpdateSharePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSharePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSharePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSharePoliciesOK creates a UpdateSharePoliciesOK with default headers values
func NewUpdateSharePoliciesOK() *UpdateSharePoliciesOK {
	return &UpdateSharePoliciesOK{}
}

/*UpdateSharePoliciesOK handles this case with default header values.

UpdateSharePoliciesOK update share policies o k
*/
type UpdateSharePoliciesOK struct {
	Payload *models.RestUpdateSharePoliciesResponse
}

func (o *UpdateSharePoliciesOK) Error() string {
	return fmt.Sprintf("[PUT /share/policies][%d] updateSharePoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdateSharePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUpdateSharePoliciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
