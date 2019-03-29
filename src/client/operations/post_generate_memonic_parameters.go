// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostGenerateMemonicParams creates a new PostGenerateMemonicParams object
// with the default values initialized.
func NewPostGenerateMemonicParams() *PostGenerateMemonicParams {
	var ()
	return &PostGenerateMemonicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGenerateMemonicParamsWithTimeout creates a new PostGenerateMemonicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGenerateMemonicParamsWithTimeout(timeout time.Duration) *PostGenerateMemonicParams {
	var ()
	return &PostGenerateMemonicParams{

		timeout: timeout,
	}
}

// NewPostGenerateMemonicParamsWithContext creates a new PostGenerateMemonicParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGenerateMemonicParamsWithContext(ctx context.Context) *PostGenerateMemonicParams {
	var ()
	return &PostGenerateMemonicParams{

		Context: ctx,
	}
}

// NewPostGenerateMemonicParamsWithHTTPClient creates a new PostGenerateMemonicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGenerateMemonicParamsWithHTTPClient(client *http.Client) *PostGenerateMemonicParams {
	var ()
	return &PostGenerateMemonicParams{
		HTTPClient: client,
	}
}

/*PostGenerateMemonicParams contains all the parameters to send to the API endpoint
for the post generate memonic operation typically these are written to a http.Request
*/
type PostGenerateMemonicParams struct {

	/*UsePassphrase
	  ask for passphrase before starting operation

	*/
	UsePassphrase *bool
	/*WordCount
	  mnemonic seed length

	*/
	WordCount int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post generate memonic params
func (o *PostGenerateMemonicParams) WithTimeout(timeout time.Duration) *PostGenerateMemonicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post generate memonic params
func (o *PostGenerateMemonicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post generate memonic params
func (o *PostGenerateMemonicParams) WithContext(ctx context.Context) *PostGenerateMemonicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post generate memonic params
func (o *PostGenerateMemonicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post generate memonic params
func (o *PostGenerateMemonicParams) WithHTTPClient(client *http.Client) *PostGenerateMemonicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post generate memonic params
func (o *PostGenerateMemonicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsePassphrase adds the usePassphrase to the post generate memonic params
func (o *PostGenerateMemonicParams) WithUsePassphrase(usePassphrase *bool) *PostGenerateMemonicParams {
	o.SetUsePassphrase(usePassphrase)
	return o
}

// SetUsePassphrase adds the usePassphrase to the post generate memonic params
func (o *PostGenerateMemonicParams) SetUsePassphrase(usePassphrase *bool) {
	o.UsePassphrase = usePassphrase
}

// WithWordCount adds the wordCount to the post generate memonic params
func (o *PostGenerateMemonicParams) WithWordCount(wordCount int64) *PostGenerateMemonicParams {
	o.SetWordCount(wordCount)
	return o
}

// SetWordCount adds the wordCount to the post generate memonic params
func (o *PostGenerateMemonicParams) SetWordCount(wordCount int64) {
	o.WordCount = wordCount
}

// WriteToRequest writes these params to a swagger request
func (o *PostGenerateMemonicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UsePassphrase != nil {

		// form param use-passphrase
		var frUsePassphrase bool
		if o.UsePassphrase != nil {
			frUsePassphrase = *o.UsePassphrase
		}
		fUsePassphrase := swag.FormatBool(frUsePassphrase)
		if fUsePassphrase != "" {
			if err := r.SetFormParam("use-passphrase", fUsePassphrase); err != nil {
				return err
			}
		}

	}

	// form param word-count
	frWordCount := o.WordCount
	fWordCount := swag.FormatInt64(frWordCount)
	if fWordCount != "" {
		if err := r.SetFormParam("word-count", fWordCount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
