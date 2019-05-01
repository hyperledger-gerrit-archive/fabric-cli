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

// NewConfigDeleteNetworkCommand creates a new "fabric config delete-network" command
func NewConfigDeleteNetworkCommand(settings *environment.Settings) *cobra.Command {
	c := DeleteNetworkCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "delete-network <network-name>",
		Short: "delete a network",
		Args:  c.ParseArgs(),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return c.Validate()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	c.AddArg(&c.Name)

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// DeleteNetworkCommand implements the config delete context command
type DeleteNetworkCommand struct {
	command.BaseCommand

	Name string
}

// Validate checks the required parameters for run
func (c *DeleteNetworkCommand) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("network name not specified")
	}

	return nil
}

// Run executes the command
func (c *DeleteNetworkCommand) Run() error {
	err := c.Settings.ModifyConfig(environment.DeleteNetwork(c.Name))
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Settings.Streams.Out, "successfully deleted network '%s'\n", c.Name)

	return nil
}
