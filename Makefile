#export PATH=$PATH:$(go env GOPATH)/bin
proto:
	protoc --go_out=:internal/pb --go_opt=paths=import --go-grpc_out=:internal/pb --go-grpc_opt=paths=import internal/proto/*.proto

gen:
	mockgen -source=internal/server/storage/repository.go -destination=internal/mocks/server/mock_storage.go
	mockgen -source=internal/client/updater/updater.go -destination=internal/mocks/client/mock_updater.go


