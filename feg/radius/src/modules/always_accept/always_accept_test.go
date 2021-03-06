/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package alwaysaccept

import (
	"fbc/cwf/radius/modules"
	"fbc/lib/go/radius"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/require"
)

func TestAccessRequest(t *testing.T) {
	// Arrange
	logger, err := zap.NewDevelopment()
	require.NoError(t, err, "failed to get logger")
	Init(logger, modules.ModuleConfig{})

	// Act
	res, err := Handle(
		&modules.RequestContext{
			RequestID:      0,
			Logger:         logger,
			SessionStorage: nil,
		},
		&radius.Request{Packet: &radius.Packet{
			Code: radius.CodeAccessRequest,
		}},
		func(c *modules.RequestContext, r *radius.Request) (*modules.Response, error) {
			require.Fail(t, "next method is called bu not expected to")
			return nil, nil
		},
	)

	// Act and Assert
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, radius.CodeAccessAccept, res.Code)
}

func TestNotAccessRequest(t *testing.T) {
	// Arrange
	logger, err := zap.NewDevelopment()
	require.NoError(t, err, "failed to get logger")
	Init(logger, modules.ModuleConfig{})

	// Act
	res, err := Handle(
		&modules.RequestContext{
			RequestID:      0,
			Logger:         logger,
			SessionStorage: nil,
		},
		&radius.Request{Packet: &radius.Packet{
			Code: radius.CodeAccountingRequest,
		}},
		func(c *modules.RequestContext, r *radius.Request) (*modules.Response, error) {
			require.Fail(t, "next method is called bu not expected to")
			return nil, nil
		},
	)

	// Act and Assert
	require.NotNil(t, err)
	require.Nil(t, res)
}
