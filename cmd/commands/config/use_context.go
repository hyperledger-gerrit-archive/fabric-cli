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

// NewConfigUseContextCommand creates a new "fabric config use-context" command
func NewConfigUseContextCommand(settings *environment.Settings) *cobra.Command {
	c := UseContextCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "use-context <context-name>",
		Short: "change current context",
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

// UseContextCommand implements the config use context command
type UseContextCommand struct {
	command.BaseCommand

	Name string
}

// Validate checks the required parameters for run
func (c *UseContextCommand) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("context id not specified")
	}

	if _, ok := c.Settings.Config.Contexts[c.Name]; !ok {
		return fmt.Errorf("context '%s' does not exist", c.Name)
	}

	return nil
}

// Run executes the command
func (c *UseContextCommand) Run() error {
	err := c.Settings.ModifyConfig(environment.SetCurrentContext(c.Name))
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Settings.Streams.Out, "successfully set current context to '%s'\n", c.Name)

	return nil
}
