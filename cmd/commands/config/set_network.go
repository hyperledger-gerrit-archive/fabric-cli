/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-cli/cmd/commands/command"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewConfigSetNetworkCommand creates a new "fabric config set-network" command
func NewConfigSetNetworkCommand(settings *environment.Settings) *cobra.Command {
	c := SetNetworkCommand{
		Network: new(environment.Network),
	}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "set-network <network-name>",
		Short: "set a network",
		Args:  c.ParseArgs(),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return c.Validate()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	c.AddArg(&c.Name)

	flags := cmd.Flags()
	flags.StringVar(&c.Network.ConfigPath, "path", "", "set the path to network configurations")

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// SetNetworkCommand implements the config set network command
type SetNetworkCommand struct {
	command.BaseCommand

	Name    string
	Network *environment.Network
}

// Validate checks the required parameters for run
func (c *SetNetworkCommand) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("network name not specified")
	}

	return nil
}

// Run executes the command
func (c *SetNetworkCommand) Run() error {
	err := c.Settings.ModifyConfig(environment.SetNetwork(c.Name, c.Network))
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Settings.Streams.Out, "successfully set network '%s'\n", c.Name)

	return nil
}
