PROTOC ?= protoc
DEBUG_FLAGS = -gcflags="all=-N -l"

build:
	go build -o build/simulation ./cmd/simulation
	go build -o build/coordinator ./cmd/coordinator
	go build -o build/server ./cmd/server

escape:
	go build -gcflags -m github.com/decentraland/webrtc-broker/pkg/commserver
	go build -gcflags -m github.com/decentraland/webrtc-broker/pkg/coordinator

compile-protocol:
	cd pkg/protocol; ${PROTOC} --js_out=import_style=commonjs,binary:. --ts_out=. --go_out=. ./broker.proto

test: build
	go test -race $(TEST_FLAGS) \
github.com/decentraland/webrtc-broker/pkg/commserver \
github.com/decentraland/webrtc-broker/pkg/coordinator

integration:
	go test -v -race -count=1 $(TEST_FLAGS) -tags=integration github.com/decentraland/webrtc-broker/pkg/simulation

bench: build
	go test -bench=. -run="NOTHING" github.com/decentraland/webrtc-broker/pkg/commserver

cover: TEST_FLAGS=-coverprofile=coverage.out
cover: test

check-cover: cover
	go tool cover -html=coverage.out

fmt:
	gofmt -w .
	goimports -w .

version:
	git rev-parse HEAD

tidy:
	go mod tidy

todo:
	grep --include "*.go" -r TODO *

lint:
	golint ./...

.PHONY: build test integration cover
