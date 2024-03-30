// Package appconf is a lightweight configuration solution for Go applications.
// It helps manage configuration settings and handle configuration inputs from
// different sources.
//
// Currently, this package supports the following configuration sources:
//
//   - JSON Files
//   - Environment Variables
//   - Command Line Flags
//
// Configuration directives are interpreted following this precedence order:
//
//  1. Command Line Flags
//  2. Environment Variables
//  3. Configuration File
//  4. Default Values
//
// Settings provided by an instance with a lower precedence order (i.e. higher priority)
// will always override those with a higher precedence order (i.e. lower priority).
package appconf

// An AppConf instance represents a configuration context for an application.
type AppConf struct {
	Options   []Option
	ConfFiles []string
}

// A ConfOption is a functional option for configuring an AppConf context
type ConfOption func(*AppConf)

// WithConfFile sets a single configuration file to be read
func WithConfFile(confFile string) ConfOption {
	return func(c *AppConf) {
		c.ConfFiles = make([]string, 1)
		c.ConfFiles[0] = confFile
	}
}

// WithConfFiles sets a list of configuration files to be read
func WithConfFiles(confFiles []string) ConfOption {
	return func(c *AppConf) {
		c.ConfFiles = confFiles
	}
}

// New creates a new AppConf context
func New(options ...ConfOption) *AppConf {
	conf := &AppConf{}
	for _, option := range options {
		option(conf)
	}
	return conf
}
