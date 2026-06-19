swagger:
	PATH="$$(go env GOPATH)/bin:$$PATH" GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models