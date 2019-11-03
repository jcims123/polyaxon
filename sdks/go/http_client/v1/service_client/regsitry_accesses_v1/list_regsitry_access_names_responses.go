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

package regsitry_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListRegsitryAccessNamesReader is a Reader for the ListRegsitryAccessNames structure.
type ListRegsitryAccessNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegsitryAccessNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegsitryAccessNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListRegsitryAccessNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRegsitryAccessNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRegsitryAccessNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRegsitryAccessNamesOK creates a ListRegsitryAccessNamesOK with default headers values
func NewListRegsitryAccessNamesOK() *ListRegsitryAccessNamesOK {
	return &ListRegsitryAccessNamesOK{}
}

/*ListRegsitryAccessNamesOK handles this case with default header values.

A successful response.
*/
type ListRegsitryAccessNamesOK struct {
	Payload *service_model.V1ListHostAccessesResponse
}

func (o *ListRegsitryAccessNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry_accesses/names][%d] listRegsitryAccessNamesOK  %+v", 200, o.Payload)
}

func (o *ListRegsitryAccessNamesOK) GetPayload() *service_model.V1ListHostAccessesResponse {
	return o.Payload
}

func (o *ListRegsitryAccessNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListHostAccessesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegsitryAccessNamesNoContent creates a ListRegsitryAccessNamesNoContent with default headers values
func NewListRegsitryAccessNamesNoContent() *ListRegsitryAccessNamesNoContent {
	return &ListRegsitryAccessNamesNoContent{}
}

/*ListRegsitryAccessNamesNoContent handles this case with default header values.

No content.
*/
type ListRegsitryAccessNamesNoContent struct {
	Payload interface{}
}

func (o *ListRegsitryAccessNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry_accesses/names][%d] listRegsitryAccessNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListRegsitryAccessNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegsitryAccessNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegsitryAccessNamesForbidden creates a ListRegsitryAccessNamesForbidden with default headers values
func NewListRegsitryAccessNamesForbidden() *ListRegsitryAccessNamesForbidden {
	return &ListRegsitryAccessNamesForbidden{}
}

/*ListRegsitryAccessNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListRegsitryAccessNamesForbidden struct {
	Payload interface{}
}

func (o *ListRegsitryAccessNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry_accesses/names][%d] listRegsitryAccessNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegsitryAccessNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegsitryAccessNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegsitryAccessNamesNotFound creates a ListRegsitryAccessNamesNotFound with default headers values
func NewListRegsitryAccessNamesNotFound() *ListRegsitryAccessNamesNotFound {
	return &ListRegsitryAccessNamesNotFound{}
}

/*ListRegsitryAccessNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListRegsitryAccessNamesNotFound struct {
	Payload interface{}
}

func (o *ListRegsitryAccessNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry_accesses/names][%d] listRegsitryAccessNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListRegsitryAccessNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegsitryAccessNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}