# Makefile

.PHONY: lint
lint:
	# Install golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.53.3

	# Run linter manually
	$(shell go env GOPATH)/bin/golangci-lint run --enable=gofmt --tests=false --exclude-use-default --timeout=5m0s

.PHONY: gomodtidy
gomodtidy:
	go mod tidy
	cd ./ops && go mod tidy

.PHONY: godoc
godoc:
	go install golang.org/x/tools/cmd/godoc@latest
	# http://localhost:6060/pkg/github.com/smartcontractkit/chainlink-relay/
	godoc -http=:6060

PHONY: install-protoc
install-protoc:
	script/install-protoc.sh 24.2 /
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31; go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 

.PHONY: mockery
mockery: $(mockery) ## Install mockery.
	go install github.com/vektra/mockery/v2@v2.28.1

PHONY: generate

generate: mockery install-protoc
# add our installed protoc to the head of the PATH
# maybe there is a cleaner way to do this
	 PATH=$$HOME/.local/bin:$$PATH go generate -x ./...
	  