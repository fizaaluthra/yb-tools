// Code generated by go-swagger; DO NOT EDIT.

package customer_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// GenerateAPITokenReader is a Reader for the GenerateAPIToken structure.
type GenerateAPITokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateAPITokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateAPITokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateAPITokenOK creates a GenerateAPITokenOK with default headers values
func NewGenerateAPITokenOK() *GenerateAPITokenOK {
	return &GenerateAPITokenOK{}
}

/* GenerateAPITokenOK describes a response with status code 200, with default header values.

successful operation
*/
type GenerateAPITokenOK struct {
	Payload *models.SessionInfo
}

func (o *GenerateAPITokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/customers/{cUUID}/api_token][%d] generateApiTokenOK  %+v", 200, o.Payload)
}
func (o *GenerateAPITokenOK) GetPayload() *models.SessionInfo {
	return o.Payload
}

func (o *GenerateAPITokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
