# BluePrint ROA Golang Grpc Gateway
# user-service
# Software Requirement

- protobuf
- grpc
- make

# Lib

- go install github.com/favadi/protoc-go-inject-tag@v1.3.0
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
- go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Usage

- GRPC : localhost:3040
- HTTP : localhost:3041
- Swagger : localhost:3041/swagger-ui
