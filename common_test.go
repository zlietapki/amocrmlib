package amocrmlib

import (
	"os"
	"testing"
)

var token *Token

func TestMain(m *testing.M) {
	const TokenFile = "token.json"

	var err error
	token, err = LoadTokenFile(TokenFile)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
