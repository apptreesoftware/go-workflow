package step

import (
	"github.com/json-iterator/go"
	"io"
	"os"
)

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
		return
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


func WriteOutputsTo(writer io.Writer, data interface{}) error {
	enc := jsoniter.NewEncoder(writer)
	return enc.Encode(data)
}

type Response struct {
	Outputs interface{}
}
