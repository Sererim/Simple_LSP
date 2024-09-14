package rpc_test

import (
	"simple_lsp/rpc"
	"testing"
)

type EncodingTest struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingTest{Testing: true})
	if expected != actual {
		t.Fatalf("\nExpected: %s\nActual: %s", expected, actual)
	}
}
