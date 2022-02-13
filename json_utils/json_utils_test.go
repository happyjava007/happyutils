package json_utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestToJsonString(t *testing.T) {
	arr := []interface{}{10, 11, 12, 13, "apple"}
	m := map[string]interface{}{
		"name": "happyjava",
		"age":  18,
		"desc": nil,
		"arr":  arr,
	}
	t.Log(ToJsonString(m))
	t.Log(ToJsonString(nil))

	t.Log(ToJsonString(arr))
}

func TestParseObject(t *testing.T) {
	jsonStr := "{\"age\":18,\"arr\":[10,11,12,13,\"apple\"],\"desc\":null,\"name\":\"happyjava\"}"
	obj, err := ParseObject(jsonStr)
	if err != nil {
		t.Fatal(err)
	}
	name, err := obj.Get("name").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("name:" + name)
	array, err := obj.Get("arr").Array()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range array {
		t.Log("index:", i, "value:", v, "type:", reflect.TypeOf(v))
	}
}

func TestParseToStruct(t *testing.T) {
	type user struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Desc    string   `json:"desc"`
		Hobbies []string `json:"hobbies"`
	}
	u := &user{"happyjava", 18, "备注", []string{"apple", "banana"}}
	jsonString, _ := ToJsonString(u)
	t.Log(jsonString)

	var nu = &user{}
	if err := json.Unmarshal([]byte(jsonString), nu); err != nil {
		t.Fatal(err)
	}
	toJsonString, err := ToJsonString(nu)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(toJsonString)
}

func TestParseStruct(t *testing.T) {
	type user struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Desc    string   `json:"desc"`
		Hobbies []string `json:"hobbies"`
	}
	{
		jsonString, err := ToJsonString(&user{"happyjava", 18, "备注", []string{"apple", "banana"}})
		if err != nil {
			t.Fatal(err)
		}
		u := new(user)
		if err := ParseStruct(jsonString, u); err != nil {
			t.Fatal(err)
		}
		toJsonString, err := ToJsonString(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("new user:", toJsonString)
	}
	{
		users := []user{{"happyjava", 18, "备注", []string{"apple", "banana"}},
			{"happyjava", 18, "备注", []string{"apple", "banana"}}}
		jsonString, err := ToJsonString(users)
		if err != nil {
			t.Fatal(err)
		}
		newUsers := new([]user)
		if err := ParseStruct(jsonString, newUsers); err != nil {
			t.Fatal(err)
		}
		toJsonString, err := ToJsonString(newUsers)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("newUsers:", toJsonString)
	}
	{
		jsonStr := `{"name":"happyjava","age":18,"desc":"备注"}`
		u := new(user)
		if err := ParseStruct(jsonStr, u); err != nil {
			t.Fatal(err)
		}
		t.Log(ToJsonString(u))
	}
}

func TestToJsonStringIgnoreErr(t *testing.T) {
	arr := []string{"apple", "banana"}
	t.Log(ToJsonStringIgnoreErr(arr))
}
