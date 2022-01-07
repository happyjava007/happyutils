package rsa_utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
)

func GenerateKey(bits int) (privateKey string, publicKey string, err error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	privateBytes := x509.MarshalPKCS1PrivateKey(private)
	publicBytes := x509.MarshalPKCS1PublicKey(&private.PublicKey)
	return hex.EncodeToString(privateBytes), hex.EncodeToString(publicBytes), nil
}

func BuildPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	decodeString, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(decodeString)
}

func BuildPublicKey(publicKey string) (*rsa.PublicKey, error) {
	decodeString, err := hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(decodeString)
}

func EncryptByPublicKey(content string, publicKey *rsa.PublicKey) (string, error) {
	encrypt, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(content))
	return hex.EncodeToString(encrypt), err
}

func DecryptByPrivateKey(content string, privateKey *rsa.PrivateKey) (string, error) {
	decodeString, err := hex.DecodeString(content)
	if err != nil {
		return "", err
	}
	decrypt, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeString)
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}
