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

func TestProfileRemoveCommand(t *testing.T) {
	cmd := NewProfileRemoveCommand(testEnvironment())

	assert.NotNil(t, cmd)
	assert.False(t, cmd.HasSubCommands())
}

func TestRemoveCommandComplete(t *testing.T) {
	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{"foobar"})

	assert.Nil(t, err)
	assert.Equal(t, pcmd.name, "foobar")
}

func TestRemoveCommandCompleteError(t *testing.T) {
	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile name not specified")
}

func TestRemoveCommandCompleteErrorTrim(t *testing.T) {
	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: &environment.Settings{},
	}

	err := pcmd.complete([]string{" "})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile name not specified")
}

func TestRemoveCommandCompleteActive(t *testing.T) {
	pcmd := profileRemoveCommand{
		out: new(bytes.Buffer),
		config: &environment.Settings{
			Profiles: []*environment.Profile{
				&environment.Profile{
					Name: "foobar",
				},
			},
		},
		active: "foobar",
	}

	err := pcmd.complete([]string{"foobar"})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "cannot remove active profile")
}

func TestRemoveCommandRun(t *testing.T) {
	mock := &mocks.MockConfig{}
	settings := &environment.Settings{
		Config: mock,
		Profiles: []*environment.Profile{
			&environment.Profile{
				Name: "foobar",
			},
		},
	}

	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: settings,
		name:   "foobar",
	}

	err := pcmd.run()

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprint(pcmd.out), "successfully removed profile 'foobar'\n")
}

func TestRemoveCommandRunError(t *testing.T) {
	mock := &mocks.MockConfig{}
	settings := &environment.Settings{
		Config: mock,
		Profiles: []*environment.Profile{
			&environment.Profile{
				Name: "foobar",
			},
		},
	}

	mock.ExpectError("save error")

	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: settings,
		name:   "foobar",
	}

	err := pcmd.run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "save error")
}

func TestRemoveCommandErrorDNE(t *testing.T) {
	mock := &mocks.MockConfig{}
	settings := &environment.Settings{
		Config: mock,
	}

	pcmd := profileRemoveCommand{
		out:    new(bytes.Buffer),
		config: settings,
		name:   "foobar",
	}

	err := pcmd.run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "profile 'foobar' was not found")
}
