/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package environment_test

import (
	"github.com/hyperledger/fabric-cli/pkg/environment"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Actions", func() {
	var (
		config *environment.Config
		action environment.Action
	)

	BeforeEach(func() {
		config = environment.NewConfig()
	})

	JustBeforeEach(func() {
		action(config)
	})

	Describe("SetCurrentContext", func() {
		BeforeEach(func() {
			action = environment.SetCurrentContext("foo")
		})

		It("should set current context", func() {
			Expect(config.CurrentContext).To(Equal("foo"))
		})
	})

	Describe("SetContext", func() {
		BeforeEach(func() {
			action = environment.SetContext("foo", &environment.Context{})
		})

		It("should set context", func() {
			Expect(config.Contexts).To(HaveKey("foo"))
		})
	})

	Describe("DeleteContext", func() {
		BeforeEach(func() {
			config = &environment.Config{
				Contexts: map[string]*environment.Context{
					"foo": {},
				},
			}
			action = environment.DeleteContext("foo")
		})

		It("should set context", func() {
			Expect(config.Contexts).NotTo(HaveKey("foo"))
		})
	})

	Describe("SetNetwork", func() {
		BeforeEach(func() {
			action = environment.SetNetwork("foo", &environment.Network{})
		})

		It("should set network", func() {
			Expect(config.Networks).To(HaveKey("foo"))
		})
	})

	Describe("DeleteNetwork", func() {
		BeforeEach(func() {
			config = &environment.Config{
				Networks: map[string]*environment.Network{
					"foo": {},
				},
			}
			action = environment.DeleteNetwork("foo")
		})

		It("should set context", func() {
			Expect(config.Networks).NotTo(HaveKey("foo"))
		})
	})
})
