.PHONY: run-api
run-api:
	go run infrastructure/api/main.go

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -o api infrastructure/api/main.go

.PHONY: proto
proto:
	protoc -I=./proto/ --go_out=./ ./proto/coinbit.proto