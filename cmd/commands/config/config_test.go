/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/hyperledger/fabric-cli/cmd/commands/config"
	"github.com/hyperledger/fabric-cli/pkg/environment"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("ConfigCommand", func() {
	var (
		cmd      *cobra.Command
		settings *environment.Settings
		out      *bytes.Buffer
	)

	Context("when creating a command from settings", func() {
		BeforeEach(func() {
			out = new(bytes.Buffer)

			settings = &environment.Settings{
				Home: environment.Home(os.TempDir()),
				Streams: environment.Streams{
					Out: out,
				},
			}
		})

		JustBeforeEach(func() {
			cmd = config.NewConfigCommand(settings)
		})

		It("should create a config commmand", func() {
			Expect(cmd.Name()).To(Equal("config"))
			Expect(cmd.HasSubCommands()).To(BeTrue())
			Expect(cmd.Execute()).Should(Succeed())
			Expect(fmt.Sprint(out)).To(ContainSubstring("config [command]"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("current-context"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("list-contexts"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("set-context"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("delete-context"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("use-context"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("list-networks"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("set-network"))
			Expect(fmt.Sprint(out)).To(ContainSubstring("delete-network"))
		})
	})
})
