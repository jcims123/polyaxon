// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

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

// NewGetRunArtifactsTreeParams creates a new GetRunArtifactsTreeParams object
// with the default values initialized.
func NewGetRunArtifactsTreeParams() *GetRunArtifactsTreeParams {
	var ()
	return &GetRunArtifactsTreeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunArtifactsTreeParamsWithTimeout creates a new GetRunArtifactsTreeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunArtifactsTreeParamsWithTimeout(timeout time.Duration) *GetRunArtifactsTreeParams {
	var ()
	return &GetRunArtifactsTreeParams{

		timeout: timeout,
	}
}

// NewGetRunArtifactsTreeParamsWithContext creates a new GetRunArtifactsTreeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunArtifactsTreeParamsWithContext(ctx context.Context) *GetRunArtifactsTreeParams {
	var ()
	return &GetRunArtifactsTreeParams{

		Context: ctx,
	}
}

// NewGetRunArtifactsTreeParamsWithHTTPClient creates a new GetRunArtifactsTreeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunArtifactsTreeParamsWithHTTPClient(client *http.Client) *GetRunArtifactsTreeParams {
	var ()
	return &GetRunArtifactsTreeParams{
		HTTPClient: client,
	}
}

/*GetRunArtifactsTreeParams contains all the parameters to send to the API endpoint
for the get run artifacts tree operation typically these are written to a http.Request
*/
type GetRunArtifactsTreeParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Path
	  Path query param.

	*/
	Path *string
	/*Project
	  Project where the run will be assigned

	*/
	Project string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithTimeout(timeout time.Duration) *GetRunArtifactsTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithContext(ctx context.Context) *GetRunArtifactsTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithHTTPClient(client *http.Client) *GetRunArtifactsTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithNamespace(namespace string) *GetRunArtifactsTreeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithOwner(owner string) *GetRunArtifactsTreeParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPath adds the path to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithPath(path *string) *GetRunArtifactsTreeParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetPath(path *string) {
	o.Path = path
}

// WithProject adds the project to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithProject(project string) *GetRunArtifactsTreeParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) WithUUID(uuid string) *GetRunArtifactsTreeParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run artifacts tree params
func (o *GetRunArtifactsTreeParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunArtifactsTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string
		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {
			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}

	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
