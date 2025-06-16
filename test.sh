#!/bin/bash

set -e

GO_VERSION="1.24.3"
MOCKGEN_VERSION="v1.6.0"

# Instala o Go (caso não exista)
if ! command -v go &> /dev/null; then
  echo "Instalando Go ${GO_VERSION}..."
  wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
  sudo rm -rf /usr/local/go
  sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
  export PATH=$PATH:/usr/local/go/bin
fi

# Exporta o PATH
export PATH=$PATH:/usr/local/go/bin

# Instala o mockgen (se ainda não existir)
if ! command -v mockgen &> /dev/null; then
  echo "Instalando mockgen ${MOCKGEN_VERSION}..."
  go install github.com/golang/mock/mockgen@${MOCKGEN_VERSION}
fi

# Confirma as versões instaladas
go version
mockgen --version

echo "Ambiente configurado com sucesso 🚀"
