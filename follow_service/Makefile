run:
	go run cmd/main.go

pb:
	protoc --proto_path=./pkg/pb/connection --go_out=paths=source_relative:./pkg/pb/connection --go-grpc_out=paths=source_relative:./pkg/pb/connection ./pkg/pb/connection/connection.proto