/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package environment

// Action is used to modify config
type Action func(c *Config)

// SetCurrentContext updates the current context
func SetCurrentContext(context string) Action {
	return func(c *Config) {
		c.CurrentContext = context
	}
}

// SetContext adds or updates the specified context
func SetContext(name string, context *Context) Action {
	return func(c *Config) {
		c.Contexts[name] = context
	}
}

// DeleteContext deletes a specified context
func DeleteContext(name string) Action {
	return func(c *Config) {
		delete(c.Contexts, name)
	}
}

// SetNetwork adds or updates the specified network
func SetNetwork(name string, network *Network) Action {
	return func(c *Config) {
		c.Networks[name] = network
	}
}

// DeleteNetwork deletes the specified network
func DeleteNetwork(name string) Action {
	return func(c *Config) {
		delete(c.Networks, name)
	}
}
