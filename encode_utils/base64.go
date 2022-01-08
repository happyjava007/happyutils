package encode_utils

import "encoding/base64"

func Base64Encoding(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decoding(data string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}
