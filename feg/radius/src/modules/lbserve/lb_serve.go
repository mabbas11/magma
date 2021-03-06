/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package lbserve

import (
	"context"
	"errors"

	"fbc/cwf/radius/modules"
	"fbc/lib/go/radius"

	"go.uber.org/zap"
)

const errMissingRequiredUpstreamHostText = "session state required field upstream host is missing, unable to serve request"

var errMissingRequiredUpstreamHost = errors.New(errMissingRequiredUpstreamHostText)

// Init module interface implementation
func Init(logger *zap.Logger, config modules.ModuleConfig) error {
	return nil
}

// Handle module interface implementation
func Handle(c *modules.RequestContext, r *radius.Request, _ modules.Middleware) (*modules.Response, error) {
	state, err := c.SessionStorage.Get()
	if err != nil {
		c.Logger.Error(
			"Error loading session state, unable to serve request",
			zap.Error(err),
		)
		return nil, err
	}

	if state.UpstreamHost == "" {
		c.Logger.Error(errMissingRequiredUpstreamHostText)
		return nil, errMissingRequiredUpstreamHost
	}

	res, err := radius.Exchange(context.Background(), r.Packet, state.UpstreamHost)
	if err != nil {
		c.Logger.Error("LB Serve received failed response", zap.Error(err))
		return nil, err
	}

	return &modules.Response{
		Code:       res.Code,
		Attributes: res.Attributes,
	}, nil
}
