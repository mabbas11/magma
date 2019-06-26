/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package dictionary

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}
