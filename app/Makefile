BIN=$$(pwd)/bin
GOBIN=$(if $$(go env GOBIN),$$(go env GOPATH)/bin,$$(go env GOBIN))

setup:
	if [ ! -e $(BIN)/golangci-lint ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN) v1.42.0; fi
	if [ ! -e $(BIN)/air ]; then curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(BIN) v1.27.3; fi
	if [ ! -e $(GOBIN)/mockgen ]; then go install github.com/golang/mock/mockgen@v1.6.0; fi

run:
	$(BIN)/air -c .air.toml

lint:
	$(BIN)/golangci-lint run --fix --allow-parallel-runners ./...

test:
	go test -cover -parallel 5 -count 1 ./...

gen:
	go generate ./...
