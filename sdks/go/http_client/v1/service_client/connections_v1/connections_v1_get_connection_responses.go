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

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ConnectionsV1GetConnectionReader is a Reader for the ConnectionsV1GetConnection structure.
type ConnectionsV1GetConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectionsV1GetConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectionsV1GetConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewConnectionsV1GetConnectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewConnectionsV1GetConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConnectionsV1GetConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConnectionsV1GetConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectionsV1GetConnectionOK creates a ConnectionsV1GetConnectionOK with default headers values
func NewConnectionsV1GetConnectionOK() *ConnectionsV1GetConnectionOK {
	return &ConnectionsV1GetConnectionOK{}
}

/*ConnectionsV1GetConnectionOK handles this case with default header values.

A successful response.
*/
type ConnectionsV1GetConnectionOK struct {
	Payload *service_model.V1ConnectionResponse
}

func (o *ConnectionsV1GetConnectionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] connectionsV1GetConnectionOK  %+v", 200, o.Payload)
}

func (o *ConnectionsV1GetConnectionOK) GetPayload() *service_model.V1ConnectionResponse {
	return o.Payload
}

func (o *ConnectionsV1GetConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1GetConnectionNoContent creates a ConnectionsV1GetConnectionNoContent with default headers values
func NewConnectionsV1GetConnectionNoContent() *ConnectionsV1GetConnectionNoContent {
	return &ConnectionsV1GetConnectionNoContent{}
}

/*ConnectionsV1GetConnectionNoContent handles this case with default header values.

No content.
*/
type ConnectionsV1GetConnectionNoContent struct {
	Payload interface{}
}

func (o *ConnectionsV1GetConnectionNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] connectionsV1GetConnectionNoContent  %+v", 204, o.Payload)
}

func (o *ConnectionsV1GetConnectionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1GetConnectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1GetConnectionForbidden creates a ConnectionsV1GetConnectionForbidden with default headers values
func NewConnectionsV1GetConnectionForbidden() *ConnectionsV1GetConnectionForbidden {
	return &ConnectionsV1GetConnectionForbidden{}
}

/*ConnectionsV1GetConnectionForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ConnectionsV1GetConnectionForbidden struct {
	Payload interface{}
}

func (o *ConnectionsV1GetConnectionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] connectionsV1GetConnectionForbidden  %+v", 403, o.Payload)
}

func (o *ConnectionsV1GetConnectionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1GetConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1GetConnectionNotFound creates a ConnectionsV1GetConnectionNotFound with default headers values
func NewConnectionsV1GetConnectionNotFound() *ConnectionsV1GetConnectionNotFound {
	return &ConnectionsV1GetConnectionNotFound{}
}

/*ConnectionsV1GetConnectionNotFound handles this case with default header values.

Resource does not exist.
*/
type ConnectionsV1GetConnectionNotFound struct {
	Payload interface{}
}

func (o *ConnectionsV1GetConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] connectionsV1GetConnectionNotFound  %+v", 404, o.Payload)
}

func (o *ConnectionsV1GetConnectionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1GetConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1GetConnectionDefault creates a ConnectionsV1GetConnectionDefault with default headers values
func NewConnectionsV1GetConnectionDefault(code int) *ConnectionsV1GetConnectionDefault {
	return &ConnectionsV1GetConnectionDefault{
		_statusCode: code,
	}
}

/*ConnectionsV1GetConnectionDefault handles this case with default header values.

An unexpected error response
*/
type ConnectionsV1GetConnectionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the connections v1 get connection default response
func (o *ConnectionsV1GetConnectionDefault) Code() int {
	return o._statusCode
}

func (o *ConnectionsV1GetConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] ConnectionsV1_GetConnection default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectionsV1GetConnectionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ConnectionsV1GetConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}