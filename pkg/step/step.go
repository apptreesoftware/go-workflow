package step

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

type Step interface {
	Name() string
	Version() string
	Description() string
	Execute()
	DefineIO() (map[string]*InputInfo, map[string]*OutputInfo)
}

func Id(s Step) string {
	return fmt.Sprintf("%s@%s", s.Name(), s.Version())
}

func StepNameAndVersion(environment *Environment) string {
	return fmt.Sprintf("%s@%s", environment.StepName, environment.StepVersion)
}

var stepReg = map[string]Step{}

func Register(step Step) {
	id := Id(step)
	stepReg[id] = step
}

func PrintStepDefinitions(packageName string) {
	packageInfo := Package{
		Name:  packageName,
		Lang:  "go",
		Exec:  "main",
		Steps: map[string]*PackageStep{},
	}

	for _, stp := range stepReg {
		in, out := stp.DefineIO()

		stepDesc := &PackageStep{
			Description: stp.Description(),
			Inputs:      in,
			Outputs:     out,
		}
		packageInfo.Steps[Id(stp)] = stepDesc
	}
	b, err := yaml.Marshal(packageInfo)
	if err != nil {
		panic(err)
	}
	_, _ = os.Stdout.Write(b)
}

func GetStep(id string) Step {
	return stepReg[id]
}

func DefineInput(in interface{}) map[string]*InputInfo {
	t := reflect.TypeOf(in)
	info := map[string]*InputInfo{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("io")
		if len(tag) == 0 {
			info[field.Name] = &InputInfo{
				Required:    true,
				Description: "",
			}
		}
		info[field.Name] = &InputInfo{Required: true}
	}
	return info
}

//func inputFromTag(tag string, field reflect.StructField) (string,*InputInfo) {
//	if len(tag) == 0 {
//		return field.Name, &InputInfo{
//			Required:             true,
//			Description:          "",
//		}
//	}
//	components := strings.Split(tag, ",")
//	name := components[0]
//	opts := components[1:]
//
//}

func DefineOutput(out interface{}) map[string]*OutputInfo {
	t := reflect.TypeOf(out)
	info := map[string]*OutputInfo{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("io")
		if len(tag) == 0 {
			info[field.Name] = &OutputInfo{
				Description: "",
			}
		}
		info[field.Name] = &OutputInfo{}
	}
	return info
}
