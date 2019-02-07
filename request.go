package go_workflow

import (
	"github.com/json-iterator/go"
	"io"
	"os"

)

func BindInputs(data interface{}) error {
	return ReadInputs(os.Stdin, data)
}

func SetOutput(data interface{}) error  {
	return WriteOutputsTo(os.Stdout, data)
}

func ReadInputs(reader io.Reader, data interface{}) error {
	dec := jsoniter.NewDecoder(reader)
	var req Request
	err := dec.Decode(&req)
	if err != nil {
		return err
	}
	return req.InputsTo(data)
}

func WriteOutputsTo(writer io.Writer, data interface{}) error {
	response := Response{
		Outputs: data,
	}
	enc := jsoniter.NewEncoder(writer)
	return enc.Encode(&response)
}

type Request struct {
	Inputs jsoniter.RawMessage
}

type Response struct {
	Outputs interface{}
}

func (req *Request) InputsTo(data interface{}) error {
	return jsoniter.Unmarshal(req.Inputs, data)
}
