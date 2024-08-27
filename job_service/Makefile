run:
	go run cmd/main.go

pb:
	protoc --proto_path=./pkg/pb/job --go_out=paths=source_relative:./pkg/pb/job --go-grpc_out=paths=source_relative:./pkg/pb/job ./pkg/pb/job/job.proto