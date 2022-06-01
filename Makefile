.PHONY: up purge shell build_server build_all proto_user_source proto_edible_inventory proto_edible_recipe proto_edible_menu proto_order_submission proto_all

# runs the script which loads the containers of the application.
up:
	@./script/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_gateway

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_gateway bash

# builds server's http entry point.
build_server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds application's command entry points.
build_command:
	@go build -o $(BIN_DIR)/ ./cmd/command

# builds all the entry points of the application.
build_all: build_server build_command

# compiles proto files related to user source.
proto_user_source:
	@protoc --go_out=internal/protos/user/source/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/user/source/ protos/user/source/*.proto

# compiles proto files related to edible inventory.
proto_edible_inventory:
	@protoc --go_out=internal/protos/edible/inventory/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/inventory/ protos/edible/inventory/*.proto

# compiles proto files related to edible recipe.
proto_edible_recipe:
	@protoc --go_out=internal/protos/edible/recipe/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/recipe/ protos/edible/recipe/*.proto

# compiles proto files related to edible menu.
proto_edible_menu:
	@protoc --go_out=internal/protos/edible/menu/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/menu/ protos/edible/menu/*.proto

# compiles proto files related to order submission.
proto_order_submission:
	@protoc --go_out=internal/protos/order/submission/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/order/submission/ protos/order/submission/*.proto

# compiles all proto files.
proto_all: proto_user_source proto_edible_inventory proto_edible_recipe proto_edible_menu proto_order_submission
