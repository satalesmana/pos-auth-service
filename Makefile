proto:
	@cd files/proto && protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. *.proto

server:
	go run cmd/main.go