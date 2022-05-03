// Code generated by go-swagger; DO NOT EDIT.

package asynchronous_replication

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

// NewSyncXClusterConfigParams creates a new SyncXClusterConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSyncXClusterConfigParams() *SyncXClusterConfigParams {
	return &SyncXClusterConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSyncXClusterConfigParamsWithTimeout creates a new SyncXClusterConfigParams object
// with the ability to set a timeout on a request.
func NewSyncXClusterConfigParamsWithTimeout(timeout time.Duration) *SyncXClusterConfigParams {
	return &SyncXClusterConfigParams{
		timeout: timeout,
	}
}

// NewSyncXClusterConfigParamsWithContext creates a new SyncXClusterConfigParams object
// with the ability to set a context for a request.
func NewSyncXClusterConfigParamsWithContext(ctx context.Context) *SyncXClusterConfigParams {
	return &SyncXClusterConfigParams{
		Context: ctx,
	}
}

// NewSyncXClusterConfigParamsWithHTTPClient creates a new SyncXClusterConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewSyncXClusterConfigParamsWithHTTPClient(client *http.Client) *SyncXClusterConfigParams {
	return &SyncXClusterConfigParams{
		HTTPClient: client,
	}
}

/* SyncXClusterConfigParams contains all the parameters to send to the API endpoint
   for the sync x cluster config operation.

   Typically these are written to a http.Request.
*/
type SyncXClusterConfigParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// TargetUniverseUUID.
	//
	// Format: uuid
	TargetUniverseUUID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sync x cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncXClusterConfigParams) WithDefaults() *SyncXClusterConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sync x cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncXClusterConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sync x cluster config params
func (o *SyncXClusterConfigParams) WithTimeout(timeout time.Duration) *SyncXClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync x cluster config params
func (o *SyncXClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync x cluster config params
func (o *SyncXClusterConfigParams) WithContext(ctx context.Context) *SyncXClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync x cluster config params
func (o *SyncXClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync x cluster config params
func (o *SyncXClusterConfigParams) WithHTTPClient(client *http.Client) *SyncXClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync x cluster config params
func (o *SyncXClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the sync x cluster config params
func (o *SyncXClusterConfigParams) WithCUUID(cUUID strfmt.UUID) *SyncXClusterConfigParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the sync x cluster config params
func (o *SyncXClusterConfigParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithTargetUniverseUUID adds the targetUniverseUUID to the sync x cluster config params
func (o *SyncXClusterConfigParams) WithTargetUniverseUUID(targetUniverseUUID *strfmt.UUID) *SyncXClusterConfigParams {
	o.SetTargetUniverseUUID(targetUniverseUUID)
	return o
}

// SetTargetUniverseUUID adds the targetUniverseUuid to the sync x cluster config params
func (o *SyncXClusterConfigParams) SetTargetUniverseUUID(targetUniverseUUID *strfmt.UUID) {
	o.TargetUniverseUUID = targetUniverseUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SyncXClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if o.TargetUniverseUUID != nil {

		// query param targetUniverseUUID
		var qrTargetUniverseUUID strfmt.UUID

		if o.TargetUniverseUUID != nil {
			qrTargetUniverseUUID = *o.TargetUniverseUUID
		}
		qTargetUniverseUUID := qrTargetUniverseUUID.String()
		if qTargetUniverseUUID != "" {

			if err := r.SetQueryParam("targetUniverseUUID", qTargetUniverseUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}