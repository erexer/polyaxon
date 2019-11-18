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

package k8s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateK8sSecretReader is a Reader for the UpdateK8sSecret structure.
type UpdateK8sSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateK8sSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateK8sSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateK8sSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateK8sSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateK8sSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateK8sSecretOK creates a UpdateK8sSecretOK with default headers values
func NewUpdateK8sSecretOK() *UpdateK8sSecretOK {
	return &UpdateK8sSecretOK{}
}

/*UpdateK8sSecretOK handles this case with default header values.

A successful response.
*/
type UpdateK8sSecretOK struct {
	Payload *service_model.V1K8sResource
}

func (o *UpdateK8sSecretOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_secrets/{k8s_resource.uuid}][%d] updateK8sSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateK8sSecretOK) GetPayload() *service_model.V1K8sResource {
	return o.Payload
}

func (o *UpdateK8sSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8sResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8sSecretNoContent creates a UpdateK8sSecretNoContent with default headers values
func NewUpdateK8sSecretNoContent() *UpdateK8sSecretNoContent {
	return &UpdateK8sSecretNoContent{}
}

/*UpdateK8sSecretNoContent handles this case with default header values.

No content.
*/
type UpdateK8sSecretNoContent struct {
	Payload interface{}
}

func (o *UpdateK8sSecretNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_secrets/{k8s_resource.uuid}][%d] updateK8sSecretNoContent  %+v", 204, o.Payload)
}

func (o *UpdateK8sSecretNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8sSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8sSecretForbidden creates a UpdateK8sSecretForbidden with default headers values
func NewUpdateK8sSecretForbidden() *UpdateK8sSecretForbidden {
	return &UpdateK8sSecretForbidden{}
}

/*UpdateK8sSecretForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateK8sSecretForbidden struct {
	Payload interface{}
}

func (o *UpdateK8sSecretForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_secrets/{k8s_resource.uuid}][%d] updateK8sSecretForbidden  %+v", 403, o.Payload)
}

func (o *UpdateK8sSecretForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8sSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8sSecretNotFound creates a UpdateK8sSecretNotFound with default headers values
func NewUpdateK8sSecretNotFound() *UpdateK8sSecretNotFound {
	return &UpdateK8sSecretNotFound{}
}

/*UpdateK8sSecretNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateK8sSecretNotFound struct {
	Payload interface{}
}

func (o *UpdateK8sSecretNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_secrets/{k8s_resource.uuid}][%d] updateK8sSecretNotFound  %+v", 404, o.Payload)
}

func (o *UpdateK8sSecretNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8sSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}