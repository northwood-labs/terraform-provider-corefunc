##
# Default Makefile Settings
# These will run by default, unless the Makefile re-defines them.
#
# Helpful links:
#
# * https://makefiletutorial.com
# * https://stackoverflow.com/questions/11958626/make-file-warning-overriding-commands-for-target/49804748
# * https://www.gnu.org/software/make/manual/html_node/Make-Control-Functions.html
##

# Shell
SHELL:=bash
GO:=GOEXPERIMENT=greenteagc,jsonv2 $(shell which go)
GOTOOLS:=$(GO) tool -modfile=go.tools.mod
GUMCHECK:=$(shell command -v $(GOTOOLS) gum 2> /dev/null)

# Color stuff
HASH := \#
FOREGROUND:=$(HASH)FFFC67
YELLOW:=echo
WHITE:=echo
ERROR:=echo
BORDER:=echo
HEADER:=echo

ifdef GUMCHECK
	YELLOW:=$(GOTOOLS) gum style --foreground='$(FOREGROUND)' --bold
	WHITE:=$(GOTOOLS) gum style --bold
	ERROR:=$(GOTOOLS) gum style --foreground='$(HASH)FFFFFF' --background='$(HASH)CC0000' --bold --padding='0 1'
	BORDER:=$(GOTOOLS) gum style --foreground='$(HASH)FFFFFF' --border='rounded' --border-foreground='240' --padding='0 1' --margin='1 0'
	HEADER:=$(GOTOOLS) gum style --foreground='$(FOREGROUND)' --border='rounded' --border-foreground='12' --bold --width='78' --padding='0 1' --margin='1 0 0 0'
endif

# Tooling
NEXT_VERSION ?= $(shell git cliff --bump --unreleased --context | jq -r .[0].version)

#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all: help

.PHONY: help
## help: [help]* Prints this help message.
help:
ifndef GUMCHECK
	@ echo "==================================================================="
	@ echo "IMPORTANT:"
	@ echo "  Run 'make install-tools' to install developer tools."
	@ echo "==================================================================="
	@ echo ""
endif

	@ $(WHITE) "Usage:"
	@ echo ""
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | \
		while IFS= read -r line; do \
			if [[ "$$line" == *"]*"* ]]; then \
				$(YELLOW) "$$line"; \
			else \
				echo "$$line"; \
			fi; \
		done | LC_ALL=C sort --ignore-nonprinting --stable --ignore-leading-blanks --field-separator=' ' --key=4

#-------------------------------------------------------------------------------
# Checks

.PHONY: __req-on-host
__req-on-host:
	@ if [[ -f /.dockerenv ]]; then $(ERROR) "This command MUST be run on the host environment."; exit 1; fi

.PHONY: __req-inside-image
__req-inside-image:
	@ if [[ ! -f /.dockerenv ]]; then $(ERROR) "This command MUST be run inside the development environment."; exit 1; fi

#-------------------------------------------------------------------------------
# Mac-specific functions

.PHONY: clean-ds
## clean-ds: [host] Clean .DS_Store files.
clean-ds:
	@ $(HEADER) "=====> Cleaning .DS_Store files..."
	find . -type f -name ".DS_Store" | xargs -I% rm -fv "%"

#-------------------------------------------------------------------------------
# Default functions that should be overwritten as appropriate in local Makefiles

%: %-default
	@ true

.PHONY: clean-default
clean-default:
	@ $(ERROR) "No 'clean' task defined."

.PHONY: docs-default
docs-default:
	@ $(ERROR) "No 'docs' task defined."

.PHONY: lint-default
lint-default:
	@ $(ERROR) "No 'lint' task defined."

.PHONY: test-default
test-default:
	@ $(ERROR) "No 'test' task defined."

.PHONY: list-tests-default
list-tests-default:
	@ $(ERROR) "No 'list-tests' task defined."

.PHONY: acc-default
acc-default:
	@ $(ERROR) "No 'acc' task defined."

.PHONY: unit-default
unit-default:
	@ $(ERROR) "No 'unit' task defined."

.PHONY: mutate-default
mutate-default:
	@ $(ERROR) "No 'mutate' task defined."

.PHONY: terratest-default
terratest-default:
	@ $(ERROR) "No 'terratest' task defined."

.PHONY: examples-default
examples-default:
	@ $(ERROR) "No 'examples' task defined."

.PHONY: fuzz-default
fuzz-default:
	@ $(ERROR) "No 'fuzz' task defined."

.PHONY: quickbench-default
quickbench-default:
	@ $(ERROR) "No 'quickbench' task defined."

.PHONY: bench-default
bench-default:
	@ $(ERROR) "No 'bench' task defined."
