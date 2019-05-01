/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package command

import (
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// BaseCommand implements common command functions
type BaseCommand struct {
	Settings *environment.Settings
	Args     []*string
}

// AddArg makes the command aware that it is expecting an argument
// The order that args are added is important
func (c *BaseCommand) AddArg(arg *string) {
	if arg == nil {
		return
	}

	c.Args = append(c.Args, arg)
}

// ParseArgs writes the arg values to the specified locations.
func (c *BaseCommand) ParseArgs() cobra.PositionalArgs {
	return func(_ *cobra.Command, args []string) error {
		for i, arg := range c.Args {
			if len(args) < i+1 {
				return nil
			}

			*arg = args[i]
		}

		return nil
	}
}
