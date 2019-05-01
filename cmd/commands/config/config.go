/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewConfigCommand creates a new "fabric config" command
func NewConfigCommand(settings *environment.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configurations",
	}

	cmd.AddCommand(
		NewConfigCurrentContextCommand(settings),
		NewConfigListContextCommand(settings),
		NewConfigUseContextCommand(settings),
		NewConfigSetContextCommand(settings),
		NewConfigDeleteContextCommand(settings),
		NewConfigListNetworkCommand(settings),
		NewConfigSetNetworkCommand(settings),
		NewConfigDeleteNetworkCommand(settings),
	)

	cmd.SetOutput(settings.Streams.Out)

	return cmd
}
