package rpc_test

import (
	"monkeylsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(content) != 15 {
		t.Fatalf("Content length was not 15. got=%d", len(content))
	}

	if method != "hi" {
		t.Fatalf("method not 'hi'. got=%s", method)
	}
}
