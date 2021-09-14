// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAlertConfigurationParams creates a new GetAlertConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertConfigurationParams() *GetAlertConfigurationParams {
	return &GetAlertConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertConfigurationParamsWithTimeout creates a new GetAlertConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetAlertConfigurationParamsWithTimeout(timeout time.Duration) *GetAlertConfigurationParams {
	return &GetAlertConfigurationParams{
		timeout: timeout,
	}
}

// NewGetAlertConfigurationParamsWithContext creates a new GetAlertConfigurationParams object
// with the ability to set a context for a request.
func NewGetAlertConfigurationParamsWithContext(ctx context.Context) *GetAlertConfigurationParams {
	return &GetAlertConfigurationParams{
		Context: ctx,
	}
}

// NewGetAlertConfigurationParamsWithHTTPClient creates a new GetAlertConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertConfigurationParamsWithHTTPClient(client *http.Client) *GetAlertConfigurationParams {
	return &GetAlertConfigurationParams{
		HTTPClient: client,
	}
}

/* GetAlertConfigurationParams contains all the parameters to send to the API endpoint
   for the get alert configuration operation.

   Typically these are written to a http.Request.
*/
type GetAlertConfigurationParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// ConfigurationUUID.
	//
	// Format: uuid
	ConfigurationUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertConfigurationParams) WithDefaults() *GetAlertConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alert configuration params
func (o *GetAlertConfigurationParams) WithTimeout(timeout time.Duration) *GetAlertConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert configuration params
func (o *GetAlertConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert configuration params
func (o *GetAlertConfigurationParams) WithContext(ctx context.Context) *GetAlertConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert configuration params
func (o *GetAlertConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert configuration params
func (o *GetAlertConfigurationParams) WithHTTPClient(client *http.Client) *GetAlertConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert configuration params
func (o *GetAlertConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get alert configuration params
func (o *GetAlertConfigurationParams) WithCUUID(cUUID strfmt.UUID) *GetAlertConfigurationParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get alert configuration params
func (o *GetAlertConfigurationParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithConfigurationUUID adds the configurationUUID to the get alert configuration params
func (o *GetAlertConfigurationParams) WithConfigurationUUID(configurationUUID strfmt.UUID) *GetAlertConfigurationParams {
	o.SetConfigurationUUID(configurationUUID)
	return o
}

// SetConfigurationUUID adds the configurationUuid to the get alert configuration params
func (o *GetAlertConfigurationParams) SetConfigurationUUID(configurationUUID strfmt.UUID) {
	o.ConfigurationUUID = configurationUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param configurationUUID
	if err := r.SetPathParam("configurationUUID", o.ConfigurationUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}