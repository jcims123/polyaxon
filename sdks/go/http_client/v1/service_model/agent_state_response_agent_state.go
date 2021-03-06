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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AgentStateResponseAgentState agent state response agent state
//
// swagger:model AgentStateResponseAgentState
type AgentStateResponseAgentState struct {

	// List of apply runs
	Apply interface{} `json:"apply,omitempty"`

	// A flag to tell the agent that the queues are still full
	Full bool `json:"full,omitempty"`

	// List of notifier runs
	Notifier interface{} `json:"notifier,omitempty"`

	// List of queued runs
	Queued interface{} `json:"queued,omitempty"`

	// List of schdules runs
	Schedules interface{} `json:"schedules,omitempty"`

	// List of stopping runs
	Stopping interface{} `json:"stopping,omitempty"`

	// List of tuners runs
	Tuners interface{} `json:"tuners,omitempty"`

	// List of watchdogs runs
	Watchdogs interface{} `json:"watchdogs,omitempty"`
}

// Validate validates this agent state response agent state
func (m *AgentStateResponseAgentState) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgentStateResponseAgentState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentStateResponseAgentState) UnmarshalBinary(b []byte) error {
	var res AgentStateResponseAgentState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
