all: build
build-node: |
	grpc_tools_node_protoc -I proto --js_out=import_style=commonjs,binary:node/gen --grpc_out=node/gen --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` common.proto step.proto
build: |
	#Generate common models
	protoc -I proto --go_out=paths=source_relative:pkg/core common.proto

	#Generate Step API
	#API used by the step to make calls to the engine that is responsible for running it
#	protoc -I proto --go_out=paths=source_relative,plugins=grpc:pkg/core step.proto


	protoc -I proto --plugin=protoc-gen-grpc=${HOME}/.nuget/packages/grpc.tools/1.17.1/tools/macosx_x64/grpc_csharp_plugin --grpc_out csharp/StepCore/gen --csharp_out=csharp/StepCore/gen step.proto common.proto
	
	protoc -I proto --js_out=library=workflow,binary:node/gen  step.proto common.proto

	##GRPC APIs
	protoc -I proto --go_out=paths=source_relative,plugins=grpc:pkg/core remote_engine.proto step.proto
#
#	#Generate Engine API
#	#API used by the remote engine to communicate back to the host engine
	##Twirp APIs
	protoc -I proto --go_out=paths=source_relative:pkg/core --twirp_out=paths=source_relative:pkg/core engine_api.proto workflow_api.proto platform_api.proto step_library.proto step_package_api.proto
	protoc -I proto --twirp_dart_out=paths=source_relative:dart/lib workflow_api.proto step_library.proto step_package_api.proto
	dartfmt -w dart/lib/.



	protoc-go-inject-tag -input=pkg/core/workflow_api.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/platform_api.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/common.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/engine_api.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/step.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/step_library.pb.go -XXX_skip=yaml,xml,bson
	protoc-go-inject-tag -input=pkg/core/step_package_api.pb.go -XXX_skip=yaml,xml,bson
#
#	#Generate Client API
#	#API used by clients (cli, web) to request data from the workflow system
#	protoc -I proto --go_out=paths=source_relative:pkg/core --twirp_out=paths=source_relative:pkg/core api.proto
#	protoc-go-inject-tag -input=pkg/core/api.pb.go -XXX_skip=yaml,xml,bson
