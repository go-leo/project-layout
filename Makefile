# note: call scripts from /scripts

GIT_VERSION=$(shell git tag | grep v | sort -r --version-sort | head -n1)
PROJECT_PATH:=$(shell pwd)

.PHONY: tools
tools:
	@sh scripts/tools.sh

.PHONY: wire_gen
wire_gen:
	@sh scripts/wire_gen.sh

.PHONY: go_gen
go_gen:
	@sh scripts/go_gen.sh

.PHONY: protoc_gen
protoc_gen:
	@sh scripts/protoc_gen.sh

.PHONY: conf_gen
conf_gen:
	@sh scripts/conf_gen.sh

.PHONY: gen
gen:
	@echo "--- generate code start ---"
	@$(MAKE) go_gen
	@$(MAKE) protoc_gen
	@$(MAKE) wire_gen
	@echo "--- generate code end ---"

.PHONY: format
format:
	@sh scripts/shell/format.sh

.PHONY: lint
lint:
	@sh scripts/shell/lint.sh

.PHONY: test
test:
	@echo "--- go test start ---"
	go test -test.bench=".*" -count=1 -v ./...
	@echo "--- go test end ---"

.PHONY: build
build:
	@echo "--- go build start ---"
	CGO_ENABLED=1 go build -o bin/$(path) -a -ldflags "-w -s -X main.Version=$(GIT_VERSION)" -tags=jsoniter ./cmd/$(path)
	@echo "--- go build end ---"

