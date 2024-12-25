# Define variables
PROTO_DIR = api
PROTO_FILE = $(PROTO_DIR)/expense.proto
GO_OUT_DIR = .

# Protobuf generation commands with source-relative paths
generate:
	protoc --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)

# Clean generated files
clean:
	rm -f $(GO_OUT_DIR)/*.pb.go

# Default target
all: generate

# Phony targets to avoid conflicts with files of the same name
.PHONY: generate clean all
