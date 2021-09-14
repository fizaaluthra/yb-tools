// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SMTPData SMTP configuration information
//
// swagger:model SmtpData
type SMTPData struct {

	// SMTP email 'from' address
	// Example: test@example.com
	EmailFrom string `json:"emailFrom,omitempty"`

	// SMTP password
	// Example: XurenRknsc
	SMTPPassword string `json:"smtpPassword,omitempty"`

	// SMTP port number
	// Example: 465
	SMTPPort int32 `json:"smtpPort,omitempty"`

	// SMTP server
	// Example: smtp.example.com
	SMTPServer string `json:"smtpServer,omitempty"`

	// SMTP email username
	// Example: testsmtp
	SMTPUsername string `json:"smtpUsername,omitempty"`

	// Connect to SMTP server using SSL
	// Example: true
	UseSSL bool `json:"useSSL,omitempty"`

	// Connect to SMTP server using TLS
	// Example: false
	UseTLS bool `json:"useTLS,omitempty"`
}

// Validate validates this Smtp data
func (m *SMTPData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Smtp data based on context it is used
func (m *SMTPData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SMTPData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SMTPData) UnmarshalBinary(b []byte) error {
	var res SMTPData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}