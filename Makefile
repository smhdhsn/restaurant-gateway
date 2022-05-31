.PHONY: up purge shell build_server build_all proto_user_source proto_all

# runs the script which loads the containers of the application.
up:
	@./script/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_gateway

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_gateway bash

build_server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds all the entry points of the application.
build_all: build_server

# compiles proto files related to user source.
proto_user_source:
	@protoc --go_out=internal/protos/user/source/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/user/source/ protos/user/source/*.proto

# compiles all proto files.
proto_all: proto_user_source
