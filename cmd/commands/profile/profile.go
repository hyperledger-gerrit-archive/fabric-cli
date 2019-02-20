/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package profile

import (
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/spf13/cobra"
)

// NewProfileCommand creates a new "fabric profile" command
func NewProfileCommand(settings *environment.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Manage profiles",
	}

	cmd.AddCommand(
		NewProfileUseCommand(settings),
		NewProfileShowCommand(settings),
		NewProfileCreateCommand(settings),
		NewProfileRemoveCommand(settings),
	)

	return cmd
}
