#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

BINARY_NAME=terraform-provider-corefunc
BINARY_VERSION=$(shell cat ./VERSION | tr -d '\n')
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

#-------------------------------------------------------------------------------
# Global stuff.

GO=$(shell which go)
BREW_PREFIX=$(shell brew --prefix)
HOMEBREW_PACKAGES=git git-lfs goreleaser/tap/goreleaser-pro graphviz jq librsvg pre-commit python@3.10 qcachegrind

# Determine which version of `echo` to use. Use version from coreutils if available.
ECHOCHECK := $(shell command -v $(BREW_PREFIX)/opt/coreutils/libexec/gnubin/echo 2> /dev/null)
ifdef ECHOCHECK
    ECHO="$(BREW_PREFIX)/opt/coreutils/libexec/gnubin/echo" -e
else
    ECHO=echo -e
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
			if [[ "$$line" == *"]\*"* ]]; then \
				$(ECHO) "\033[1;33m$$line\033[0m"; \
			else \
				$(ECHO) "$$line"; \
			fi; \
		done

#-------------------------------------------------------------------------------
# Installation

.PHONY: _validate_deps
_validate_deps:
	@ cd ./scripts && python3 ./dependency-check.py && cd ..

.PHONY: install-tools-go
## install-tools-go: [tools]* Install/upgrade the required tools for macOS, including Go packages.
install-tools-go: _validate_deps
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/google/osv-scanner/cmd/osv-scanner@v1

