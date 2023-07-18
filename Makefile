client:
	@go run ./cmd/client

server:
	@go run ./cmd/server

test:
	go test -v ./...

### -- Protocol Buffers -- ###
GOPATH=/go
PROTO_PATH=pkg/proto/rpc
REPO=/workspace/protocol
GO_PROTOCOL=pkg/proto
TWIRP_VERSION=@v8.1.3+incompatible
PROTOC_GEN_VERSION=@v1.0.2
PROTOC_GEN_TWIRP=@v7.2.0
PB_REL="https://github.com/protocolbuffers/protobuf/releases"

echo:
	@echo ${PWD}/${TS_PROTOCOL}
generate_go: 
	@protoc \
	-I ${GOPATH}/pkg/mod \
	-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate${PROTOC_GEN_VERSION} \
	--proto_path=${PROTO_PATH} \
	${PROTO_PATH}/*.proto \
	--go_out=paths=source_relative:${GO_PROTOCOL} \
	--twirp_out=paths=source_relative:${GO_PROTOCOL} \
	--validate_out="lang=go,paths=source_relative:${GO_PROTOCOL}"

generate: generate_go 

install:  manual_install 
	# sudo apt install -y protobuf-compiler 
	# protoc --version
	# sudo apt-get -y install protoc-gen-go

manual_install: 
	curl -LO ${PB_REL}/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
	unzip protoc-3.15.8-linux-x86_64.zip -d ${HOME}/.local
	export PATH="${PATH}:${HOME}/.local/bin"
	rm ./protoc-3.15.8-linux-x86_64.zip
	protoc --version
	go install github.com/envoyproxy/protoc-gen-validate
	go install github.com/twitchtv/twirp/protoc-gen-twirp
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	export PATH="${PATH}:${GOPATH}/bin"
	
version:
	protoc --version
	protoc-gen-go --version
	protoc-gen-validate --version

clear:
	rm -rf ${GO_PROTOCOL}/*
	mkdir -p ${GO_PROTOCOL}