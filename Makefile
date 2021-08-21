.PHONY: all
all:

.PHONY: deps
deps:
	go mod tidy && go mod verify

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -v ./...
