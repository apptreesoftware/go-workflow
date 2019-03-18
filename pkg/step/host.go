package step

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/core"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Host struct {
}

func (Host) GetPackageInfo(context.Context, *core.EmptyMessage) (*core.Package, error) {
	b, err := ioutil.ReadFile("package.yaml")
	if err != nil {
		return nil, fmt.Errorf("Unable to read package.yaml.")
	}
	pkg := core.Package{}
	err = yaml.Unmarshal(b, &pkg)
	return &pkg, err
}

func (Host) RunStep(ctx context.Context, req *core.RunStepRequest) (*core.StepOutput, error) {
	out, err := runStep(req.Environment, &ByteInput{bytes: req.Input})
	if err != nil {
		return nil, err
	}

	var outBytes []byte
	if out != nil {
		b, err := jsoniter.Marshal(out)
		if err != nil {
			return nil, err
		}
		outBytes = b
	}
	return &core.StepOutput{
		Output: outBytes,
	}, nil
}
