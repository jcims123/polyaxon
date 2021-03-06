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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadRunLogsReader is a Reader for the UploadRunLogs structure.
type UploadRunLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadRunLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadRunLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUploadRunLogsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUploadRunLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadRunLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadRunLogsOK creates a UploadRunLogsOK with default headers values
func NewUploadRunLogsOK() *UploadRunLogsOK {
	return &UploadRunLogsOK{}
}

/*UploadRunLogsOK handles this case with default header values.

A successful response.
*/
type UploadRunLogsOK struct {
}

func (o *UploadRunLogsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/logs/upload][%d] uploadRunLogsOK ", 200)
}

func (o *UploadRunLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadRunLogsNoContent creates a UploadRunLogsNoContent with default headers values
func NewUploadRunLogsNoContent() *UploadRunLogsNoContent {
	return &UploadRunLogsNoContent{}
}

/*UploadRunLogsNoContent handles this case with default header values.

No content.
*/
type UploadRunLogsNoContent struct {
	Payload interface{}
}

func (o *UploadRunLogsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/logs/upload][%d] uploadRunLogsNoContent  %+v", 204, o.Payload)
}

func (o *UploadRunLogsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadRunLogsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadRunLogsForbidden creates a UploadRunLogsForbidden with default headers values
func NewUploadRunLogsForbidden() *UploadRunLogsForbidden {
	return &UploadRunLogsForbidden{}
}

/*UploadRunLogsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UploadRunLogsForbidden struct {
	Payload interface{}
}

func (o *UploadRunLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/logs/upload][%d] uploadRunLogsForbidden  %+v", 403, o.Payload)
}

func (o *UploadRunLogsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadRunLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadRunLogsNotFound creates a UploadRunLogsNotFound with default headers values
func NewUploadRunLogsNotFound() *UploadRunLogsNotFound {
	return &UploadRunLogsNotFound{}
}

/*UploadRunLogsNotFound handles this case with default header values.

Resource does not exist.
*/
type UploadRunLogsNotFound struct {
	Payload interface{}
}

func (o *UploadRunLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/logs/upload][%d] uploadRunLogsNotFound  %+v", 404, o.Payload)
}

func (o *UploadRunLogsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadRunLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
