package appconf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TraverseJsonFile recursively traverses the JSON structure and assembles keys
// with nested keys separated by a delimiter
func TraverseJsonFile(data interface{}, prefix string) (map[string]Value, error) {
	result := make(map[string]Value)
	switch value := data.(type) {
	case map[string]interface{}:
		for key, val := range value {
			nestedKey := fmt.Sprintf("%s%s%s", prefix, key, ".")
			res, err := TraverseJsonFile(val, nestedKey)
			if err != nil {
				return nil, err
			}
			result = MergeMaps(result, res)
		}
	case []interface{}:
		for i, val := range value {
			nestedKey := fmt.Sprintf("%s%d%s", prefix, i, ".")
			res, err := TraverseJsonFile(val, nestedKey)
			if err != nil {
				return nil, err
			}
			result = MergeMaps(result, res)
		}
	case int:
		key := strings.TrimSuffix(prefix, ".")
		result[key] = IntValue(value)
	case string:
		key := strings.TrimSuffix(prefix, ".")
		result[key] = StringValue(value)
	case float64:
		key := strings.TrimSuffix(prefix, ".")
		result[key] = FloatValue(value)
	case bool:
		key := strings.TrimSuffix(prefix, ".")
		result[key] = BoolValue(value)
	default:
		return nil, ErrInvalidType
	}
	return result, nil
}

// ParseJsonFile reads an arbitrary JSON file into a flat key/value map, where nested
// keys are represented by address strings
func ParseJsonFile(path string) (map[string]Value, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}(file)

	var data interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	result, err := TraverseJsonFile(data, "")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFromJsonFile updates configuration options with data extracted from
// the specified JSON file
func (conf *AppConf) UpdateFromJsonFile(path string) error {
	data, err := ParseJsonFile(path)
	if err != nil {
		return err
	}
	for key, value := range data {
		for optKey, option := range conf.Options {
			if option.Json == key {
				conf.Options[optKey].Value = value
			}
		}
	}
	return nil
}
