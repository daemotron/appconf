package appconf

import "os"

// UpdateFromEnv updates configuration option values from environment variables
func (conf *AppConf) UpdateFromEnv() error {
	for optKey, option := range conf.Options {
		val, ok := os.LookupEnv(option.Env)
		if ok {
			value := StringValue(val)
			switch option.Default.(type) {
			case *IntValue:
				v, err := value.ToInt()
				if err != nil {
					return err
				}
				iv := IntValue(v)
				conf.Options[optKey].Value = iv.Copy()
			case *FloatValue:
				v, err := value.ToFloat64()
				if err != nil {
					return err
				}
				fv := FloatValue(v)
				conf.Options[optKey].Value = fv.Copy()
			case *BoolValue:
				v, err := value.ToBool()
				if err != nil {
					return err
				}
				bv := BoolValue(v)
				conf.Options[optKey].Value = bv.Copy()
			case *StringValue:
				conf.Options[optKey].Value = value.Copy()
			default:
				return ErrInvalidType
			}
		}
	}
	return nil
}
