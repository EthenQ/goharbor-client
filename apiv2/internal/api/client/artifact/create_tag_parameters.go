// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// NewCreateTagParams creates a new CreateTagParams object
// with the default values initialized.
func NewCreateTagParams() *CreateTagParams {
	var ()
	return &CreateTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTagParamsWithTimeout creates a new CreateTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTagParamsWithTimeout(timeout time.Duration) *CreateTagParams {
	var ()
	return &CreateTagParams{

		timeout: timeout,
	}
}

// NewCreateTagParamsWithContext creates a new CreateTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTagParamsWithContext(ctx context.Context) *CreateTagParams {
	var ()
	return &CreateTagParams{

		Context: ctx,
	}
}

// NewCreateTagParamsWithHTTPClient creates a new CreateTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTagParamsWithHTTPClient(client *http.Client) *CreateTagParams {
	var ()
	return &CreateTagParams{
		HTTPClient: client,
	}
}

/*CreateTagParams contains all the parameters to send to the API endpoint
for the create tag operation typically these are written to a http.Request
*/
type CreateTagParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectName
	  The name of the project

	*/
	ProjectName string
	/*Reference
	  The reference of the artifact, can be digest or tag

	*/
	Reference string
	/*RepositoryName
	  The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb

	*/
	RepositoryName string
	/*Tag
	  The JSON object of tag.

	*/
	Tag *model.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create tag params
func (o *CreateTagParams) WithTimeout(timeout time.Duration) *CreateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tag params
func (o *CreateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tag params
func (o *CreateTagParams) WithContext(ctx context.Context) *CreateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tag params
func (o *CreateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) WithHTTPClient(client *http.Client) *CreateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create tag params
func (o *CreateTagParams) WithXRequestID(xRequestID *string) *CreateTagParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create tag params
func (o *CreateTagParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the create tag params
func (o *CreateTagParams) WithProjectName(projectName string) *CreateTagParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the create tag params
func (o *CreateTagParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the create tag params
func (o *CreateTagParams) WithReference(reference string) *CreateTagParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the create tag params
func (o *CreateTagParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the create tag params
func (o *CreateTagParams) WithRepositoryName(repositoryName string) *CreateTagParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the create tag params
func (o *CreateTagParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WithTag adds the tag to the create tag params
func (o *CreateTagParams) WithTag(tag *model.Tag) *CreateTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the create tag params
func (o *CreateTagParams) SetTag(tag *model.Tag) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
