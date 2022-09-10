.PHONY: run-api
run-api:
	go run infrastructure/api/main.go

.PHONY: run-processor
run-processor:
	go run infrastructure/processor/main.go

.PHONY: build-api
build-api:
	GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -o api infrastructure/api/main.go

.PHONY: build-processor
build-processor:
	GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -o processor infrastructure/processor/main.go

.PHONY: proto
proto:
	protoc -I=./proto/ --go_out=./ ./proto/coinbit.proto

.PHONY: test
test:
	go clean -testcache && go test -count=1 ./domain