BIN=$$(pwd)/bin
TMP=$$(pwd)/tmp
GOBIN=$(if $$(go env GOBIN),$$(go env GOPATH)/bin,$$(go env GOBIN))

setup:
	if [ ! -e $(BIN)/migrate ]; then\
		if [ ! -d $(TMP) ]; then mkdir $(TMP); fi;\
		curl -sSL https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz -o $(TMP)/migrate.tar.gz;\
		tar -xvzf $(TMP)/migrate.tar.gz -C $(TMP);\
		mv $(TMP)/migrate $(BIN);\
		rm -rf $(TMP);\
	fi
	if [ ! -e $(BIN)/buf ]; then\
		curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.0.0-rc6/buf-$$(uname -s)-$$(uname -m)" -o $(BIN)/buf;\
		chmod 755 $(BIN)/buf;\
	fi
	if [ ! -e $(GOBIN)/proto-gen-go ]; then	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1; fi
	if [ ! -e $(GOBIN)/protoc-gen-go-grpc ]; then go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0; fi
	if [ ! -e $(GOBIN)/prototool ]; then GO111MODULE=on GOBIN=$(GOBIN) go install github.com/uber/prototool/cmd/prototool@dev; fi
	if [ ! -d $(BIN)/include ]; then\
		if [ ! -d $(TMP) ]; then mkdir $(TMP); fi;\
		curl -sSL "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-$$(uname -s)-$$(uname -m).zip" -o $(TMP)/protoc.zip;\
		unzip $(TMP)/protoc.zip -d $(TMP);\
		mv $(TMP)/include $(BIN);\
		rm -rf $(TMP);\
	fi

migrate-up:
	$(BIN)/migrate -source file://migrations -database mysql://localhost:3306 up

run:
	$(BIN)/up.sh

lint:
	$(BIN)/buf lint

gen:
	protoc -I=./proto --go_out=./app/proto --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./app/proto --go-grpc_opt=paths=source_relative $$(find ./proto/app/services/ -name "*.proto")
