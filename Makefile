APP_MODE ?= local

# runs the script which loads the containers of the application.
up:
	@./scripts/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_gateway_app

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_gateway_app bash

# builds server's http entry point.
build-server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds application's command entry points.
build-command:
	@go build -o $(BIN_DIR)/ ./cmd/command

# builds all the entry points of the application.
build-all: build-server build-command

# compiles proto files related to edible inventory.
proto-edible-inventory:
	@protoc --go_out=internal/protos/edible/inventory/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/inventory/ protos/edible/inventory/*.proto

# compiles proto files related to edible recipe.
proto-edible-recipe:
	@protoc --go_out=internal/protos/edible/recipe/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/recipe/ protos/edible/recipe/*.proto

# compiles proto files related to edible menu.
proto-edible-menu:
	@protoc --go_out=internal/protos/edible/menu/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/menu/ protos/edible/menu/*.proto

# compiles proto files related to order submission.
proto-order-submission:
	@protoc --go_out=internal/protos/order/submission/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/order/submission/ protos/order/submission/*.proto

# compiles all proto files.
proto-all:  proto-edible-inventory proto-edible-recipe proto-edible-menu proto-order-submission

.PHONY: up purge shell build-server build-command build-all  proto-edible-inventory proto-edible-recipe proto-edible-menu proto-order-submission proto-all