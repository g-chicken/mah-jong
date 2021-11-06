setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.42.0
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.27.3
	go install github.com/golang/mock/mockgen@v1.6.0

run:
	air -c .air.toml

lint:
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners ./...

test:
	go test -cover -parallel 5 -count 1 ./...

gen:
	go generate ./...
