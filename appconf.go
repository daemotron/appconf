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
	Options   map[string]*Option
	ConfFiles []string
	Name      string
	Author    string
	Version   string
	Roaming   bool
}

// A ConfOption is a functional option for configuring an AppConf context
type ConfOption func(*AppConf)

// WithAuthor sets the application author or publisher
func WithAuthor(author string) ConfOption {
	return func(conf *AppConf) {
		conf.Author = author
	}
}

// WithConfFile sets a single configuration file to be read
func WithConfFile(confFile string) ConfOption {
	return func(conf *AppConf) {
		conf.ConfFiles = make([]string, 1)
		conf.ConfFiles[0] = confFile
	}
}

// WithConfFiles sets a list of configuration files to be read
func WithConfFiles(confFiles []string) ConfOption {
	return func(conf *AppConf) {
		conf.ConfFiles = confFiles
	}
}

// WithRoaming sets the roaming flag (applies to Windows only)
func WithRoaming() ConfOption {
	return func(conf *AppConf) {
		conf.Roaming = true
	}
}

// WithVersion sets the application version
func WithVersion(version string) ConfOption {
	return func(conf *AppConf) {
		conf.Version = version
	}
}

// NewConf creates a new AppConf context
func NewConf(appName string, options ...ConfOption) *AppConf {
	conf := &AppConf{Name: appName, Roaming: false}
	conf.Options = make(map[string]*Option)
	for _, option := range options {
		option(conf)
	}
	return conf
}
