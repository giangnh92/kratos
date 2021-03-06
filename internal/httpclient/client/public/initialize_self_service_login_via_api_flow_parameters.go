// Code generated by go-swagger; DO NOT EDIT.

package public

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
	"github.com/go-openapi/swag"
)

// NewInitializeSelfServiceLoginViaAPIFlowParams creates a new InitializeSelfServiceLoginViaAPIFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeSelfServiceLoginViaAPIFlowParams() *InitializeSelfServiceLoginViaAPIFlowParams {
	return &InitializeSelfServiceLoginViaAPIFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceLoginViaAPIFlowParamsWithTimeout creates a new InitializeSelfServiceLoginViaAPIFlowParams object
// with the ability to set a timeout on a request.
func NewInitializeSelfServiceLoginViaAPIFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceLoginViaAPIFlowParams {
	return &InitializeSelfServiceLoginViaAPIFlowParams{
		timeout: timeout,
	}
}

// NewInitializeSelfServiceLoginViaAPIFlowParamsWithContext creates a new InitializeSelfServiceLoginViaAPIFlowParams object
// with the ability to set a context for a request.
func NewInitializeSelfServiceLoginViaAPIFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceLoginViaAPIFlowParams {
	return &InitializeSelfServiceLoginViaAPIFlowParams{
		Context: ctx,
	}
}

// NewInitializeSelfServiceLoginViaAPIFlowParamsWithHTTPClient creates a new InitializeSelfServiceLoginViaAPIFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeSelfServiceLoginViaAPIFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceLoginViaAPIFlowParams {
	return &InitializeSelfServiceLoginViaAPIFlowParams{
		HTTPClient: client,
	}
}

/* InitializeSelfServiceLoginViaAPIFlowParams contains all the parameters to send to the API endpoint
   for the initialize self service login via API flow operation.

   Typically these are written to a http.Request.
*/
type InitializeSelfServiceLoginViaAPIFlowParams struct {

	/* Refresh.

	     Refresh a login session

	If set to true, this will refresh an existing login session by
	asking the user to sign in again. This will reset the
	authenticated_at time of the session.
	*/
	Refresh *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize self service login via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WithDefaults() *InitializeSelfServiceLoginViaAPIFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize self service login via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceLoginViaAPIFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceLoginViaAPIFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceLoginViaAPIFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceLoginViaAPIFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefresh adds the refresh to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WithRefresh(refresh *bool) *InitializeSelfServiceLoginViaAPIFlowParams {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the initialize self service login via API flow params
func (o *InitializeSelfServiceLoginViaAPIFlowParams) SetRefresh(refresh *bool) {
	o.Refresh = refresh
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceLoginViaAPIFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Refresh != nil {

		// query param refresh
		var qrRefresh bool

		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := swag.FormatBool(qrRefresh)
		if qRefresh != "" {

			if err := r.SetQueryParam("refresh", qRefresh); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
