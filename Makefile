CPUS ?= $(shell (nproc --all || sysctl -n hw.ncpu) 2>/dev/null || echo 1)
MAKEFLAGS += --jobs=$(CPUS)
PROTO_FILES := $(shell find protobufs -mindepth 2 -maxdepth 2 -type f -name "*.proto"  -path "**/steam/**")

.PHONY: help
help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: fakes
fakes: ## Generate fakes for unit testing
	@go install go.uber.org/mock/mockgen@latest
	@go generate ./...

.PHONY: build
build: ## Build the application
	@go build ./...

.PHONY: test
test: ## Run unit tests
	@go test ./...

.PHONY: lint
lint: ## Lint files
	@golangci-lint run ./...

.PHONY: coverage
coverage: coverage-html coverage-xml ## Generate and report coverage
	@printf "Coverage: %0.1f%%\n" $(shell echo "$(shell grep -oPm1 '(?<=line-rate=")\d+\.\d+' coverage/coverage.xml) * 100" | bc)
	@printf "Lines   : %d/%d\n" $(shell grep -oPm1 '(?<=lines-covered=")[\d]+' coverage/coverage.xml) $(shell grep -oP '(?<=lines-valid=")[\d]+' coverage/coverage.xml)

.PHONY: coverage-profile
coverage-profile:
	@mkdir -p coverage
	@go test -v -cover -covermode=count -coverprofile=coverage/profile.cov ./...

.PHONY: coverage-html
coverage-html: coverage-profile
	@go tool cover -html=coverage/profile.cov -o coverage/coverage.html

.PHONY: coverage-xml
coverage-xml: coverage-profile
	@go install github.com/axw/gocov/gocov@latest
	@go install github.com/AlekSi/gocov-xml@latest
	@gocov convert coverage/profile.cov | gocov-xml > coverage/coverage.xml

.PHONY: language
language: ## Generate language files from protobufs
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@rm -rf language
	@$(MAKE) $(PROTO_FILES)

define compile_proto
	@mkdir -p language/$(shell basename $(dir $(1)))
	@-protoc \
		--proto_path=$(dir $(1)) \
		--experimental_allow_proto3_optional \
		$(foreach proto,$(wildcard $(dir $(1))*.proto),--go_opt=M$(shell basename $(proto))=./$(shell basename $(dir $(1)))) \
		--go_out=language \
		$(1)
endef
define compile_proto_target
.PHONY: $(1)
$(1):
	$$(call compile_proto,$(1))
endef
$(foreach proto,$(PROTO_FILES),$(eval $(call compile_proto_target,$(proto))))
