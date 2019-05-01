/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package environment

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

// DefaultConfigFilename is the default config filename
const DefaultConfigFilename = "config.yaml"

// Config contains information needed to manage fabric networks
type Config struct {
	Networks       map[string]*Network `yaml:",omitempty"`
	Contexts       map[string]*Context `yaml:",omitempty"`
	CurrentContext string              `yaml:"current-context,omitempty"`
}

// Network contains a fabric network's configurations
type Network struct {
	// path to fabric go sdk config file
	ConfigPath string `yaml:"path,omitempty"`
}

// Context contains network interaction parameters
type Context struct {
	Network      string   `yaml:",omitempty"`
	Organization string   `yaml:",omitempty"`
	User         string   `yaml:",omitempty"`
	Channel      string   `yaml:",omitempty"`
	Orderers     []string `yaml:",omitempty"`
	Peers        []string `yaml:",omitempty"`
}

// AddFlags appeneds config flags onto an existing flag set
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.CurrentContext, "context", c.CurrentContext, "override the current context")
}

// LoadFromFile populates config based on the specified path
func (c *Config) LoadFromFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}

// Save writes it's current value to the specified path
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, data, 0600); err != nil {
		return err
	}

	return nil
}

// GetCurrentContext returns the current context
func (c *Config) GetCurrentContext() (*Context, error) {
	if c == nil || len(c.CurrentContext) == 0 {
		return nil, errors.New("current context is not set")
	}

	context, ok := c.Contexts[c.CurrentContext]
	if !ok {
		return nil, errors.New("current context does not exist")
	}

	return context, nil
}

// GetCurrentNetwork returns the current network
func (c *Config) GetCurrentNetwork() (*Network, error) {
	context, err := c.GetCurrentContext()
	if err != nil {
		return nil, err
	}

	if len(context.Network) == 0 {
		return nil, errors.New("current network is not set")
	}

	network, ok := c.Networks[context.Network]
	if !ok {
		return nil, errors.New("current network does not exist")
	}

	return network, nil
}

// NewConfig returns a new config
func NewConfig() *Config {
	return &Config{
		Networks: make(map[string]*Network),
		Contexts: make(map[string]*Context),
	}
}
