// Code generated by go-swagger; DO NOT EDIT.

package metadata

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

// NewGetShareMetricsParams creates a new GetShareMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetShareMetricsParams() *GetShareMetricsParams {
	return &GetShareMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetShareMetricsParamsWithTimeout creates a new GetShareMetricsParams object
// with the ability to set a timeout on a request.
func NewGetShareMetricsParamsWithTimeout(timeout time.Duration) *GetShareMetricsParams {
	return &GetShareMetricsParams{
		timeout: timeout,
	}
}

// NewGetShareMetricsParamsWithContext creates a new GetShareMetricsParams object
// with the ability to set a context for a request.
func NewGetShareMetricsParamsWithContext(ctx context.Context) *GetShareMetricsParams {
	return &GetShareMetricsParams{
		Context: ctx,
	}
}

// NewGetShareMetricsParamsWithHTTPClient creates a new GetShareMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetShareMetricsParamsWithHTTPClient(client *http.Client) *GetShareMetricsParams {
	return &GetShareMetricsParams{
		HTTPClient: client,
	}
}

/*
GetShareMetricsParams contains all the parameters to send to the API endpoint

	for the get share metrics operation.

	Typically these are written to a http.Request.
*/
type GetShareMetricsParams struct {

	// Duration.
	Duration *string

	// ShrToken.
	ShrToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get share metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShareMetricsParams) WithDefaults() *GetShareMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get share metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShareMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get share metrics params
func (o *GetShareMetricsParams) WithTimeout(timeout time.Duration) *GetShareMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get share metrics params
func (o *GetShareMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get share metrics params
func (o *GetShareMetricsParams) WithContext(ctx context.Context) *GetShareMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get share metrics params
func (o *GetShareMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get share metrics params
func (o *GetShareMetricsParams) WithHTTPClient(client *http.Client) *GetShareMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get share metrics params
func (o *GetShareMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDuration adds the duration to the get share metrics params
func (o *GetShareMetricsParams) WithDuration(duration *string) *GetShareMetricsParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the get share metrics params
func (o *GetShareMetricsParams) SetDuration(duration *string) {
	o.Duration = duration
}

// WithShrToken adds the shrToken to the get share metrics params
func (o *GetShareMetricsParams) WithShrToken(shrToken string) *GetShareMetricsParams {
	o.SetShrToken(shrToken)
	return o
}

// SetShrToken adds the shrToken to the get share metrics params
func (o *GetShareMetricsParams) SetShrToken(shrToken string) {
	o.ShrToken = shrToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetShareMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Duration != nil {

		// query param duration
		var qrDuration string

		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := qrDuration
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}
	}

	// path param shrToken
	if err := r.SetPathParam("shrToken", o.ShrToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
