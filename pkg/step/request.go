package step

import (
	"flag"
	"github.com/json-iterator/go"
	"io"
	"os"
	"strings"
)

func BindInputs(data interface{}) {
	var input string
	flag.StringVar(&input, "input", "", "Pass input on startup rather than stdin")
	flag.Parse()
	var reader io.Reader
	if len(input) > 0 {
		reader = strings.NewReader(input)
	} else {
		reader = os.Stdin
	}
	err := ReadInputs(reader, data)
	if err != nil {
		ReportError(err)
	}
}

func ReportError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
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
	enc := jsoniter.NewEncoder(writer)
	return enc.Encode(data)
}

type Response struct {
	Outputs interface{}
}
