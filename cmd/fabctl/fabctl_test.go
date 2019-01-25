/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"testing"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	settings := &environment.Settings{
		Home: environment.DefaultHome,
		Streams: environment.Streams{
			In:  new(bytes.Buffer),
			Out: new(bytes.Buffer),
			Err: new(bytes.Buffer),
		},
	}
	cmd := NewFabctl(settings)

	assert.NotNil(t, cmd)
	assert.True(t, cmd.HasSubCommands())

	err := cmd.Execute()

	assert.Nil(t, err)
}
