// Copyright 2019 Polyaxon, Inc.
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

package config_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchConfigResourceParams creates a new PatchConfigResourceParams object
// with the default values initialized.
func NewPatchConfigResourceParams() *PatchConfigResourceParams {
	var ()
	return &PatchConfigResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConfigResourceParamsWithTimeout creates a new PatchConfigResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConfigResourceParamsWithTimeout(timeout time.Duration) *PatchConfigResourceParams {
	var ()
	return &PatchConfigResourceParams{

		timeout: timeout,
	}
}

// NewPatchConfigResourceParamsWithContext creates a new PatchConfigResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConfigResourceParamsWithContext(ctx context.Context) *PatchConfigResourceParams {
	var ()
	return &PatchConfigResourceParams{

		Context: ctx,
	}
}

// NewPatchConfigResourceParamsWithHTTPClient creates a new PatchConfigResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConfigResourceParamsWithHTTPClient(client *http.Client) *PatchConfigResourceParams {
	var ()
	return &PatchConfigResourceParams{
		HTTPClient: client,
	}
}

/*PatchConfigResourceParams contains all the parameters to send to the API endpoint
for the patch config resource operation typically these are written to a http.Request
*/
type PatchConfigResourceParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1ConfigResource
	/*ConfigResourceUUID
	  UUID

	*/
	ConfigResourceUUID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch config resource params
func (o *PatchConfigResourceParams) WithTimeout(timeout time.Duration) *PatchConfigResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch config resource params
func (o *PatchConfigResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch config resource params
func (o *PatchConfigResourceParams) WithContext(ctx context.Context) *PatchConfigResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch config resource params
func (o *PatchConfigResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch config resource params
func (o *PatchConfigResourceParams) WithHTTPClient(client *http.Client) *PatchConfigResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch config resource params
func (o *PatchConfigResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch config resource params
func (o *PatchConfigResourceParams) WithBody(body *service_model.V1ConfigResource) *PatchConfigResourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch config resource params
func (o *PatchConfigResourceParams) SetBody(body *service_model.V1ConfigResource) {
	o.Body = body
}

// WithConfigResourceUUID adds the configResourceUUID to the patch config resource params
func (o *PatchConfigResourceParams) WithConfigResourceUUID(configResourceUUID string) *PatchConfigResourceParams {
	o.SetConfigResourceUUID(configResourceUUID)
	return o
}

// SetConfigResourceUUID adds the configResourceUuid to the patch config resource params
func (o *PatchConfigResourceParams) SetConfigResourceUUID(configResourceUUID string) {
	o.ConfigResourceUUID = configResourceUUID
}

// WithOwner adds the owner to the patch config resource params
func (o *PatchConfigResourceParams) WithOwner(owner string) *PatchConfigResourceParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch config resource params
func (o *PatchConfigResourceParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConfigResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param config_resource.uuid
	if err := r.SetPathParam("config_resource.uuid", o.ConfigResourceUUID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}