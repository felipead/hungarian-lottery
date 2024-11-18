.PHONY: build
build:
	@go build -o hungarian-lottery cmd/main.go

.PHONY: clean
clean:
	@rm -rf ./hungarian-lottery

.PHONY: ci
ci:
	@make test

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test
test:
	@go test -v ./...
