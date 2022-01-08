package encode_utils

import (
	"net/url"
	"testing"
)

func TestCommon(t *testing.T) {
	escape := url.QueryEscape("%&=")
	t.Log(escape)
	unescape, err := url.QueryUnescape(escape)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(unescape)

	t.Log(url.QueryEscape(""))
}

func TestUrlEncoding(t *testing.T) {
	encoding := UrlEncoding("&**)+^%%$&")
	t.Log(encoding)
	decoding, err := UrlDecoding(encoding)
	t.Log(decoding)
	t.Log(err)
}
