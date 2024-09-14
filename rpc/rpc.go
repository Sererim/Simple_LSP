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

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.
		Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("ERROR: no separator found")
	}

	content_length_bytes := header[len("Content-Length: "):]

	content_length, err := strconv.Atoi(string(content_length_bytes))
	if err != nil {
		return "", nil, err
	}

	var base_msg BaseMessage
	if err := json.
		Unmarshal(content[:content_length], &base_msg); err != nil {
		return "", nil, err
	}

	return base_msg.Method, content[:content_length], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {

	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	content_length_bytes := header[len("Content-Length: "):]
	content_length, err := strconv.Atoi(string(content_length_bytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < content_length {
		return 0, nil, nil
	}

	total_length := len(header) + 4 + content_length

	return total_length, data[:total_length], nil
}
