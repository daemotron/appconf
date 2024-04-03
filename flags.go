package appconf

import "flag"

var registeredFlags = make(map[string]bool)
var flagActions = make(map[string]bool)

// registerFlags registers all defined option flags with Go's flag package.
func (conf *AppConf) registerFlags() error {
	for _, option := range conf.Options {
		if option.Flag != "" {
			if !registeredFlags[option.Flag] {
				switch v := option.Value.(type) {
				case *IntValue:
					iv, err := option.Default.ToInt()
					if err != nil {
						return err
					}
					flag.IntVar((*int)(v), option.Flag, iv, option.Help)
				case *FloatValue:
					fv, err := option.Default.ToFloat64()
					if err != nil {
						return err
					}
					flag.Float64Var((*float64)(v), option.Flag, fv, option.Help)
				case *BoolValue:
					bv, err := option.Default.ToBool()
					if err != nil {
						return err
					}
					flag.BoolVar((*bool)(v), option.Flag, bv, option.Help)
				case *StringValue:
					flag.StringVar((*string)(v), option.Flag, option.Default.ToString(), option.Help)
				default:
					return ErrInvalidType
				}
				registeredFlags[option.Flag] = true
			}
		}
	}
	return nil
}

// UpdateFromFlags updates configuration options from command line flags
func (conf *AppConf) UpdateFromFlags() error {
	if flagActions["parse"] {
		return ErrFlagsAlreadyParsed
	}
	if !flagActions["register"] {
		err := conf.registerFlags()
		if err != nil {
			return err
		}
		flagActions["register"] = true
	}
	flag.Parse()
	flagActions["parse"] = true

	// Update option values from parsed flags
	for _, option := range conf.Options {
		if option.Flag != "" {
			checkFlag := flag.Lookup(option.Flag)
			if checkFlag != nil {
				err := option.Value.FromString(checkFlag.Value.String())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
