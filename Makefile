# Desenvolvimento
dev:
	@air

# Production
run:
	@echo "Running server..."
	@go run main.go	

# Build
build:
	@go build -o ./bin ./main.go

# Testes
test:
	@go test -v ./...

test-coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Limpeza
clean:
	@rm -rf tmp/
	@rm -f build-errors.log
	@rm -f coverage.out

vite: 
	bun run vite ./web/public --port 3000	

build-vite:
	bun run vite build ./web/public	--config ./web/vite.config.js

tailwindcss:
	bun run tailwindcss --config ./web/tailwind.config.js -i ./web/public/assets/styles.css -o ./dist/css/styles.css	

help:
	@echo "Available commands:"
	@echo "  make run          - Run the application in production mode with hot reload"
	@echo "  make dev          - Run the application in development mode with hot reload"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean temporary files"
	@echo "  make vite         - Run vite server"

.PHONY: run dev build test test-coverage clean vite tailwindcss help 
