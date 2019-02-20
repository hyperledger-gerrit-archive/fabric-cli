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
	"github.com/hyperledger/fabric-cli/pkg/environment/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProfileCreateCommand(t *testing.T) {
	cmd := NewProfileCreateCommand(testEnvironment())

	assert.NotNil(t, cmd)
	assert.False(t, cmd.HasSubCommands())
}

func TestCreateCommandComplete(t *testing.T) {
	pcmd := profileCreateCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{"foo"})

	assert.Nil(t, err)
	assert.Equal(t, pcmd.name, "foo")
}

func TestInstallCommandCompleteError(t *testing.T) {
	pcmd := profileCreateCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile name not specified")
}

func TestInstallCommandCompleteErrorTrim(t *testing.T) {
	pcmd := profileCreateCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{" "})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile name not specified")
}

func TestInstallCommandCompleteExists(t *testing.T) {
	pcmd := profileCreateCommand{
		out: new(bytes.Buffer),
		config: &environment.Settings{
			Profiles: []*environment.Profile{
				&environment.Profile{
					Name: "foo",
				},
			},
		},
	}

	err := pcmd.complete([]string{"foo"})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile 'foo' already exists")
}

func TestInstallCommandRun(t *testing.T) {
	mock := &mocks.MockConfig{}
	settings := &environment.Settings{Config: mock}

	pcmd := profileCreateCommand{
		out:    new(bytes.Buffer),
		config: settings,
		name:   "foobar",
	}

	err := pcmd.run()

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprint(pcmd.out), "successfully created profile 'foobar'\n")
}

func TestInstallCommandRunErr(t *testing.T) {
	mock := &mocks.MockConfig{}
	settings := &environment.Settings{Config: mock}

	mock.ExpectError("profile exists")

	pcmd := profileCreateCommand{
		out:    new(bytes.Buffer),
		config: settings,
	}

	err := pcmd.run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile exists")
}
