package util

import "github.com/mitchellh/mapstructure"

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	err := mapstructure.Decode(obj, &output)
	return output, err
}
