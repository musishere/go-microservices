generate:
	protoc --go_out=paths=source_relative:./proto/gen --go-grpc_out=paths=source_relative:./proto/gen --proto_path=./proto ./proto/main.proto
	protoc --go_out=paths=source_relative:./proto/gen --go-grpc_out=paths=source_relative:./proto/gen --proto_path=./proto/user ./proto/user/user.proto
	protoc --go_out=paths=source_relative:./proto/gen --go-grpc_out=paths=source_relative:./proto/gen --proto_path=./proto/order --proto_path=./proto ./proto/order/order.proto