.PHONY: install-tools-mac
## install-tools-mac: [tools]* Install/upgrade the required tools for macOS, including Go packages.
install-tools-mac: _validate_deps install-tools-go
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Installing required packages for macOS (Homebrew)...\033[0m"
	brew update && brew install $(HOMEBREW_PACKAGES) && brew upgrade $(HOMEBREW_PACKAGES)
	cargo install --locked cocogitto
	curl -sSLf https://raw.githubusercontent.com/mtdowling/chag/master/install.sh | bash

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33mTo update to the latest versions, run:\033[0m"
	@ $(ECHO) "\033[1;33m    brew update && brew upgrade\033[0m"
	@ $(ECHO) " "

.PHONY: install-hooks
## install-hooks: [tools]* Install/upgrade the Git hooks used for ensuring consistency.
install-hooks: _validate_deps
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Installing Git hooks...\033[0m"
	pre-commit install

	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33mLearn more about `cog` at:\033[0m"
	@ $(ECHO) "\033[1;33m    https://docs.cocogitto.io/guide/\033[0m"
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
	mkdir -p ./bin
	$(GO) mod tidy -go=1.21 -v
	$(GO) mod download -x
	$(GO) get -v ./...

.PHONY: build-release-prep
## build-release-prep: [build] Post-development, ready to release steps.
build-release-prep:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Download the Go dependencies...\033[0m"
	$(GO) mod download

.PHONY: build
## build: [build]* Builds and installs the Terraform provider.
build: tidy
	go install -v .

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
clean-tests: clean-logs
	@ echo " "
	@ echo "=====> Cleaning artifacts from tests..."
	- find . -type d -name ".terraform" | xargs rm -Rf
	- find ./ -type f -name ".terraform.lock.hcl" | xargs rm -Rf
	- find ./examples -type d -name "\.*" | xargs rm -Rf

.PHONY: clean-logs
## clean-logs: [clean] Removes the log files.
clean-logs:
	@ echo " "
	@ echo "=====> Cleaning log files..."
	- rm -Rf ./terratest-*
	- rm -Rf ./terraform.tfstate*
	- rm -Rf ./**/terratest-*
	- rm -Rf ./**/terraform.tfstate*

.PHONY: clean-tf
## clean-tf: [clean] Clean Terraform leftovers.
clean-tf:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Cleaning Terraform artifacts...\033[0m"
	find . -type d -name "terraform.d" | xargs -I% rm -Rfv "%"
	find . -type d -name ".terraform" | xargs -I% rm -Rfv "%"
	find . -type f -name ".terraform.lock.hcl" | xargs -I% rm -fv "%"

.PHONY: clean
## clean: [clean]* Runs ALL cleaning tasks.
clean: clean-tf clean-tests clean-logs

#-------------------------------------------------------------------------------
# Documentation

.PHONY: docs
## docs: [docs]* Runs ALL documentation tasks.
docs:
	go generate -v ./...

#-------------------------------------------------------------------------------
# Linting

.PHONY: golint
## golint: [lint] Runs `golangci-lint` (static analysis, formatting) against all Golang (*.go) code.
golint:
	@ echo " "
	@ echo "=====> Running golangci-lint..."
	golangci-lint run --fix ./...

.PHONY: markdownlint
## markdownlint: [lint] Runs `markdownlint` (formatting, spelling) against all Markdown (*.md) documents.
markdownlint:
	@ echo " "
	@ echo "=====> Running Markdownlint..."
	npx -y markdownlint-cli --fix '**/*.md' --ignore 'node_modules'

# Formats all Terraform code to its canonical format.
.PHONY: fmt
## fmt: [lint] Runs `terraform fmt` (formatting) against all Terraform (*.tf) code.
fmt: clean-tests
	@ echo " "
	@ echo "=====> Running Terraform format..."
	terraform fmt
	for dir in $$(find examples/ -type f -name "*.tf" | xargs -I% dirname % | uniq); do \
		cd "$$dir" && \
			terraform fmt ; \
		cd $(current_dir) ; \
	done

.PHONY: vuln
## vuln: [lint] Runs `govulncheck` (vulnerability scanning) against all Golang (*.go) code.
vuln:
	@ echo " "
	@ echo "=====> Running govulncheck (https://go.dev/blog/vuln)..."
	govulncheck -v ./...

	@ echo " "
	@ echo "=====> Running govulncheck -test (https://go.dev/blog/vuln)..."
	govulncheck -test -v ./...

	@ echo " "
	@ echo "=====> Running osv-scanner (https://osv.dev)..."
	osv-scanner -r .

.PHONY: lint
## lint: [lint]* Runs ALL linting/validation tasks.
lint: fmt markdownlint golint vuln

#-------------------------------------------------------------------------------
# Profiling

.PHONY: binsize
## binsize: [profiling] Generate a chart of dependencies and their sizes in the binary.
binsize:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Generating binsize chart...\033[0m"
	go tool nm -size bin/$(BINARY_NAME) | go-binsize-treemap > binsize.svg && \
		rsvg-convert --width=2000 --format=png --output="binsize.png" "binsize.svg";

.PHONY: pprof
## pprof: [profiling] Reads the profiler data (--profiler) and generates callgrind files.
pprof:
	@ $(ECHO) " "
	@ $(ECHO) "\033[1;33m=====> Generating callgrind files from pprof files...\033[0m"
	@ find . -type f -name "*.pprof" | xargs -I% bash -c '\
		$(ECHO) "%" && \
		cat "%" | pprof --callgrind bin/$(BINARY_NAME) "%" > "callgrind.%.profile";
	'

#-------------------------------------------------------------------------------
# Testing
# https://github.com/golang/go/wiki/TableDrivenTests
# https://go.dev/doc/tutorial/fuzz

.PHONY: test
## test: [test] Runs ALL tests -- Unit, Acceptance, and Fuzz.
test:
	TF_ACC=1 go test -count=1 -parallel=$(shell nproc) -timeout 30s -v ./...

.PHONY: unittest
## unittest: [test] Runs Unit and Fuzz tests. Faster than Acceptance tests.
unittest:
	go test -count=1 -parallel=$(shell nproc) -timeout 30s -v ./...

.PHONY: fuzz
## fuzz: [test] Runs the fuzzer for 10 minutes.
fuzz:
	go test -fuzz=Fuzz -fuzztime 10m -v ./corefunc/.

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
