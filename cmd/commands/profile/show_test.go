/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package profile

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/stretchr/testify/assert"
)

func TestProfileShowCommand(t *testing.T) {
	cmd := NewProfileShowCommand(testEnvironment())

	assert.NotNil(t, cmd)
	assert.False(t, cmd.HasSubCommands())
}

func TestShowCommandRun(t *testing.T) {
	pcmd := profileShowCommand{
		out: new(bytes.Buffer),
		profiles: []*environment.Profile{
			&environment.Profile{
				Name: "foobar",
			},
		},
		active: "foobar",
	}

	err := pcmd.run()

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprint(pcmd.out), "Name: foobar\n")
}

func TestShowCommandRunError(t *testing.T) {
	pcmd := profileShowCommand{
		out: new(bytes.Buffer),
	}

	err := pcmd.run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile not selected")
}
