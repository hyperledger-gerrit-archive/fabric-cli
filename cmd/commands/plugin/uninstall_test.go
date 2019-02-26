/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package plugin_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/hyperledger/fabric-cli/cmd/commands/plugin"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/hyperledger/fabric-cli/pkg/plugin/mocks"
)

var _ = Describe("PluginUninstallCommand", func() {
	var (
		cmd      *cobra.Command
		settings *environment.Settings
		out      *bytes.Buffer

		args []string
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)

		settings = &environment.Settings{
			Home: environment.Home(os.TempDir()),
			Streams: environment.Streams{
				Out: out,
			},
		}

		args = os.Args
	})

	JustBeforeEach(func() {
		cmd = plugin.NewPluginUninstallCommand(settings)
	})

	AfterEach(func() {
		os.Args = args
	})

	It("should create a plugin uninstall commmand", func() {
		Expect(cmd.Name()).To(Equal("uninstall"))
		Expect(cmd.HasSubCommands()).To(BeFalse())
	})

	It("should provide a help prompt", func() {
		os.Args = append(os.Args, "--help")

		Expect(cmd.Execute()).Should(Succeed())
		Expect(fmt.Sprint(out)).To(ContainSubstring("uninstall <name>"))
	})
})

var _ = Describe("PluginUninstallImplementation", func() {
	var (
		impl    *plugin.UninstallCommand
		out     *bytes.Buffer
		handler *mocks.PluginHandler
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)
		handler = &mocks.PluginHandler{}
	})

	JustBeforeEach(func() {
		impl = &plugin.UninstallCommand{
			Out:     out,
			Handler: handler,
		}
	})

	It("should not be nil", func() {
		Expect(impl).ShouldNot(BeNil())
	})

	Describe("Complete", func() {
		It("should fail without args", func() {
			err := impl.Complete([]string{})

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("plugin name not specified"))
		})

		It("should fail with empty string", func() {
			err := impl.Complete([]string{" "})

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("plugin name not specified"))
		})

		It("should succeed with a name", func() {
			Expect(impl.Complete([]string{"foo"})).Should(Succeed())
		})
	})

	Describe("Run", func() {
		BeforeEach(func() {
			err := impl.Complete([]string{"foo"})

			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail if handler fails", func() {
			handler.UninstallPluginReturns(errors.New("handler error"))

			err := impl.Run()

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("handler error"))
		})

		Context("when plugins have been installed", func() {
			JustBeforeEach(func() {
				err := handler.InstallPlugin("foo")

				Expect(err).NotTo(HaveOccurred())
			})

			It("should successfully uninstall a plugin", func() {
				Expect(impl.Run()).Should(Succeed())
				Expect(fmt.Sprint(out)).To(ContainSubstring("successfully uninstalled the plugin\n"))
			})
		})
	})
})
