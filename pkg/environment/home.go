/*
Copyright © 2019 State Street Bank and Trust Company.  All rights reserved

SPDX-License-Identifier: Apache-2.0
*/

package environment

import (
	"os"
	"path/filepath"
)

// DefaultHome is the default directory for the user
const DefaultHome = Home("${HOME}/.fabctl")

// Home is the location of the configuration files
// By default, the files are stored in ~/.fab
type Home string

// String resolves variables and returns home as a string
func (h Home) String() string {
	return os.ExpandEnv(string(h))
}

// Path appends compoenents to home and returns it as a string
func (h Home) Path(components ...string) string {
	return filepath.Join(append([]string{h.String()}, components...)...)
}

// Plugins returns the path to the plugins directory
func (h Home) Plugins() string {
	return h.Path("plugins")
}
