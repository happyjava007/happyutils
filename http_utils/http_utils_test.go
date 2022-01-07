package http_utils

import "testing"

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
