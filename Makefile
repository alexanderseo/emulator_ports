GRPC_PKG_DIR:=./pkg/api/grpc
GRPC_API_PROTO_PATH=./api/grpc

grpc-rm:
	rm -rf $(GRPC_PKG_DIR)/*.go


install-tools:
	./scripts/install_protoc.sh
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


grpc-gen: grpc-rm
	mkdir -p $(GRPC_PKG_DIR)
	./third_party/protoc/bin/protoc \
		-I ./third_party/protoc/include/google/protobuf \
        -I ./third_party/protoc/include/google/protobuf/compiler \
		-I ./api/grpc \
        	--go_out=./pkg/api/grpc \
        	--go_opt=paths=source_relative \
        	--go-grpc_out=./pkg/api/grpc \
        	--go-grpc_opt=paths=source_relative \
        	./api/grpc/*.proto


install_linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1
	golangci-lint --version

run_linter:
	golangci-lint run -v
