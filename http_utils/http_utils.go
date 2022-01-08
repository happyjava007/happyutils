package http_utils

import (
	"fmt"
	"github.com/happyjava007/happyutils/encode_utils"
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
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(respBytes), nil
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

func PostFormWithHeaders(url string, headers map[string]string, data map[string]string) (*http.Response, error) {
	var strArr []string
	payloadData := ""

	if data != nil && len(data) > 0 {
		for k, v := range data {
			strArr = append(strArr, fmt.Sprintf("%s=%s", encode_utils.UrlEncoding(k), encode_utils.UrlEncoding(v)))
		}
		payloadData = strings.Join(strArr, "&")
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(payloadData))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	return client.Do(request)
}

func PostFormReturnBody(url string, data map[string]string) (string, error) {
	resp, err := PostFormWithHeaders(url, nil, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}
