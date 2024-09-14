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

func TestDecode(t *testing.T) {
	incoming_msg := "Content-Length: 16\r\n\r\n{\"Testing\"true}"
	content_length, err := rpc.DecodeMessage([]byte(incoming_msg))
	if err != nil {
		t.Fatal(err)
	}

	if content_length != 16 {
		t.Fatalf("Expected: 16\nActual: %s", content_length)
	}
}
