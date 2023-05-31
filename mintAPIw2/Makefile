CONTRACT_PATH ?= ./contracts/Shows.sol
BUILD_PATH ?= ./contracts/Show/abi/

ABI_FILE ?= ./contracts/Show/abi/Shows_sol_Shows.abi
PACKAGE_NAME ?= sticket
OUTPUT_PATH ?= ./internal/infra/quicknode/contract/sticket/ShowsTicket.go

BINARY_NAME ?= api
DIRECTORY_PATH ?= .

.PHONY: run api generate_domain solidity_conversor contract_conversor

run:
	go run main.go

api:
	go build -o $(BINARY_NAME) -ldflags "-s -w" $(DIRECTORY_PATH)

generate_domain:
	protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/core/domain/domain.proto

solidity_conversor:
	solc --optimize --abi $(CONTRACT_PATH) -o $(BUILD_PATH)

contract_conversor:
	abigen --abi=$(ABI_FILE) --pkg=$(PACKAGE_NAME) --out=$(OUTPUT_PATH)

default: api