/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package context

import (
	"fmt"

	"github.com/hyperledger/fabric-cli/cmd/commands/command"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewContextCurrentCommand creates a new "fabric context current" command
func NewContextCurrentCommand(settings *environment.Settings) *cobra.Command {
	c := CurrentCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "current",
		Short: "view the current context",
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// CurrentCommand implements the current command
type CurrentCommand struct {
	command.BaseCommand
}

// Run executes the command
func (c *CurrentCommand) Run() error {
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
