.PHONY: test, coverage and mocks

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	
generate-mocks:
	mockgen --source=./daemon/internal/store/service/interface.go --destination=./daemon/internal/mocks/store_mock.go  --package=mocks
	mockgen --source=./daemon/internal/vector/service/interface.go --destination=./daemon/internal/mocks/vector_mock.go --package=mocks
	mockgen --source=./daemon/internal/nativeops/interfaces.go --destination=./daemon/internal/mocks/operations_c_mock.go --package=mocks
	mockgen --source=./daemon/internal/index/service/interface.go --destination=./daemon/internal/mocks/index_mock.go --package=mocks
	

