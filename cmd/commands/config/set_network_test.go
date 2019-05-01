/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/hyperledger/fabric-cli/cmd/commands/config"
	"github.com/hyperledger/fabric-cli/pkg/environment"
)

var _ = Describe("SetNetworkCommand", func() {
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
		cmd = config.NewConfigSetNetworkCommand(settings)
	})

	AfterEach(func() {
		os.Args = args
	})

	It("should create a set network commmand", func() {
		Expect(cmd.Name()).To(Equal("set-network"))
		Expect(cmd.HasSubCommands()).To(BeFalse())
	})

	It("should provide a help prompt", func() {
		os.Args = append(os.Args, "--help")

		Expect(cmd.Execute()).Should(Succeed())
		Expect(fmt.Sprint(out)).To(ContainSubstring("set-network"))
	})
})

var _ = Describe("SetNetworkImplementation", func() {
	var (
		impl     *config.SetNetworkCommand
		err      error
		out      *bytes.Buffer
		settings *environment.Settings
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)

		settings = environment.NewDefaultSettings()
		settings.Home = environment.Home(os.TempDir())
		settings.Streams = environment.Streams{Out: out}

		impl = &config.SetNetworkCommand{}
		impl.Settings = settings
	})

	It("should not be nil", func() {
		Expect(impl).ShouldNot(BeNil())
	})

	Describe("Validate", func() {
		JustBeforeEach(func() {
			err = impl.Validate()
		})

		It("should fail without network name", func() {
			Expect(err).NotTo(BeNil())
		})

		Context("when network name is set", func() {
			BeforeEach(func() {
				impl.Name = "foo"
			})

			It("should successfully validate", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Run", func() {
		JustBeforeEach(func() {
			err = impl.Run()
		})

		JustAfterEach(func() {
			os.RemoveAll(impl.Settings.Home.Path(environment.DefaultConfigFilename))
		})

		Context("when a network exists", func() {
			BeforeEach(func() {
				impl.Name = "bar"
			})

			It("should set the network", func() {
				Expect(err).To(BeNil())
				Expect(fmt.Sprint(out)).To(ContainSubstring("successfully set network 'bar'"))
			})
		})

		Context("when config file is invalid", func() {
			BeforeEach(func() {
				data, _ := yaml.Marshal(struct {
					Networks string
				}{
					Networks: "foo",
				})

				os.MkdirAll(settings.Home.String(), 0777)

				ioutil.WriteFile(
					settings.Home.Path(environment.DefaultConfigFilename),
					data,
					0777,
				)
			})

			It("should fail to unmarshal config file", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
