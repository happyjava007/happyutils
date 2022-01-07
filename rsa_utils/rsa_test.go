package rsa_utils

import "testing"

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey, err := GenerateKey(2048)
	t.Log(privateKey)
	t.Log(publicKey)
	t.Log(err)
}

func TestEncrypt(t *testing.T) {
	privateKey, publicKey, _ := GenerateKey(2048)
	content := "happyjava"

	priKey, _ := BuildPrivateKey(privateKey)
	pubKey, _ := BuildPublicKey(publicKey)

	encrypt, _ := EncryptByPublicKey(content, pubKey)
	t.Log(encrypt)

	decrypt, _ := DecryptByPrivateKey(encrypt, priKey)
	t.Log(decrypt)

}
