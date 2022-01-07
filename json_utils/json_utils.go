package json_utils

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

func ToJsonString(v interface{}) (string, error) {
	if v == nil {
		return "{}", nil
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ParseObject(json string) (*simplejson.Json, error) {
	return simplejson.NewJson([]byte(json))
}

func ParseStruct(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
