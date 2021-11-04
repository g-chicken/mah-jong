setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.42.0
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.27.3
	go install github.com/golang/mock/mockgen@v1.6.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	GO111MODULE=on GOBIN=$$(go env GOPATH)/bin go install github.com/uber/prototool/cmd/prototool@dev

run:
	air -c .air.toml

lint:
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners ./...

test:
	go test -cover -parallel 5 -count 1 ./...

gen:
	go generate ./...

proto-gen:
	protoc -I=../proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./proto --go-grpc_opt=paths=source_relative $$(find ../proto/app/services/ -name "*.proto")
