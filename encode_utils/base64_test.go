package encode_utils

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestBase64Encoding(t *testing.T) {
	data := "哈皮niubility!@#$%^&^&*(*)_       =0"
	encoding := Base64Encoding(data)
	t.Log(encoding)
	decoding, err := Base64Decoding(encoding)
	t.Log(err)
	t.Log(decoding)
	assert.Equal(t, data, decoding)
}
