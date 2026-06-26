module github.com/musishere/working-package

go 1.26.1

require google.golang.org/grpc v1.81.1 // indirect

require (
	github.com/gabriel-vasile/mimetype v1.4.13 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.3
	github.com/gorilla/mux v1.8.1
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/musishere/grpc v0.0.0
	golang.org/x/crypto v0.52.0 // indirect
	golang.org/x/net v0.54.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/musishere/grpc => ../Grpc
