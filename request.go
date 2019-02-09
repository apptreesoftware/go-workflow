package go_workflow

import (
	"github.com/json-iterator/go"
	"io"
	"os"
)

func BindInputs(data interface{}) {
	err := ReadInputs(os.Stdin, data)
	if err != nil {
		ReportError(err)
	}
}

func ReportError(err error) {
	println(err.Error())
	os.Exit(1)
}

func SetOutput(data interface{}) {
	writer, err := getDefaultOutput()
	if err != nil {
		ReportError(err)
	}
	defer writer.Close()
	err = WriteOutputsTo(writer, data)
	if err != nil {
		ReportError(err)
	}
}

func getDefaultOutput() (io.WriteCloser, error) {
	path := os.Getenv("WORKFLOW_OUTPUT")
	if len(path) > 0 {
		return os.Create(path)
	}
	return os.Stdout, nil
}

func ReadInputs(reader io.Reader, data interface{}) error {
	dec := jsoniter.NewDecoder(reader)
	return dec.Decode(data)
}

func WriteOutputsTo(writer io.Writer, data interface{}) error {
	response := Response{
		Outputs: data,
	}
	enc := jsoniter.NewEncoder(writer)
	return enc.Encode(&response)
}

type Response struct {
	Outputs interface{}
}
