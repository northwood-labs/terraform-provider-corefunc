##
# My muscle memory is to write `make` instead of `task`. This is a _very_
# lightweight shim which converts Makefile muscle memory to Taskfile patterns.
##

# @config-manager:start makefile
.PHONY: all
all:
	@ task --list --sort alphanumeric

.PHONY: %
%:
	@ task $@
# @config-manager:end makefile
