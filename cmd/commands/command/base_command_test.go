/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package command_test

import (
	"testing"

	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger/fabric-cli/cmd/commands/command"
)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Command Suite")
}

var _ = Describe("BaseCommand", func() {
	var (
		c *command.BaseCommand
	)

	BeforeEach(func() {
		c = &command.BaseCommand{}
	})

	Describe("AddArg", func() {
		It("should not add nil arg", func() {
			c.AddArg(nil)
			Expect(len(c.Args)).To(BeZero())
		})

		It("should add arg", func() {
			var foo string
			c.AddArg(&foo)
			Expect(len(c.Args)).To(Equal(1))
		})
	})

	Describe("ParseArgs", func() {
		var (
			f cobra.PositionalArgs
		)

		JustBeforeEach(func() {
			f = c.ParseArgs()
		})

		It("should do nothing when Args are nil", func() {
			f := c.ParseArgs()
			f(nil, []string{"foo"})
		})

		It("should do nothing when input args are nil", func() {
			var foo string
			c.AddArg(&foo)
			f(nil, nil)
			Expect(len(foo)).To(BeZero())
		})

		It("should set variable", func() {
			var foo string
			c.AddArg(&foo)
			f(nil, []string{"foo"})
			Expect(foo).To(Equal("foo"))
		})
	})
})
