// Code generated by go-swagger; DO NOT EDIT.

package universe_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSlowQueriesReader is a Reader for the GetSlowQueries structure.
type GetSlowQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSlowQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSlowQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSlowQueriesOK creates a GetSlowQueriesOK with default headers values
func NewGetSlowQueriesOK() *GetSlowQueriesOK {
	return &GetSlowQueriesOK{}
}

/* GetSlowQueriesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSlowQueriesOK struct {
	Payload interface{}
}

func (o *GetSlowQueriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/universes/{uniUUID}/slow_queries][%d] getSlowQueriesOK  %+v", 200, o.Payload)
}
func (o *GetSlowQueriesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetSlowQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}