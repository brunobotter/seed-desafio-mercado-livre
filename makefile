MOCKGEN := $(shell go env GOPATH)/bin/mockgen

.PHONY: all tidy mocks test clean

all: tidy mocks test

tidy:
	go mod tidy

mocks:
	@echo "Gerando mocks..."
	$(MOCKGEN) -source=internal/domain/contract/service.go -destination=internal/mock/service_mock.go -package=mock
	$(MOCKGEN) -source=internal/domain/contract/repo.go -destination=internal/mock/repo_mock.go -package=mock

test:
	go test ./... -v

clean:
	rm -rf internal/mock/*.go
