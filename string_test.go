package string

import (
	"encoding/base64"
	"testing"
)

func TestReverseSecretToPlainText(t *testing.T) {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	s := String{}

	whatIsIt = s.Reverse(sd)

	if whatIsIt != "Join:us:at:LINE:MAN:Wongnai" {
		t.Errorf("Expect Join:us:at:LINE:MAN:Wongnai but receive %s", whatIsIt)
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := String{}
		s.Reverse([]byte("iangnoW:NAM:ENIL:ta:su:nioJ"))
	}
}
