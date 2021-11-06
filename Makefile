setup:
	mkdir tmp_migrate
	curl -sSOL https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz
	mv *tar.gz tmp_migrate && cd tmp_migrate && tar -xvzf *.tar.gz
	sudo mv tmp_migrate/migrate /usr/local/bin
	rm -r tmp_migrate
	curl -sSLO "https://github.com/bufbuild/buf/releases/download/v1.0.0-rc6/buf-$$(uname -s)-$$(uname -m)"
	chmod 755 buf-$$(uname -s)-$$(uname -m)
	sudo mv ./buf-$$(uname -s)-$$(uname -m)  /usr/local/bin/buf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	GO111MODULE=on GOBIN=$$(go env GOPATH)/bin go install github.com/uber/prototool/cmd/prototool@dev
	mkdir tmp_protoc
	curl -sSLO "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-$$(uname -s)-$$(uname -m).zip"
	mv *.zip tmp_protoc && cd tmp_protoc &&	unzip protoc-3.19.1-$$(uname -s)-$$(uname -m).zip
	sudo mv ./tmp_protoc/include /usr/local/bin
	sudo mv ./tmp_protoc/bin/protoc /usr/local/bin
	rm -rf tmp_protoc

migrate-up:
	migrate -source file://migrations -database mysql://localhost:3306 up

run:
	./bin/up.sh

lint:
	buf lint

gen:
	protoc -I=./proto --go_out=./app/proto --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./app/proto --go-grpc_opt=paths=source_relative $$(find ./proto/app/services/ -name "*.proto")
