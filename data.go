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
	ToString() string
	ToInt() (int, error)
	ToFloat64() (float64, error)
	ToBool() (bool, error)
	Copy() Value
	FromString(string) error
}

// A StringValue represents a string configuration value
type StringValue string

// ToString returns the string representation of the value.
func (sv *StringValue) ToString() string {
	return string(*sv)
}

// ToInt returns the int representation of the value.
func (sv *StringValue) ToInt() (int, error) {
	return strconv.Atoi(string(*sv))
}

// ToFloat64 returns the float64 representation of the value.
func (sv *StringValue) ToFloat64() (float64, error) {
	return strconv.ParseFloat(string(*sv), 64)
}

// ToBool returns the bool representation of the value.
func (sv *StringValue) ToBool() (bool, error) {
	return strconv.ParseBool(string(*sv))
}

// Copy creates a deep copy of the string value.
func (sv *StringValue) Copy() Value {
	dup := *sv
	return &dup
}

// FromString updates the value from a string.
func (sv *StringValue) FromString(value string) error {
	*sv = StringValue(value)
	return nil
}

// An IntValue represents an int configuration value
type IntValue int

// ToString returns the string representation of the value
func (iv *IntValue) ToString() string {
	return strconv.Itoa(int(*iv))
}

// ToInt returns the int representation of the value.
func (iv *IntValue) ToInt() (int, error) {
	return int(*iv), nil
}

// ToFloat64 returns the float64 representation of the value.
func (iv *IntValue) ToFloat64() (float64, error) {
	return float64(*iv), nil
}

// ToBool returns the bool representation of the value.
func (iv *IntValue) ToBool() (bool, error) {
	return *iv != 0, nil
}

// Copy creates a deep copy of the int value.
func (iv *IntValue) Copy() Value {
	dup := *iv
	return &dup
}

// FromString updates the value from a string.
func (iv *IntValue) FromString(value string) error {
	val, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*iv = IntValue(val)
	return nil
}

// A FloatValue represents a float64 configuration value
type FloatValue float64

// ToString returns the string representation of the value
func (fv *FloatValue) ToString() string {
	return strconv.FormatFloat(float64(*fv), 'f', -1, 64)
}

// ToInt returns the int representation of the value.
func (fv *FloatValue) ToInt() (int, error) {
	return int(*fv), nil
}

// ToFloat64 returns the float64 representation of the value.
func (fv *FloatValue) ToFloat64() (float64, error) {
	return float64(*fv), nil
}

// ToBool returns the bool representation of the value.
func (fv *FloatValue) ToBool() (bool, error) {
	return *fv != 0, nil
}

// Copy creates a deep copy of the float value.
func (fv *FloatValue) Copy() Value {
	dup := *fv
	return &dup
}

// FromString updates the value from a string.
func (fv *FloatValue) FromString(value string) error {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	*fv = FloatValue(val)
	return nil
}

// A BoolValue represents a bool configuration value
type BoolValue bool

// ToString returns the string representation of the value
func (bv *BoolValue) ToString() string {
	return strconv.FormatBool(bool(*bv))
}

// ToInt returns the int representation of the value.
func (bv *BoolValue) ToInt() (int, error) {
	if *bv {
		return 1, nil
	}
	return 0, nil
}

// ToFloat64 returns the float64 representation of the value.
func (bv *BoolValue) ToFloat64() (float64, error) {
	if *bv {
		return 1, nil
	}
	return 0, nil
}

// ToBool returns the bool representation of the value.
func (bv *BoolValue) ToBool() (bool, error) {
	return bool(*bv), nil
}

// Copy creates a deep copy of the bool value.
func (bv *BoolValue) Copy() Value {
	dup := *bv
	return &dup
}

// FromString updates the value from a string.
func (bv *BoolValue) FromString(value string) error {
	val, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*bv = BoolValue(val)
	return nil
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

// createOption creates a new configuration option
func createOption(key string) *Option {
	return &Option{Key: key}
}
