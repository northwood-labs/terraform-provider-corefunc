#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

SHELL:=bash
BINARY_NAME=terraform-provider-corefunc
BINARY_VERSION=$(shell cat ./VERSION | tr -d '\n')
GOBIN=$(shell ./find-go-bin.sh)
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

#-------------------------------------------------------------------------------
# Global stuff.

GO=$(shell which go)
HOMEBREW_PACKAGES=bash coreutils findutils go jq nodejs pre-commit python@3.11

# Determine the operating system and CPU arch.
OS=$(shell uname -o | tr '[:upper:]' '[:lower:]')

# Determine which version of `echo` to use. Use version from coreutils if available.
ECHOCHECK_HOMEBREW_AMD64 := $(shell command -v /usr/local/opt/coreutils/libexec/gnubin/echo 2> /dev/null)
ECHOCHECK_HOMEBREW_ARM64 := $(shell command -v /opt/homebrew/opt/coreutils/libexec/gnubin/echo 2> /dev/null)

ifdef ECHOCHECK_HOMEBREW_AMD64
	ECHO=/usr/local/opt/coreutils/libexec/gnubin/echo -e
else ifdef ECHOCHECK_HOMEBREW_ARM64
	ECHO=/opt/homebrew/opt/coreutils/libexec/gnubin/echo -e
else ifeq ($(findstring linux,$(OS)), linux)
	ECHO=echo -e
else
	ECHO=echo
endif

#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all: help

.PHONY: help
## help: [help]* Prints this help message.
help:
	@ $(ECHO) "Usage:"
	@ $(ECHO) ""
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | \
		while IFS= read -r line; do \
			if [[ "$$line" == *"]*"* ]]; then \
				$(ECHO) "\033[1;33m$$line\033[0m"; \
			else \
				$(ECHO) "$$line"; \
			fi; \
		done

#-------------------------------------------------------------------------------
# Installation

