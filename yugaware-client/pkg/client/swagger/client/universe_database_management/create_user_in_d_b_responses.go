// Code generated by go-swagger; DO NOT EDIT.

package universe_database_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// CreateUserInDBReader is a Reader for the CreateUserInDB structure.
type CreateUserInDBReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserInDBReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserInDBOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserInDBOK creates a CreateUserInDBOK with default headers values
func NewCreateUserInDBOK() *CreateUserInDBOK {
	return &CreateUserInDBOK{}
}

/* CreateUserInDBOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateUserInDBOK struct {
	Payload *models.YBPSuccess
}

func (o *CreateUserInDBOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/customers/{cUUID}/universes/{uniUUID}/create_db_credentials][%d] createUserInDBOK  %+v", 200, o.Payload)
}
func (o *CreateUserInDBOK) GetPayload() *models.YBPSuccess {
	return o.Payload
}

func (o *CreateUserInDBOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPSuccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}