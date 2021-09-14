// Code generated by go-swagger; DO NOT EDIT.

package encryption_at_rest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RetrieveKeyReader is a Reader for the RetrieveKey structure.
type RetrieveKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveKeyOK creates a RetrieveKeyOK with default headers values
func NewRetrieveKeyOK() *RetrieveKeyOK {
	return &RetrieveKeyOK{}
}

/* RetrieveKeyOK describes a response with status code 200, with default header values.

successful operation
*/
type RetrieveKeyOK struct {
	Payload map[string]interface{}
}

func (o *RetrieveKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/customers/{cUUID}/universes/{uniUUID}/kms][%d] retrieveKeyOK  %+v", 200, o.Payload)
}
func (o *RetrieveKeyOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *RetrieveKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}