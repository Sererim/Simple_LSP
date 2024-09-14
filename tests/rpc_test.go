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
	incoming_msg := "Content-Length: 17\r\n\r\n{\"Method\":\"test\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming_msg))
	if err != nil {
		t.Fatal(err)
	}
	content_length := len(content)

	if content_length != 17 {
		t.Fatalf("Expected: 15\nActual: %d", content_length)
	}

	if method != "test" {
		t.Fatalf("Expected: 'test', Got: %s", method)
	}

}
