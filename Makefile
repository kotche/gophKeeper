#export PATH=$PATH:$(go env GOPATH)/bin
proto:
	protoc --go_out=:internal/pb --go_opt=paths=import --go-grpc_out=:internal/pb --go-grpc_opt=paths=import internal/proto/*.proto

