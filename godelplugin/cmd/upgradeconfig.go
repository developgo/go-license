// Copyright (c) 2016 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi"

	"github.com/palantir/go-license/golicense/config"
)

var upgradeConfigCmd = pluginapi.CobraUpgradeConfigCmd(config.UpgradeConfig)

func init() {
	RootCmd.AddCommand(upgradeConfigCmd)
}
