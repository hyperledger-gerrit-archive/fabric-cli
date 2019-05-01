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

// NewConfigListContextCommand creates a new "fabric config list-context" command
func NewConfigListContextCommand(settings *environment.Settings) *cobra.Command {
	c := ListContextCommand{}

	c.Settings = settings

	cmd := &cobra.Command{
		Use:   "list-contexts",
		Short: "list all contexts",
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	cmd.SetOutput(c.Settings.Streams.Out)

	return cmd
}

// ListContextCommand implements the config delete context command
type ListContextCommand struct {
	command.BaseCommand
}

// Run executes the command
func (c *ListContextCommand) Run() error {
	if len(c.Settings.Config.Contexts) == 0 {
		return errors.New("no contexts currently exist")
	}

	var names []string
	for name := range c.Settings.Config.Contexts {
		if name == c.Settings.Config.CurrentContext {
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
