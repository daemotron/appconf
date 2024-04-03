package appconf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// traverseJsonFile recursively traverses the JSON structure and assembles keys
// with nested keys separated by a delimiter
func traverseJsonFile(data interface{}, prefix string) (map[string]Value, error) {
	result := make(map[string]Value)
	switch value := data.(type) {
	case map[string]interface{}:
		for key, val := range value {
			nestedKey := fmt.Sprintf("%s%s%s", prefix, key, ".")
			res, err := traverseJsonFile(val, nestedKey)
			if err != nil {
				return nil, err
			}
			result = mergeMaps(result, res)
		}
	case []interface{}:
		for i, val := range value {
			nestedKey := fmt.Sprintf("%s%d%s", prefix, i, ".")
			res, err := traverseJsonFile(val, nestedKey)
			if err != nil {
				return nil, err
			}
			result = mergeMaps(result, res)
		}
	case int:
		key := strings.TrimSuffix(prefix, ".")
		iv := IntValue(value)
		result[key] = iv.Copy()
	case string:
		key := strings.TrimSuffix(prefix, ".")
		sv := StringValue(value)
		result[key] = sv.Copy()
	case float64:
		key := strings.TrimSuffix(prefix, ".")
		fv := FloatValue(value)
		result[key] = fv.Copy()
	case bool:
		key := strings.TrimSuffix(prefix, ".")
		bv := BoolValue(value)
		result[key] = bv.Copy()
	default:
		return nil, ErrInvalidType
	}
	return result, nil
}

// parseJsonFile reads an arbitrary JSON file into a flat key/value map, where nested
// keys are represented by address strings
func parseJsonFile(path string) (map[string]Value, error) {
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

	result, err := traverseJsonFile(data, "")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateFromJsonFile updates configuration options with data extracted from
// the specified JSON file
func (conf *AppConf) updateFromJsonFile(path string) error {
	data, err := parseJsonFile(path)
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
