/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package profile

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewProfileRemoveCommand creates a new "fabric profile remove" command
func NewProfileRemoveCommand(settings *environment.Settings) *cobra.Command {
	pcmd := profileRemoveCommand{
		out:    settings.Streams.Out,
		active: settings.ActiveProfile,
	}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a configuration profile",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := settings.FromFile()
			if err != nil {
				return err
			}

			pcmd.config = config

			return pcmd.complete(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return pcmd.run()
		},
	}

	return cmd
}

type profileRemoveCommand struct {
	out    io.Writer
	config *environment.Settings
	active string

	name string
}

func (cmd *profileRemoveCommand) complete(args []string) error {
	if len(args) == 0 {
		return errors.New("profile name not specified")
	}

	cmd.name = strings.TrimSpace(args[0])

	if len(cmd.name) == 0 {
		return errors.New("profile name not specified")
	}

	if cmd.name == cmd.active {
		return errors.New("cannot remove active profile")
	}

	return nil
}

func (cmd *profileRemoveCommand) run() error {
	var found bool
	for i, p := range cmd.config.Profiles {
		if p.Name == cmd.name {
			found = true
			cmd.config.Profiles = append(cmd.config.Profiles[:i], cmd.config.Profiles[i+1:]...)
		}
	}

	if !found {
		return fmt.Errorf("profile '%s' was not found", cmd.name)
	}

	err := cmd.config.Save()
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.out, "successfully removed profile '%s'\n", cmd.name)

	return nil
}
