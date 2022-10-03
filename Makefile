
.PHONY: proto generate
gen:
	protoc --proto_path=proto proto/*.proto --go_out=api/v1/nginx --go-grpc_out=api/v1/nginx

clean:
	rm -rf server/pb/
	rm -rf client/pb/

.PHONY: run server
server:
	go run server/main.go redis

.PHONY: run client
client:
	go run client/main.go

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
		 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
											sudo apt install protobuf
	sudo apt install clang-format
	sudo apt install grpcurl
	export GO_PATH=~/go 
	export PATH=$PATH:/$GO_PATH/bin
	go mod vendor
path:
	export GO_PATH=~/go 
	export PATH=$PATH:/$GO_PATH/bin

test:
	rm -rf tmp && mkdir tmp
	go test -cover -race serializer/*.go

run:
	go run . ./.env

mod:
	go mod tidy
	go mod vendor