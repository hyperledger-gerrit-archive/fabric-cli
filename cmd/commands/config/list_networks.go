/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"errors"
	"fmt"
	"sort"

	"github.com/hyperledger/fabric-cli/cmd/commands/command"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewConfigListNetworkCommand creates a new "fabric config list-network" command
func NewConfigListNetworkCommand(settings *environment.Settings) *cobra.Command {
	c := ListNetworkCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "list-networks",
		Short: "list all networks",
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// ListNetworkCommand implements the config delete context command
type ListNetworkCommand struct {
	command.BaseCommand
}

// Run executes the command
func (c *ListNetworkCommand) Run() error {
	if len(c.Settings.Config.Networks) == 0 {
		return errors.New("no networks currently exist")
	}

	context, _ := c.Settings.Config.GetCurrentContext()

	var names []string
	for name := range c.Settings.Config.Networks {
		if context != nil && name == context.Network {
			name += " (current)"
		}
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Fprintln(c.Settings.Streams.Out, name)
	}

	return nil
}
