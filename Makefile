all: build
build: |
	./scripts/build_steps.sh
	protoc --go_out=paths=source_relative,plugins=grpc:. pkg/cache/cache.proto
	protoc --plugin=protoc-gen-grpc=${HOME}/.nuget/packages/grpc.tools/1.17.1/tools/macosx_x64/grpc_csharp_plugin --csharp_out=lang/csharp pkg/cache/cache.proto
clean: |
	./scripts/clean_steps.sh