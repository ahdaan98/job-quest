run:
	go run cmd/main.go

pb:
	protoc --proto_path=./pkg/pb/auth --go_out=paths=source_relative:./pkg/pb/auth --go-grpc_out=paths=source_relative:./pkg/pb/auth ./pkg/pb/auth/auth.proto

pb:
	protoc --proto_path=./pkg/pb/chat --go_out=paths=source_relative:./pkg/pb/chat --go-grpc_out=paths=source_relative:./pkg/pb/chat ./pkg/pb/chat/chat.proto