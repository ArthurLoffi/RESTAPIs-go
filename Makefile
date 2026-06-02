.PHONY: docs

build:
	go mod tidy
	go build ./cmd/api/.

run:
	./api

docs:
	swag init -g ./cmd/api/main.go --output ./docs

# Ao instalar o repo, rodar este comando
init:
	go mod download

test:
	./setup_tests.sh