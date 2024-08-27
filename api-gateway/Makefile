run:
	go run cmd/main.go

job_pb:
	protoc --proto_path=./pkg/pb/job --go_out=paths=source_relative:./pkg/pb/job --go-grpc_out=paths=source_relative:./pkg/pb/job ./pkg/pb/job/job.proto

auth_pb:
	protoc --proto_path=./pkg/pb/auth --go_out=paths=source_relative:./pkg/pb/auth --go-grpc_out=paths=source_relative:./pkg/pb/auth ./pkg/pb/auth/auth.proto