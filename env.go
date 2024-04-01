package appconf

import "os"

// UpdateFromEnv updates configuration option values from environment variables
func (conf *AppConf) UpdateFromEnv() error {
	for optKey, option := range conf.Options {
		val, ok := os.LookupEnv(option.Env)
		if ok {
			value := StringValue(val)
			switch option.Default.(type) {
			case IntValue:
				v, err := value.ToInt()
				if err != nil {
					return err
				}
				conf.Options[optKey].Value = IntValue(v)
			case FloatValue:
				v, err := value.ToFloat64()
				if err != nil {
					return err
				}
				conf.Options[optKey].Value = FloatValue(v)
			case BoolValue:
				v, err := value.ToBool()
				if err != nil {
					return err
				}
				conf.Options[optKey].Value = BoolValue(v)
			case StringValue:
				conf.Options[optKey].Value = value
			default:
				return ErrInvalidType
			}
		}
	}
	return nil
}
