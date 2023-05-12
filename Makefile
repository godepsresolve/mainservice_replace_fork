.PHONY: static_check
static_check:
	go fmt ./... && go vet ./... && staticcheck ./...

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: deps
deps:
	go mod tidy
