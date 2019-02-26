/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/hyperledger/fabric-cli/pkg/plugin"
	"github.com/hyperledger/fabric-cli/pkg/plugin/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func TestFabricCommand(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Fabric Suite")
}

var _ = Describe("FabricCommand", func() {
	var (
		cmd      *cobra.Command
		settings *environment.Settings
		out      *bytes.Buffer
	)

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
		cmd = NewFabricCommand(settings)
	})

	It("should create a fabric commmand", func() {
		Expect(cmd.Name()).To(Equal("fabric"))
		Expect(cmd.HasSubCommands()).To(BeTrue())
		Expect(cmd.Execute()).Should(Succeed())
		Expect(fmt.Sprint(out)).To(ContainSubstring("fabric [command]"))
		Expect(fmt.Sprint(out)).To(ContainSubstring("profile"))
		Expect(fmt.Sprint(out)).To(ContainSubstring("plugin"))
	})

})

var _ = Describe("LoadPlugins", func() {
	var (
		out, err *bytes.Buffer
		cmd      *cobra.Command
		settings *environment.Settings
		handler  *mocks.PluginHandler
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)
		err = new(bytes.Buffer)
		cmd = &cobra.Command{}

		cmd.SetOutput(out)

		settings = &environment.Settings{
			Home: environment.Home(os.TempDir()),
			Streams: environment.Streams{
				Err: err,
			},
		}

		handler = &mocks.PluginHandler{}
	})

	JustBeforeEach(func() {
		loadPlugins(cmd, settings, handler)
	})

	Context("when plugins are disabled", func() {
		BeforeEach(func() {
			settings.DisablePlugins = true

			err := handler.InstallPlugin("foo")

			Expect(err).NotTo(HaveOccurred())
		})

		It("should not have foo command", func() {
			Expect(cmd.Execute()).Should(Succeed())
			Expect(cmd.HasSubCommands()).Should(BeFalse())
		})
	})

	Context("when plugin handler fails", func() {
		BeforeEach(func() {
			handler.GetPluginsReturns(nil, errors.New("handler error"))
		})

		It("should fail to load plugins", func() {
			Expect(fmt.Sprint(err)).To(ContainSubstring("An error occurred while loading plugins: handler error"))
		})
	})

	Context("when plugins have been installed", func() {
		BeforeEach(func() {
			handler.GetPluginsReturns([]*plugin.Plugin{
				&plugin.Plugin{
					Name: "foo",
				},
			}, nil)
		})

		It("should load plugins", func() {
			Expect(cmd.Execute()).Should(Succeed())
			Expect(cmd.HasSubCommands()).Should(BeTrue())
		})
	})
})
