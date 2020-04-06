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

package project_dashboards_v1

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

// NewProjectDashboardsV1PromoteProjectDashboardParams creates a new ProjectDashboardsV1PromoteProjectDashboardParams object
// with the default values initialized.
func NewProjectDashboardsV1PromoteProjectDashboardParams() *ProjectDashboardsV1PromoteProjectDashboardParams {
	var ()
	return &ProjectDashboardsV1PromoteProjectDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectDashboardsV1PromoteProjectDashboardParamsWithTimeout creates a new ProjectDashboardsV1PromoteProjectDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectDashboardsV1PromoteProjectDashboardParamsWithTimeout(timeout time.Duration) *ProjectDashboardsV1PromoteProjectDashboardParams {
	var ()
	return &ProjectDashboardsV1PromoteProjectDashboardParams{

		timeout: timeout,
	}
}

// NewProjectDashboardsV1PromoteProjectDashboardParamsWithContext creates a new ProjectDashboardsV1PromoteProjectDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectDashboardsV1PromoteProjectDashboardParamsWithContext(ctx context.Context) *ProjectDashboardsV1PromoteProjectDashboardParams {
	var ()
	return &ProjectDashboardsV1PromoteProjectDashboardParams{

		Context: ctx,
	}
}

// NewProjectDashboardsV1PromoteProjectDashboardParamsWithHTTPClient creates a new ProjectDashboardsV1PromoteProjectDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectDashboardsV1PromoteProjectDashboardParamsWithHTTPClient(client *http.Client) *ProjectDashboardsV1PromoteProjectDashboardParams {
	var ()
	return &ProjectDashboardsV1PromoteProjectDashboardParams{
		HTTPClient: client,
	}
}

/*ProjectDashboardsV1PromoteProjectDashboardParams contains all the parameters to send to the API endpoint
for the project dashboards v1 promote project dashboard operation typically these are written to a http.Request
*/
type ProjectDashboardsV1PromoteProjectDashboardParams struct {

	/*DashboardUUID
	  UUID

	*/
	DashboardUUID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithTimeout(timeout time.Duration) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithContext(ctx context.Context) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithHTTPClient(client *http.Client) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardUUID adds the dashboardUUID to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithDashboardUUID(dashboardUUID string) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetDashboardUUID(dashboardUUID)
	return o
}

// SetDashboardUUID adds the dashboardUuid to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetDashboardUUID(dashboardUUID string) {
	o.DashboardUUID = dashboardUUID
}

// WithOwner adds the owner to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithOwner(owner string) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WithProject(project string) *ProjectDashboardsV1PromoteProjectDashboardParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the project dashboards v1 promote project dashboard params
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectDashboardsV1PromoteProjectDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard.uuid
	if err := r.SetPathParam("dashboard.uuid", o.DashboardUUID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}