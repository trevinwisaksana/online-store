PROTOC ?= protoc
PROTOINCLUDE ?= googleapis
OUTPUT ?= ./gen

.PHONY: proto

proto:
	@if [ -z "$(SOURCE_DIR)" ]; then \
		echo "Error: SOURCE_DIR is required. Usage: make proto SOURCE_DIR=<path>"; \
		exit 1; \
	fi
	
	$(PROTOC) \
		--proto_path=$(PROTOINCLUDE) \
		--proto_path=. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--grpc-gateway_out=. \
		--go-grpc_opt=paths=source_relative \
		$(SOURCE_DIR)