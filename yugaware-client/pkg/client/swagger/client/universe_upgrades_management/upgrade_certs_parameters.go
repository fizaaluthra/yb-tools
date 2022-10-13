// Code generated by go-swagger; DO NOT EDIT.

package universe_upgrades_management

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

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// NewUpgradeCertsParams creates a new UpgradeCertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeCertsParams() *UpgradeCertsParams {
	return &UpgradeCertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCertsParamsWithTimeout creates a new UpgradeCertsParams object
// with the ability to set a timeout on a request.
func NewUpgradeCertsParamsWithTimeout(timeout time.Duration) *UpgradeCertsParams {
	return &UpgradeCertsParams{
		timeout: timeout,
	}
}

// NewUpgradeCertsParamsWithContext creates a new UpgradeCertsParams object
// with the ability to set a context for a request.
func NewUpgradeCertsParamsWithContext(ctx context.Context) *UpgradeCertsParams {
	return &UpgradeCertsParams{
		Context: ctx,
	}
}

// NewUpgradeCertsParamsWithHTTPClient creates a new UpgradeCertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeCertsParamsWithHTTPClient(client *http.Client) *UpgradeCertsParams {
	return &UpgradeCertsParams{
		HTTPClient: client,
	}
}

/* UpgradeCertsParams contains all the parameters to send to the API endpoint
   for the upgrade certs operation.

   Typically these are written to a http.Request.
*/
type UpgradeCertsParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	/* CertsRotateParams.

	   Certs Rotate Params
	*/
	CertsRotateParams *models.CertsRotateParams

	// UniUUID.
	//
	// Format: uuid
	UniUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade certs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCertsParams) WithDefaults() *UpgradeCertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade certs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade certs params
func (o *UpgradeCertsParams) WithTimeout(timeout time.Duration) *UpgradeCertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade certs params
func (o *UpgradeCertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade certs params
func (o *UpgradeCertsParams) WithContext(ctx context.Context) *UpgradeCertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade certs params
func (o *UpgradeCertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade certs params
func (o *UpgradeCertsParams) WithHTTPClient(client *http.Client) *UpgradeCertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade certs params
func (o *UpgradeCertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the upgrade certs params
func (o *UpgradeCertsParams) WithCUUID(cUUID strfmt.UUID) *UpgradeCertsParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the upgrade certs params
func (o *UpgradeCertsParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithCertsRotateParams adds the certsRotateParams to the upgrade certs params
func (o *UpgradeCertsParams) WithCertsRotateParams(certsRotateParams *models.CertsRotateParams) *UpgradeCertsParams {
	o.SetCertsRotateParams(certsRotateParams)
	return o
}

// SetCertsRotateParams adds the certsRotateParams to the upgrade certs params
func (o *UpgradeCertsParams) SetCertsRotateParams(certsRotateParams *models.CertsRotateParams) {
	o.CertsRotateParams = certsRotateParams
}

// WithUniUUID adds the uniUUID to the upgrade certs params
func (o *UpgradeCertsParams) WithUniUUID(uniUUID strfmt.UUID) *UpgradeCertsParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the upgrade certs params
func (o *UpgradeCertsParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}
	if o.CertsRotateParams != nil {
		if err := r.SetBodyParam(o.CertsRotateParams); err != nil {
			return err
		}
	}

	// path param uniUUID
	if err := r.SetPathParam("uniUUID", o.UniUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}