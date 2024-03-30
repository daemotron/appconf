package appconf

import "strconv"

// A Value instance represents a generic configuration value
// and can be of one of these types:
//
//   - string
//   - int
//   - float64
//   - bool
//
// Please note: this is an abstract interface, use the type-specific
// implementations to actually handle configuration values
type Value interface {
	String() string
}

// A StringValue represents a string configuration value
type StringValue string

// String returns the string representation of the value.
func (sv StringValue) String() string {
	return string(sv)
}

// An IntValue represents an int configuration value
type IntValue int

// String returns the string representation of the value
func (iv IntValue) String() string {
	return strconv.Itoa(int(iv))
}

// A FloatValue represents a float64 configuration value
type FloatValue float64

// String returns the string representation of the value
func (fv FloatValue) String() string {
	return strconv.FormatFloat(float64(fv), 'f', -1, 64)
}

// A BoolValue represents a bool configuration value
type BoolValue bool

// String returns the string representation of the value
func (bv BoolValue) String() string {
	return strconv.FormatBool(bool(bv))
}

// An Option represents a configuration option
type Option struct {
	Key     string // Key identifies the option and shall be unique
	Default Value  // Default represents the default option value
	Value   Value  // Value represents the current option value
	Flag    string // Flag represents the option's command line flag
	Json    string // Json represents the option's JSON address
	Env     string // Env represents the option's environment variable
	Help    string // Help represents a help string describing the option
}
