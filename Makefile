PROTO_DIR := proto
PROTO_SRC := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .

.PHONY: generate-proto
build_proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_SRC)

build_proto_0:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/cashier.proto

build_proto_cashier:
	protoc --go_out=. --go-grpc_out=. ./proto/cashier.proto

build_proto_balance:
	protoc --go_out=. --go-grpc_out=. ./proto/balance.proto