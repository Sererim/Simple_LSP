package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(
		"Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.
		Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("ERROR: no separator found!")
	}

	// Content-Length: <NUMBER>

	content_length_bytes := header[len("Content-Length: "):]
	content_length, err := strconv.Atoi(string(content_length_bytes))
	if err != nil {
		return 0, err
	}

	// TODO: finish work with content.
	_ = content

	return content_length, nil
}
