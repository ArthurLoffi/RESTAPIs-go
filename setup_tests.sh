#!/bin/bash
# setup_tests.sh — instala dependências de teste e roda todos os testes

set -e

echo "Adicionando gorm sqlite"
go get gorm.io/driver/sqlite

echo "Adicionando testify"
go get github.com/stretchr/testify@latest

echo "Dependências instaladas."

echo ""
echo "Rodando todos os testes"
go test ./... -v -count=1

echo ""
echo "Cobertura de testes:"
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out