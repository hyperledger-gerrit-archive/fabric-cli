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

// NewConfigSetContextCommand creates a new "fabric config set-context" command
func NewConfigSetContextCommand(settings *environment.Settings) *cobra.Command {
	c := SetContextCommand{
		Context: new(environment.Context),
	}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "set-context <context-name>",
		Short: "set a context",
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
	flags.StringVar(&c.Context.Network, "network", "", "set the network context")
	flags.StringVar(&c.Context.Organization, "organization", "", "set the organization context")
	flags.StringVar(&c.Context.Channel, "channel", "", "set the channel context")
	flags.StringVar(&c.Context.User, "user", "", "set the users context")

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// SetContextCommand implements the config set context command
type SetContextCommand struct {
	command.BaseCommand

	Name    string
	Context *environment.Context
}

// Validate checks the required parameters for run
func (c *SetContextCommand) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("context name not specified")
	}

	return nil
}

// Run executes the command
func (c *SetContextCommand) Run() error {
	err := c.Settings.ModifyConfig(environment.SetContext(c.Name, c.Context))
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Settings.Streams.Out, "successfully set context '%s'\n", c.Name)

	return nil
}
