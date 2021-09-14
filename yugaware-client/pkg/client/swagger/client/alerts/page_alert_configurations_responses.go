// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// PageAlertConfigurationsReader is a Reader for the PageAlertConfigurations structure.
type PageAlertConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PageAlertConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPageAlertConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPageAlertConfigurationsOK creates a PageAlertConfigurationsOK with default headers values
func NewPageAlertConfigurationsOK() *PageAlertConfigurationsOK {
	return &PageAlertConfigurationsOK{}
}

/* PageAlertConfigurationsOK describes a response with status code 200, with default header values.

successful operation
*/
type PageAlertConfigurationsOK struct {
	Payload *models.AlertConfigurationPagedResponse
}

func (o *PageAlertConfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/customers/{cUUID}/alert_configurations/page][%d] pageAlertConfigurationsOK  %+v", 200, o.Payload)
}
func (o *PageAlertConfigurationsOK) GetPayload() *models.AlertConfigurationPagedResponse {
	return o.Payload
}

func (o *PageAlertConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertConfigurationPagedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}