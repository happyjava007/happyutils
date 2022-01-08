package encode_utils

import "net/url"

func UrlEncoding(data string) string {
	return url.QueryEscape(data)
}

func UrlDecoding(data string) (string, error) {
	return url.QueryUnescape(data)
}
