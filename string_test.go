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

func TestReversePlainTextToSecret(t *testing.T) {
	var secret string

	plainText := "Join:us:at:LINE:MAN:Wongnai"

	s := String{}

	reversedPlainText := s.Reverse([]byte(plainText))

	secret = base64.StdEncoding.EncodeToString([]byte(reversedPlainText))

	if secret != "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K" {
		t.Errorf("Expect aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K")
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := String{}
		s.Reverse([]byte("iangnoW:NAM:ENIL:ta:su:nioJ"))
	}
}
