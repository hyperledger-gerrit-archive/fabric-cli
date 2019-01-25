/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/hyperledger/fabric-cli/pkg/plugin"
	"github.com/spf13/cobra"
)

// LoadPlugins processes all of the installed plugins, wraps them with cobra,
// and adds them to the root command
func LoadPlugins(cmd *cobra.Command, settings *environment.Settings, handler plugin.Handler) {
	if settings.DisablePlugins {
		return
	}

	plugins, err := handler.GetPlugins()
	if err != nil {
		fmt.Fprintf(settings.Streams.Err, "An error occurred while loading plugins: %s", err)
		return
	}

	for _, plugin := range plugins {
		p := plugin
		c := &cobra.Command{
			Use:   p.Name,
			Short: p.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				e := exec.Command(os.ExpandEnv(p.Command.Base),
					append(p.Command.Args, args...)...)
				e.Stdin = settings.Streams.In
				e.Stdout = settings.Streams.Out
				e.Stderr = settings.Streams.Err
				return e.Run()
			},
		}

		cmd.AddCommand(c)
	}
}
