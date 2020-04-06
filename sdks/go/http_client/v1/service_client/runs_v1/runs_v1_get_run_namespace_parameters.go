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

// NewRunsV1GetRunNamespaceParams creates a new RunsV1GetRunNamespaceParams object
// with the default values initialized.
func NewRunsV1GetRunNamespaceParams() *RunsV1GetRunNamespaceParams {
	var ()
	return &RunsV1GetRunNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1GetRunNamespaceParamsWithTimeout creates a new RunsV1GetRunNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1GetRunNamespaceParamsWithTimeout(timeout time.Duration) *RunsV1GetRunNamespaceParams {
	var ()
	return &RunsV1GetRunNamespaceParams{

		timeout: timeout,
	}
}

// NewRunsV1GetRunNamespaceParamsWithContext creates a new RunsV1GetRunNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1GetRunNamespaceParamsWithContext(ctx context.Context) *RunsV1GetRunNamespaceParams {
	var ()
	return &RunsV1GetRunNamespaceParams{

		Context: ctx,
	}
}

// NewRunsV1GetRunNamespaceParamsWithHTTPClient creates a new RunsV1GetRunNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1GetRunNamespaceParamsWithHTTPClient(client *http.Client) *RunsV1GetRunNamespaceParams {
	var ()
	return &RunsV1GetRunNamespaceParams{
		HTTPClient: client,
	}
}

/*RunsV1GetRunNamespaceParams contains all the parameters to send to the API endpoint
for the runs v1 get run namespace operation typically these are written to a http.Request
*/
type RunsV1GetRunNamespaceParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project

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

// WithTimeout adds the timeout to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithTimeout(timeout time.Duration) *RunsV1GetRunNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithContext(ctx context.Context) *RunsV1GetRunNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithHTTPClient(client *http.Client) *RunsV1GetRunNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithOwner(owner string) *RunsV1GetRunNamespaceParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithProject(project string) *RunsV1GetRunNamespaceParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) WithUUID(uuid string) *RunsV1GetRunNamespaceParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the runs v1 get run namespace params
func (o *RunsV1GetRunNamespaceParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1GetRunNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
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