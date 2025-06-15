# Variáveis
MOCKGEN=mockgen
MOCK_DIR=internal/mock
CONTRACT_DIR=internal/domain/contract

# Targets

.PHONY: all mocks test install-mockgen tidy

all: tidy mocks test

install-mockgen:
	go install github.com/golang/mock/mockgen@latest

tidy:
	go mod tidy

mocks:
	@echo "Gerando mocks..."
	$(MOCKGEN) -source=$(CONTRACT_DIR)/service.go -destination=$(MOCK_DIR)/service_mock.go -package=mock
	$(MOCKGEN) -source=$(CONTRACT_DIR)/repo.go -destination=$(MOCK_DIR)/repo_mock.go -package=mock

test:
	go test ./... -v
