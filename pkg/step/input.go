package step

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
)

type Context interface {
	BindInputs(data interface{}) error
	InputMap() (map[string]interface{}, error)
}

type StdInput struct {

}

func (StdInput) BindInputs(data interface{}) error {
	return readInputs(os.Stdin, data)
}

func (StdInput) InputMap() (map[string]interface{}, error) {
	mapData := map[string]interface{}{}
	err := readInputs(os.Stdin, &mapData)
	return mapData, err
}

type ByteInput struct {
	bytes []byte
}

func (b *ByteInput) BindInputs(data interface{}) error {
	return readInputs(bytes.NewReader(b.bytes), data)
}

func (b *ByteInput) InputMap() (map[string]interface{}, error) {
	reader := bytes.NewReader(b.bytes)
	mapData := map[string]interface{}{}
	err := readInputs(reader, &mapData)
	return mapData, err
}

func readInputs(reader io.Reader, data interface{}) error {
	dec := jsoniter.NewDecoder(reader)
	return dec.Decode(data)
}