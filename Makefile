.PHONY: build
build:
	@go build -o lottery cmd/main.go

.PHONY: clean
clean:
	@rm -rf ./lottery

.PHONY: ci
ci:
	@make lint
	@make test

.PHONY: lint
lint:
	@echo all good for now

.PHONY: test
test:
	go test -v ./...
