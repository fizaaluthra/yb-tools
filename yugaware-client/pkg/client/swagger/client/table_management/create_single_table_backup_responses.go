// Code generated by go-swagger; DO NOT EDIT.

package table_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// CreateSingleTableBackupReader is a Reader for the CreateSingleTableBackup structure.
type CreateSingleTableBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSingleTableBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSingleTableBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSingleTableBackupOK creates a CreateSingleTableBackupOK with default headers values
func NewCreateSingleTableBackupOK() *CreateSingleTableBackupOK {
	return &CreateSingleTableBackupOK{}
}

/* CreateSingleTableBackupOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateSingleTableBackupOK struct {
	Payload *models.YBPTask
}

func (o *CreateSingleTableBackupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/create_backup][%d] createSingleTableBackupOK  %+v", 200, o.Payload)
}
func (o *CreateSingleTableBackupOK) GetPayload() *models.YBPTask {
	return o.Payload
}

func (o *CreateSingleTableBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}