/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"

	"github.com/hyperledger/fabric-cli/cmd/commands/command"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewConfigCurrentContextCommand creates a new "fabric config current-context" command
func NewConfigCurrentContextCommand(settings *environment.Settings) *cobra.Command {
	c := CurrentContextCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "current-context",
		Short: "view the current context",
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// CurrentContextCommand implements the config current context command
type CurrentContextCommand struct {
	command.BaseCommand
}

// Run executes the command
func (c *CurrentContextCommand) Run() error {
	context, err := c.Settings.Config.GetCurrentContext()
	if err != nil {
		return err
	}

	fmt.Fprintln(c.Settings.Streams.Out, "Name:", c.Settings.Config.CurrentContext)
	fmt.Fprintln(c.Settings.Streams.Out, "Network:", context.Network)
	fmt.Fprintln(c.Settings.Streams.Out, "Organization:", context.Organization)
	fmt.Fprintln(c.Settings.Streams.Out, "User:", context.User)
	fmt.Fprintln(c.Settings.Streams.Out, "Channel:", context.Channel)
	fmt.Fprintln(c.Settings.Streams.Out, "Orderers: ", context.Orderers)
	fmt.Fprintln(c.Settings.Streams.Out, "Peers: ", context.Peers)

	return nil
}
