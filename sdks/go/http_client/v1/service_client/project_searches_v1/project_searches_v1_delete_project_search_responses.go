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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectSearchesV1DeleteProjectSearchReader is a Reader for the ProjectSearchesV1DeleteProjectSearch structure.
type ProjectSearchesV1DeleteProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectSearchesV1DeleteProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectSearchesV1DeleteProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectSearchesV1DeleteProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectSearchesV1DeleteProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectSearchesV1DeleteProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectSearchesV1DeleteProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectSearchesV1DeleteProjectSearchOK creates a ProjectSearchesV1DeleteProjectSearchOK with default headers values
func NewProjectSearchesV1DeleteProjectSearchOK() *ProjectSearchesV1DeleteProjectSearchOK {
	return &ProjectSearchesV1DeleteProjectSearchOK{}
}

/*ProjectSearchesV1DeleteProjectSearchOK handles this case with default header values.

A successful response.
*/
type ProjectSearchesV1DeleteProjectSearchOK struct {
}

func (o *ProjectSearchesV1DeleteProjectSearchOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] projectSearchesV1DeleteProjectSearchOK ", 200)
}

func (o *ProjectSearchesV1DeleteProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectSearchesV1DeleteProjectSearchNoContent creates a ProjectSearchesV1DeleteProjectSearchNoContent with default headers values
func NewProjectSearchesV1DeleteProjectSearchNoContent() *ProjectSearchesV1DeleteProjectSearchNoContent {
	return &ProjectSearchesV1DeleteProjectSearchNoContent{}
}

/*ProjectSearchesV1DeleteProjectSearchNoContent handles this case with default header values.

No content.
*/
type ProjectSearchesV1DeleteProjectSearchNoContent struct {
	Payload interface{}
}

func (o *ProjectSearchesV1DeleteProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] projectSearchesV1DeleteProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *ProjectSearchesV1DeleteProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1DeleteProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1DeleteProjectSearchForbidden creates a ProjectSearchesV1DeleteProjectSearchForbidden with default headers values
func NewProjectSearchesV1DeleteProjectSearchForbidden() *ProjectSearchesV1DeleteProjectSearchForbidden {
	return &ProjectSearchesV1DeleteProjectSearchForbidden{}
}

/*ProjectSearchesV1DeleteProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectSearchesV1DeleteProjectSearchForbidden struct {
	Payload interface{}
}

func (o *ProjectSearchesV1DeleteProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] projectSearchesV1DeleteProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *ProjectSearchesV1DeleteProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1DeleteProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1DeleteProjectSearchNotFound creates a ProjectSearchesV1DeleteProjectSearchNotFound with default headers values
func NewProjectSearchesV1DeleteProjectSearchNotFound() *ProjectSearchesV1DeleteProjectSearchNotFound {
	return &ProjectSearchesV1DeleteProjectSearchNotFound{}
}

/*ProjectSearchesV1DeleteProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectSearchesV1DeleteProjectSearchNotFound struct {
	Payload interface{}
}

func (o *ProjectSearchesV1DeleteProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] projectSearchesV1DeleteProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *ProjectSearchesV1DeleteProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1DeleteProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1DeleteProjectSearchDefault creates a ProjectSearchesV1DeleteProjectSearchDefault with default headers values
func NewProjectSearchesV1DeleteProjectSearchDefault(code int) *ProjectSearchesV1DeleteProjectSearchDefault {
	return &ProjectSearchesV1DeleteProjectSearchDefault{
		_statusCode: code,
	}
}

/*ProjectSearchesV1DeleteProjectSearchDefault handles this case with default header values.

An unexpected error response
*/
type ProjectSearchesV1DeleteProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project searches v1 delete project search default response
func (o *ProjectSearchesV1DeleteProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *ProjectSearchesV1DeleteProjectSearchDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] ProjectSearchesV1_DeleteProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectSearchesV1DeleteProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectSearchesV1DeleteProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}