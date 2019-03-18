package step

import (
	"flag"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/core"
	"google.golang.org/grpc"
	"net"
)

func Run() {
	srv := false
	port := 4000
	flag.BoolVar(&srv, "serve", false, "Run this step package as a service for development")
	flag.IntVar(&port, "port", 4000, "The port to server this package on")
	flag.Parse()

	if !srv {
		out, err := runStep(GetEnvironment(), &StdInput{})
		if err != nil {
			ReportError(err)
		}
		if out != nil {
			SetOutput(out)
		}
		return
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	core.RegisterStepHostServer(grpcServer, &Host{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("Staring package server on port %d", port))
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func runStep(env *core.Environment, ctx Context) (interface{}, error) {
	stepId := StepNameAndVersion(env)

	s := GetStep(stepId)
	if s == nil {
		return nil, fmt.Errorf("step not found in package: %s", stepId)
	}
	return s.Execute(ctx)
}
