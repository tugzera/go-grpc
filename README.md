# gRPC - Golang

# 1 - compile proto files

protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

# 2 - test client on cli

https://github.com/ktr0731/evans#not-recommended-go-get