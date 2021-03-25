// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVersioningPolicyParams creates a new GetVersioningPolicyParams object
// with the default values initialized.
func NewGetVersioningPolicyParams() *GetVersioningPolicyParams {
	var ()
	return &GetVersioningPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersioningPolicyParamsWithTimeout creates a new GetVersioningPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersioningPolicyParamsWithTimeout(timeout time.Duration) *GetVersioningPolicyParams {
	var ()
	return &GetVersioningPolicyParams{

		timeout: timeout,
	}
}

// NewGetVersioningPolicyParamsWithContext creates a new GetVersioningPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersioningPolicyParamsWithContext(ctx context.Context) *GetVersioningPolicyParams {
	var ()
	return &GetVersioningPolicyParams{

		Context: ctx,
	}
}

// NewGetVersioningPolicyParamsWithHTTPClient creates a new GetVersioningPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersioningPolicyParamsWithHTTPClient(client *http.Client) *GetVersioningPolicyParams {
	var ()
	return &GetVersioningPolicyParams{
		HTTPClient: client,
	}
}

/*GetVersioningPolicyParams contains all the parameters to send to the API endpoint
for the get versioning policy operation typically these are written to a http.Request
*/
type GetVersioningPolicyParams struct {

	/*Description*/
	Description *string
	/*IgnoreFilesGreaterThan*/
	IgnoreFilesGreaterThan *string
	/*MaxSizePerFile*/
	MaxSizePerFile *string
	/*MaxTotalSize*/
	MaxTotalSize *string
	/*Name*/
	Name *string
	/*UUID*/
	UUID string
	/*VersionsDataSourceBucket*/
	VersionsDataSourceBucket *string
	/*VersionsDataSourceName*/
	VersionsDataSourceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get versioning policy params
func (o *GetVersioningPolicyParams) WithTimeout(timeout time.Duration) *GetVersioningPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versioning policy params
func (o *GetVersioningPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versioning policy params
func (o *GetVersioningPolicyParams) WithContext(ctx context.Context) *GetVersioningPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versioning policy params
func (o *GetVersioningPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versioning policy params
func (o *GetVersioningPolicyParams) WithHTTPClient(client *http.Client) *GetVersioningPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versioning policy params
func (o *GetVersioningPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the get versioning policy params
func (o *GetVersioningPolicyParams) WithDescription(description *string) *GetVersioningPolicyParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get versioning policy params
func (o *GetVersioningPolicyParams) SetDescription(description *string) {
	o.Description = description
}

// WithIgnoreFilesGreaterThan adds the ignoreFilesGreaterThan to the get versioning policy params
func (o *GetVersioningPolicyParams) WithIgnoreFilesGreaterThan(ignoreFilesGreaterThan *string) *GetVersioningPolicyParams {
	o.SetIgnoreFilesGreaterThan(ignoreFilesGreaterThan)
	return o
}

// SetIgnoreFilesGreaterThan adds the ignoreFilesGreaterThan to the get versioning policy params
func (o *GetVersioningPolicyParams) SetIgnoreFilesGreaterThan(ignoreFilesGreaterThan *string) {
	o.IgnoreFilesGreaterThan = ignoreFilesGreaterThan
}

// WithMaxSizePerFile adds the maxSizePerFile to the get versioning policy params
func (o *GetVersioningPolicyParams) WithMaxSizePerFile(maxSizePerFile *string) *GetVersioningPolicyParams {
	o.SetMaxSizePerFile(maxSizePerFile)
	return o
}

// SetMaxSizePerFile adds the maxSizePerFile to the get versioning policy params
func (o *GetVersioningPolicyParams) SetMaxSizePerFile(maxSizePerFile *string) {
	o.MaxSizePerFile = maxSizePerFile
}

// WithMaxTotalSize adds the maxTotalSize to the get versioning policy params
func (o *GetVersioningPolicyParams) WithMaxTotalSize(maxTotalSize *string) *GetVersioningPolicyParams {
	o.SetMaxTotalSize(maxTotalSize)
	return o
}

// SetMaxTotalSize adds the maxTotalSize to the get versioning policy params
func (o *GetVersioningPolicyParams) SetMaxTotalSize(maxTotalSize *string) {
	o.MaxTotalSize = maxTotalSize
}

// WithName adds the name to the get versioning policy params
func (o *GetVersioningPolicyParams) WithName(name *string) *GetVersioningPolicyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get versioning policy params
func (o *GetVersioningPolicyParams) SetName(name *string) {
	o.Name = name
}

// WithUUID adds the uuid to the get versioning policy params
func (o *GetVersioningPolicyParams) WithUUID(uuid string) *GetVersioningPolicyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get versioning policy params
func (o *GetVersioningPolicyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WithVersionsDataSourceBucket adds the versionsDataSourceBucket to the get versioning policy params
func (o *GetVersioningPolicyParams) WithVersionsDataSourceBucket(versionsDataSourceBucket *string) *GetVersioningPolicyParams {
	o.SetVersionsDataSourceBucket(versionsDataSourceBucket)
	return o
}

// SetVersionsDataSourceBucket adds the versionsDataSourceBucket to the get versioning policy params
func (o *GetVersioningPolicyParams) SetVersionsDataSourceBucket(versionsDataSourceBucket *string) {
	o.VersionsDataSourceBucket = versionsDataSourceBucket
}

// WithVersionsDataSourceName adds the versionsDataSourceName to the get versioning policy params
func (o *GetVersioningPolicyParams) WithVersionsDataSourceName(versionsDataSourceName *string) *GetVersioningPolicyParams {
	o.SetVersionsDataSourceName(versionsDataSourceName)
	return o
}

// SetVersionsDataSourceName adds the versionsDataSourceName to the get versioning policy params
func (o *GetVersioningPolicyParams) SetVersionsDataSourceName(versionsDataSourceName *string) {
	o.VersionsDataSourceName = versionsDataSourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersioningPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param Description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("Description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.IgnoreFilesGreaterThan != nil {

		// query param IgnoreFilesGreaterThan
		var qrIgnoreFilesGreaterThan string
		if o.IgnoreFilesGreaterThan != nil {
			qrIgnoreFilesGreaterThan = *o.IgnoreFilesGreaterThan
		}
		qIgnoreFilesGreaterThan := qrIgnoreFilesGreaterThan
		if qIgnoreFilesGreaterThan != "" {
			if err := r.SetQueryParam("IgnoreFilesGreaterThan", qIgnoreFilesGreaterThan); err != nil {
				return err
			}
		}

	}

	if o.MaxSizePerFile != nil {

		// query param MaxSizePerFile
		var qrMaxSizePerFile string
		if o.MaxSizePerFile != nil {
			qrMaxSizePerFile = *o.MaxSizePerFile
		}
		qMaxSizePerFile := qrMaxSizePerFile
		if qMaxSizePerFile != "" {
			if err := r.SetQueryParam("MaxSizePerFile", qMaxSizePerFile); err != nil {
				return err
			}
		}

	}

	if o.MaxTotalSize != nil {

		// query param MaxTotalSize
		var qrMaxTotalSize string
		if o.MaxTotalSize != nil {
			qrMaxTotalSize = *o.MaxTotalSize
		}
		qMaxTotalSize := qrMaxTotalSize
		if qMaxTotalSize != "" {
			if err := r.SetQueryParam("MaxTotalSize", qMaxTotalSize); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param Name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("Name", qName); err != nil {
				return err
			}
		}

	}

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if o.VersionsDataSourceBucket != nil {

		// query param VersionsDataSourceBucket
		var qrVersionsDataSourceBucket string
		if o.VersionsDataSourceBucket != nil {
			qrVersionsDataSourceBucket = *o.VersionsDataSourceBucket
		}
		qVersionsDataSourceBucket := qrVersionsDataSourceBucket
		if qVersionsDataSourceBucket != "" {
			if err := r.SetQueryParam("VersionsDataSourceBucket", qVersionsDataSourceBucket); err != nil {
				return err
			}
		}

	}

	if o.VersionsDataSourceName != nil {

		// query param VersionsDataSourceName
		var qrVersionsDataSourceName string
		if o.VersionsDataSourceName != nil {
			qrVersionsDataSourceName = *o.VersionsDataSourceName
		}
		qVersionsDataSourceName := qrVersionsDataSourceName
		if qVersionsDataSourceName != "" {
			if err := r.SetQueryParam("VersionsDataSourceName", qVersionsDataSourceName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
