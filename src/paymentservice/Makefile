GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install go-micro.dev/v4/cmd/protoc-gen-micro@v4

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/paymentservice.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/health.proto
	

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o paymentservice *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t paymentservice:latest .