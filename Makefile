.PHONY: all
all:

.PHONY: add-pre-commit
add-pre-commit:
	lefthook add pre-commit

.PHONY: deps
deps:
	go mod tidy && go mod vendor && go mod verify

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -v ./...

.PHONY: update-deps
update-deps:
	go get -t -u ./...
