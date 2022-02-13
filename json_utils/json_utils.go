package json_utils

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
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

func ToJsonStringIgnoreErr(v interface{}) string {
	if v == nil {
		return "{}"
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Println("ToJsonStringIgnoreErr error:", err)
		return "{}"
	}
	return string(bytes)
}

func ParseObject(json string) (*simplejson.Json, error) {
	return simplejson.NewJson([]byte(json))
}

func ParseStruct(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
