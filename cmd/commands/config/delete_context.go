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

// NewConfigDeleteContextCommand creates a new "fabric config delete-context" command
func NewConfigDeleteContextCommand(settings *environment.Settings) *cobra.Command {
	c := DeleteContextCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "delete-context <context-name>",
		Short: "delete a context",
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

// DeleteContextCommand implements the config delete context command
type DeleteContextCommand struct {
	command.BaseCommand

	Name string
}

// Validate checks the required parameters for run
func (c *DeleteContextCommand) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("context name not specified")
	}

	return nil
}

// Run executes the command
func (c *DeleteContextCommand) Run() error {
	err := c.Settings.ModifyConfig(environment.DeleteContext(c.Name))
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Settings.Streams.Out, "successfully deleted context '%s'\n", c.Name)

	return nil
}
