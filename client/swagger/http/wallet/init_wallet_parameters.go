// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewInitWalletParams creates a new InitWalletParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitWalletParams() *InitWalletParams {
	return &InitWalletParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitWalletParamsWithTimeout creates a new InitWalletParams object
// with the ability to set a timeout on a request.
func NewInitWalletParamsWithTimeout(timeout time.Duration) *InitWalletParams {
	return &InitWalletParams{
		timeout: timeout,
	}
}

// NewInitWalletParamsWithContext creates a new InitWalletParams object
// with the ability to set a context for a request.
func NewInitWalletParamsWithContext(ctx context.Context) *InitWalletParams {
	return &InitWalletParams{
		Context: ctx,
	}
}

// NewInitWalletParamsWithHTTPClient creates a new InitWalletParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitWalletParamsWithHTTPClient(client *http.Client) *InitWalletParams {
	return &InitWalletParams{
		HTTPClient: client,
	}
}

/*
InitWalletParams contains all the parameters to send to the API endpoint

	for the init wallet operation.

	Typically these are written to a http.Request.
*/
type InitWalletParams struct {

	/* Address.

	   Address
	*/
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the init wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitWalletParams) WithDefaults() *InitWalletParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the init wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitWalletParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the init wallet params
func (o *InitWalletParams) WithTimeout(timeout time.Duration) *InitWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the init wallet params
func (o *InitWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the init wallet params
func (o *InitWalletParams) WithContext(ctx context.Context) *InitWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the init wallet params
func (o *InitWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the init wallet params
func (o *InitWalletParams) WithHTTPClient(client *http.Client) *InitWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the init wallet params
func (o *InitWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the init wallet params
func (o *InitWalletParams) WithAddress(address string) *InitWalletParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the init wallet params
func (o *InitWalletParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *InitWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
