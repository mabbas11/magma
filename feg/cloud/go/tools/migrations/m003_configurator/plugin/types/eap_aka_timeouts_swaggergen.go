/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EapAkaTimeouts eap aka timeouts
// swagger:model eap_aka_timeouts
type EapAkaTimeouts struct {

	// challenge ms
	ChallengeMs uint32 `json:"challenge_ms,omitempty"`

	// error notification ms
	ErrorNotificationMs uint32 `json:"error_notification_ms,omitempty"`

	// session authenticated ms
	SessionAuthenticatedMs uint32 `json:"session_authenticated_ms,omitempty"`

	// session ms
	SessionMs uint32 `json:"session_ms,omitempty"`
}

// Validate validates this eap aka timeouts
func (m *EapAkaTimeouts) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EapAkaTimeouts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EapAkaTimeouts) UnmarshalBinary(b []byte) error {
	var res EapAkaTimeouts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
