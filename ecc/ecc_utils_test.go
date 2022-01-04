package ecc

import "testing"

const pubKey = "MTEzODUyMzM5MzU4ODk1MzEzNzg3MzY2NzYwODUwMjE2MjQzNTE0OTc0NjkzMzM2NzA4OTYzNDY0MDk1MTUxMDkyMjY5MjM0NjE3NjMyKzc5NzY5NDM5NTI4NjcxMzc4OTQ3ODYyMzY0MjQ1ODg4ODA0MTEyMDgwOTM4MjgyNjI0MTY4NDUwNjE0NTg4MDc3MDk0MTUxNzk4NzM3"
const priKey = "MHcCAQEEIBC7/AKbbIHyVF4XJQmvqwrJMX8c8dU2JP6NcReYDlJ1oAoGCCqGSM49AwEHoUQDQgAE+7Yj9jOgDCRccssUMp1NVJExBrJCv6H8LsYUqS8lfSCwW+cdXcnmNDHM5Z2K05bJywyDIWU3f+53z0HK0I4/0Q=="
const content = "Happyjava not only java"
const signature = "36393733373630313538383036323831323937313232353733313837333336373132363935303734303237373233373433363139373530313438383239303532303635383433333733313336362b35363137383339343939383534303631323730373630343438313135333633383339333535303836313830343737333331383935333532383537333730363035313839383939353330363939"

func TestGenKeyPair(t *testing.T) {
	privateKey, publicKey, e := GenKeyPair()
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("pubKey:", publicKey)
	t.Log("priKey:", privateKey)
}

func TestBuildPublicKey(t *testing.T) {
	publicKey, e := BuildPublicKey(pubKey)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("publicKey:", publicKey)
}

func TestBuildPrivateKey(t *testing.T) {
	privateKey, e := BuildPrivateKey(priKey)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("privateKey:", privateKey)
}

func TestSign(t *testing.T) {
	signature, e := Sign(content, priKey)
	if e != nil {
		t.Error(e)
	} else {
		t.Log("signature:", signature)
	}
}

func TestVerifySign(t *testing.T) {
	verify := VerifySign(content, signature, pubKey)
	t.Log("verify:", verify)
}
