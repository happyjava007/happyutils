package http_utils

import (
	"github.com/happyjava007/happyutils/encode_utils"
	"io/ioutil"
	"testing"
)

func TestPostJson(t *testing.T) {
	json := `
{
    "days": 3650,
    "email": null
}`
	result, err := PostJsonReturnBody("https://app.happyjava.cn/github/postHosta", json)
	t.Log(err)
	t.Log(result)
}

func TestPostFormWhitHeaders(t *testing.T) {
	resp, err := PostFormWithHeaders("http://localhost:8080/testFormData",
		map[string]string{"token": encode_utils.UrlEncoding("哈皮niubility")},
		map[string]string{"name": "happyjava", "desc": "niubility&&asf哈皮！@#￥%……&!@#$%^&&*()_+=--089"})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	resultBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(resultBytes))
}

func TestPostFormReturnBody(t *testing.T) {
	body, err := PostFormReturnBody("http://localhost:8080/testFormData",
		map[string]string{"name": "happyjava", "desc": "niubility&&asf哈皮！@#￥%……&!@#$%^&&*()_+=--089"})
	t.Log(err)
	t.Log(body)
}
