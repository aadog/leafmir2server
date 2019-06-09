package main

import (
	"fmt"
	"leafmir2server/base"
	"testing"
)

func Test_enc(t *testing.T) {
	decbase64 := base.Base64DecodeEx_EDcode([]byte("b^jYq]DsVwTAEQW&ez&ge&UgZQlpgym>URjO&<IQHAe="), 32)
	enc1 := decbase64[:16]
	dec1 := base.DecryptAES_EDcode(enc1)
	fmt.Println(dec1)
}
