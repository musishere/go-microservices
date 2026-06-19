swagger:
	PATH="$$(go env GOPATH)/bin:$$PATH" GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

protos:
	PATH="/usr/local/bin:$$(go env GOPATH)/bin:$$PATH" protoc -I Grpc/protos Grpc/protos/currency.proto \
		--go_out=Grpc/protos --go_opt=paths=source_relative \
		--go-grpc_out=Grpc/protos --go-grpc_opt=paths=source_relative