.PHONY: install-tools-go
## install-tools-go: [tools]* Install/upgrade the required Go packages.
install-tools-go:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Installing Go packages...\033[0m"
	go install github.com/google/osv-scanner/cmd/osv-scanner@v1
	go install golang.org/x/perf/cmd/benchstat@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: install-tools-mac
## install-tools-mac: [tools]* Install/upgrade the required tools for macOS, including Go packages.
install-tools-mac: install-tools-go
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Installing required packages for macOS (Homebrew)...\033[0m"
	brew update && brew install $(HOMEBREW_PACKAGES) && brew upgrade $(HOMEBREW_PACKAGES)
	curl -sSLf https://raw.githubusercontent.com/mtdowling/chag/master/install.sh | bash

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33mTo update to the latest versions, run:\033[0m"
	@ $(ECHO) "\033[1;33m    brew update && brew upgrade\033[0m"
	@ $(ECHO) " "

.PHONY: install-hooks
## install-hooks: [tools]* Install/upgrade the Git hooks used for ensuring consistency.
install-hooks:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Installing Git hooks...\033[0m"
	pre-commit install

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33mLearn more about `pre-commit` at:\033[0m"
	@ $(ECHO) "\033[1;33m    https://pre-commit.com/\033[0m"
	@ $(ECHO) " "

#-------------------------------------------------------------------------------
# Compile

.PHONY: tidy
## tidy: [build] Updates go.mod and downloads dependencies.
tidy:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Tidy and download the Go dependencies...\033[0m"
	$(GO) mod tidy -go=1.21 -v
	$(GO) mod download -x
	$(GO) get -v ./...

.PHONY: build
## build: [build]* Builds and installs the Terraform provider.
build: tidy
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Building and installing the provider...\033[0m"
	$(GO) install -v .
	@ $(ECHO) " "
	@ ls -lahF $(GOBIN)/$(BINARY_NAME)

#-------------------------------------------------------------------------------
# Clean

.PHONY: clean-go
## clean-go: [clean] Clean Go's module cache.
clean-go:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Cleaning Go cache...\033[0m"
	$(GO) clean -i -r -x -testcache -modcache -cache

.PHONY: clean-tests
## clean-tests: [clean] Cleans all files/folders inside the examples directory which begin with a ".".
clean-tests:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Cleaning artifacts from tests...\033[0m"
	- find . -type d -name ".terraform" | xargs rm -Rf
	- find . -type d -name "terratest-*" | xargs rm -Rf
	- find . -type f -name "terraform.tfstate*" | xargs rm -Rf
	- find ./examples -type d -name "\.*" | xargs rm -Rf

.PHONY: clean-bench
## clean-bench: [clean] Cleans all benchmarking-related files.
clean-bench:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Cleaning artifacts from benchmarks...\033[0m"
	- find . -type f -name "__*.out" | xargs rm -Rf
	- find . -type f -name "*.test" | xargs rm -Rf

.PHONY: clean-tf
## clean-tf: [clean] Clean Terraform leftovers.
clean-tf:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Cleaning Terraform artifacts...\033[0m"
	find . -type d -name "terraform.d" | xargs -I% rm -Rfv "%"
	find . -type d -name ".terraform" | xargs -I% rm -Rfv "%"
	find . -type f -name ".terraform.lock.hcl" | xargs -I% rm -fv "%"

.PHONY: clean
## clean: [clean]* Runs ALL cleaning tasks (except the Go cache).
clean: clean-bench clean-tf clean-tests

#-------------------------------------------------------------------------------
# Documentation

.PHONY: docs
## docs: [docs]* Runs primary documentation tasks.
docs: docs-provider docs-cli

.PHONY: docs-provider
## docs-provider: [docs] Generate Terraform Registry documentation.
docs-provider:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Generating Terraform provider documentation...\033[0m"
	$(GO) generate -v ./...

.PHONY: docs-cli
## docs-cli: [docs] Preview the Go library documentation on the CLI.
docs-cli:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Displaying Go CLI documentation...\033[0m"
	$(GO) doc -C corefunc/ -all

.PHONY: docs-serve
## docs-serve: [docs] Preview the Go library documentation as displayed on pkg.go.dev.
docs-serve:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Displaying Go HTTP documentation...\033[0m"
	open http://localhost:6060/pkg/github.com/northwood-labs/terraform-provider-corefunc/corefunc/
	godoc -index -links

#-------------------------------------------------------------------------------
# Linting

.PHONY: vuln
## vuln: [lint]* Runs `govulncheck` (vulnerability scanning) against all Golang (*.go) code.
vuln:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running govulncheck (https://go.dev/blog/vuln)...\033[0m"
	govulncheck ./...

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running govulncheck -test (https://go.dev/blog/vuln)...\033[0m"
	govulncheck -test ./...

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running osv-scanner (https://osv.dev)...\033[0m"
	osv-scanner -r .

.PHONY: lint
## lint: [lint]* Runs ALL linting/validation tasks.
lint: vuln
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running pre-commit...\033[0m"
	pre-commit run --all-files

#-------------------------------------------------------------------------------
# Testing
# https://github.com/golang/go/wiki/TableDrivenTests
# https://go.dev/doc/tutorial/fuzz
# https://pkg.go.dev/testing
# https://pkg.go.dev/golang.org/x/perf/cmd/benchstat

.PHONY: test
## test: [test]* Runs ALL tests.
test: unit acc

.PHONY: acc
## acc: [test] Runs Terraform provider acceptance tests. Set NAME= (without 'Test') to run a specific test by name
acc:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running acceptance tests...\033[0m"
	TF_ACC=1 $(GO) test -run=TestAcc$(NAME) -count=1 -parallel=$(shell nproc) -timeout 30s -coverpkg=./corefuncprovider/... -coverprofile=__coverage.out -v ./corefuncprovider/...

.PHONY: unit
## unit: [test] Runs unit tests. Set NAME= (without 'Test') to run a specific test by name
unit:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running unit tests...\033[0m"
	$(GO) test -run=Test$(NAME) -count=1 -parallel=$(shell nproc) -timeout 30s -coverpkg=./corefunc/... -coverprofile=__coverage.out -v ./corefunc/...

.PHONY: fuzz
## fuzz: [test]* Runs the fuzzer for 10 minutes. Set NAME= (without 'Fuzz') to run a specific test by name
fuzz:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Running the fuzzer (https://go.dev/doc/tutorial/fuzz)...\033[0m"
	$(GO) test -run='^$$' -fuzz=Fuzz$(NAME) -fuzztime 10m -parallel=$(shell nproc) -v ./corefunc/...

.PHONY: quickbench
## quickbench: [test]* Runs the benchmarks with minimal data for a quick check
quickbench:
	$(GO) test -bench=. -timeout 60m ./corefunc

.PHONY: bench
## bench: [test]* Runs the benchmarks with enough data for analysis with benchstat.
bench:
	$(GO) test -bench=. -count=6 -timeout 60m -benchmem -cpuprofile=__cpu.out -memprofile=__mem.out -trace=__trace.out ./corefunc | tee __bench-$(shell date --utc "+%Y%m%dT%H%M%SZ").out

.PHONY: view-cov
## view-cov: [test] After running test or unittest, this will launch a browser to view the coverage report.
view-cov:
	$(GO) tool cover -html=__coverage.out

.PHONY: view-cpupprof
## view-cpupprof: [test] After running bench, this will launch a browser to view the CPU profiler results.
view-cpupprof:
	$(GO) tool pprof -http :8080 __cpu.out

.PHONY: view-mempprof
## view-mempprof: [test] After running bench, this will launch a browser to view the memory profiler results.
view-mempprof:
	$(GO) tool pprof -http :8080 __mem.out

.PHONY: view-trace
## view-trace: [test] After running bench, this will launch a browser to view the trace results.
view-trace:
	$(GO) tool trace _trace.out

#-------------------------------------------------------------------------------
# Git Tasks

.PHONY: tag
## tag: [release]* Tags (and GPG-signs) the release.
tag:
	@ if [ $$(git status -s -uall | wc -l) != 1 ]; then echo 'ERROR: Git workspace must be clean.'; exit 1; fi;

	@echo "This release will be tagged as: $$(cat ./VERSION)"
	@echo "This version should match your release. If it doesn't, re-run 'make version'."
	@echo "---------------------------------------------------------------------"
	@read -p "Press any key to continue, or press Control+C to cancel. " x;

	@echo " "
	@chag update $$(cat ./VERSION)
	@echo " "

	@echo "These are the contents of the CHANGELOG for this release. Are these correct?"
	@echo "---------------------------------------------------------------------"
	@chag contents
	@echo "---------------------------------------------------------------------"
	@echo "Are these release notes correct? If not, cancel and update CHANGELOG.md."
	@read -p "Press any key to continue, or press Control+C to cancel. " x;

	@echo " "

	git add .
	git commit -a -m "Preparing the $$(cat ./VERSION) release."
	chag tag --sign

.PHONY: version
## version: [release]* Sets the version for the next release; pre-req for a release tag.
version:
	@echo "Current version: $$(cat ./VERSION)"
	@read -p "Enter new version number: " nv; \
	printf "$$nv" > ./VERSION
