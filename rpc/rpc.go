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
		panic(err) // TODO: change this if we care
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find seperator")
	}

	contentLengthBytes := header[len("Content-Length: "):] // slice everything after the set intro

	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil

}

// scanner split function to read one message at a time while doing some basic checking on length
func Split(data []byte, _ bool) (adavance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil // Wait
	}

	contentLengthBytes := header[len("Content-Length: "):] // slice everything after the set intro

	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		// haven't read enough bytes from content stream
		return 0, nil, nil // wait
	}

	seperatorLength := 4
	totalLength := len(header) + seperatorLength + contentLength

	return totalLength, data[:totalLength], nil
}
