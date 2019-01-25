/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hyperledger/fabric-cli/cmd/fabctl/commands"
	"github.com/hyperledger/fabric-cli/cmd/fabctl/util"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/hyperledger/fabric-cli/pkg/plugin"
)

// NewFabctl returns a new root command for fabctl
func NewFabctl(settings *environment.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fabctl",
		Short: "The command line interface for Hyperledger Fabric",
	}

	// load all built in commands into the root command
	cmd.AddCommand(commands.All(settings)...)

	// load all plugins into the root command
	util.LoadPlugins(cmd, settings, &plugin.DefaultHandler{
		Dir:              settings.Home.Plugins(),
		MetadataFileName: plugin.DefaultMetadataFileName,
	})

	return cmd
}

func main() {
	settings, err := environment.GetSettings()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	cmd := NewFabctl(settings)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
