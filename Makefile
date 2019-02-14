all: build
build: |
	protoc -I proto --go_out=paths=source_relative,plugins=grpc:pkg/step step.proto
	protoc --plugin=protoc-gen-grpc=${HOME}/.nuget/packages/grpc.tools/1.17.1/tools/macosx_x64/grpc_csharp_plugin --grpc_out csharp/StepCore/gen --csharp_out=csharp/StepCore/gen proto/step.proto
	protoc-go-inject-tag -input=pkg/step/step.pb.go -XXX_skip=yaml,xml,bson