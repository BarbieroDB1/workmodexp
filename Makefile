

build-proto ::
	@echo "Building proto files"
	cd ./proto && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

test-all ::
	@echo "Testing all modules"
	go test github.com/barbierodb1/workmodexp/...