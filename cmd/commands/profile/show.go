/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package profile

import (
	"errors"
	"fmt"
	"io"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewProfileShowCommand creates a new "fabric profile show" command
func NewProfileShowCommand(settings *environment.Settings) *cobra.Command {
	pcmd := profileShowCommand{
		out:      settings.Streams.Out,
		profiles: settings.Profiles,
		active:   settings.ActiveProfile,
	}

	cmd := &cobra.Command{
		Use:   "show",
		Short: "show the metadata of the active profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return pcmd.run()
		},
	}

	return cmd
}

type profileShowCommand struct {
	out      io.Writer
	profiles []*environment.Profile
	active   string
}

func (cmd *profileShowCommand) run() error {
	if len(cmd.active) == 0 {
		return errors.New("profile not selected")
	}

	for _, p := range cmd.profiles {
		if p.Name == cmd.active {
			fmt.Fprintf(cmd.out, "Name: %s\n", p.Name)
		}
	}

	return nil
}
