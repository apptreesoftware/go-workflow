package step

import (
	"bytes"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/core"
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
)

type Context interface {
	BindInputs(data interface{}) error
	InputMap() (map[string]interface{}, error)
	Environment() *core.Environment
	Engine() Engine
}

type ContextBase struct {
	environment *core.Environment
}

func (c ContextBase) Environment() *core.Environment {
	return c.environment
}

func (c ContextBase) Engine() Engine {
	return GetEngine(c.environment)
}

type IpcContext struct {
	ContextBase
}

func (IpcContext) BindInputs(data interface{}) error {
	return readInputs(os.Stdin, data)
}

func (IpcContext) InputMap() (map[string]interface{}, error) {
	mapData := map[string]interface{}{}
	err := readInputs(os.Stdin, &mapData)
	return mapData, err
}

type RpcContext struct {
	bytes []byte
	ContextBase
}

func (b *RpcContext) BindInputs(data interface{}) error {
	return readInputs(bytes.NewReader(b.bytes), data)
}

func (b *RpcContext) InputMap() (map[string]interface{}, error) {
	reader := bytes.NewReader(b.bytes)
	mapData := map[string]interface{}{}
	err := readInputs(reader, &mapData)
	return mapData, err
}

func readInputs(reader io.Reader, data interface{}) error {
	var buf bytes.Buffer
	tee := io.TeeReader(reader, &buf)
	dec := jsoniter.NewDecoder(tee)
	readErr := dec.Decode(data)
	if readErr != nil {
		return fmt.Errorf("Unable to read step inputs. The data provided was not in a valid json format.\nData: %s", buf.String())
	}
	return nil
}