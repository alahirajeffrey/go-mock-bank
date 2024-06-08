build:
	@go build -o bin/go-mock-bank

run: build
	@./bin/go-mock-bank

test:
	@go test -v ./...