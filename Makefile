.PHONY: all
all:

.PHONE: add-pre-commit
add-pre-commit:
	lefthook add pre-commit

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
