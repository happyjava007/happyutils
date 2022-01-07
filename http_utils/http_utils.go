package http_utils

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{Timeout: time.Minute * 1}

func PostJsonReturnBody(url, json string) (string, error) {
	resp, err := PostJsonWhitHeaders(url, json, nil)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func PostJsonWhitHeaders(url, json string, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(json))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json;charset=utf8")
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	return client.Do(request)
}
