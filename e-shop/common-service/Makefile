# Define variables
PROTO_DIRS = expense users
GO_OUT_DIR = .

# Find all .proto files in specified directories
PROTO_FILES = $(wildcard $(patsubst %,%/*.proto,$(PROTO_DIRS)))

# Protobuf generation commands with source-relative paths
generate:
	$(foreach PROTO_FILE, $(PROTO_FILES), \
		protoc --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
			--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
			$(PROTO_FILE) && echo Generated $(PROTO_FILE);)

# Clean generated files
clean:
	rm -f $(GO_OUT_DIR)/*.pb.go

# Default target
all: generate

# Phony targets to avoid conflicts with files of the same name
.PHONY: generate clean all
