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

// A AppOption is a functional option for configuring an AppConf context
type AppOption func(*AppConf)

// WithAuthor sets the application author or publisher
func WithAuthor(author string) AppOption {
	return func(conf *AppConf) {
		conf.Author = author
	}
}

// WithConfFile sets a single configuration file to be read
func WithConfFile(confFile string) AppOption {
	return func(conf *AppConf) {
		conf.ConfFiles = make([]string, 1)
		conf.ConfFiles[0] = confFile
	}
}

// WithConfFiles sets a list of configuration files to be read
func WithConfFiles(confFiles []string) AppOption {
	return func(conf *AppConf) {
		conf.ConfFiles = confFiles
	}
}

// WithRoaming sets the roaming flag (applies to Windows only)
func WithRoaming() AppOption {
	return func(conf *AppConf) {
		conf.Roaming = true
	}
}

// WithVersion sets the application version
func WithVersion(version string) AppOption {
	return func(conf *AppConf) {
		conf.Version = version
	}
}

// NewConf creates a new AppConf context
func NewConf(appName string, options ...AppOption) *AppConf {
	conf := &AppConf{Name: appName, Roaming: false}
	conf.Options = make(map[string]*Option)
	for _, option := range options {
		option(conf)
	}
	return conf
}

// A OptOption is a functional option for configuring an Option object
type OptOption func(option *Option)

// WithDefaultInt sets the default int value for an option
func WithDefaultInt(value int) OptOption {
	return func(opt *Option) {
		v := IntValue(value)
		opt.Default = v.Copy()
		opt.Value = v.Copy()
	}
}

// WithDefaultFloat sets the default float64 value for an option
func WithDefaultFloat(value float64) OptOption {
	return func(opt *Option) {
		v := FloatValue(value)
		opt.Default = v.Copy()
		opt.Value = v.Copy()
	}
}

// WithDefaultBool sets the default bool value for an option
func WithDefaultBool(value bool) OptOption {
	return func(opt *Option) {
		v := BoolValue(value)
		opt.Default = v.Copy()
		opt.Value = v.Copy()
	}
}

// WithDefaultString sets the default string value for an option
func WithDefaultString(value string) OptOption {
	return func(opt *Option) {
		v := StringValue(value)
		opt.Default = v.Copy()
		opt.Value = v.Copy()
	}
}

// WithFlag sets the command line flag for an option
func WithFlag(flag string) OptOption {
	return func(opt *Option) {
		opt.Flag = flag
	}
}

// WithJson sets the JSON address for an option
func WithJson(json string) OptOption {
	return func(opt *Option) {
		opt.Json = json
	}
}

// WithEnv sets the environment variable for an option
func WithEnv(env string) OptOption {
	return func(opt *Option) {
		opt.Env = env
	}
}

// WithHelp sets the help text for an option
func WithHelp(help string) OptOption {
	return func(opt *Option) {
		opt.Help = help
	}
}

// NewOption creates and registers a new Option within the AppConf context
func (conf *AppConf) NewOption(key string, options ...OptOption) error {
	_, ok := conf.Options[key]
	if ok {
		return ErrOptionExists
	}
	opt := createOption(key)
	for _, option := range options {
		option(opt)
	}
	conf.Options[key] = opt
	return nil
}
