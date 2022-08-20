PROJECT = ConsoleUI
VERSION := $(shell git describe --tag --abbrev=0)
SHA1 := $(shell git rev-parse HEAD)
NOW := $(shell date -u +'%Y%m%d-%H%M%S')
GIT_COMMIT := $(shell git rev-list -1 HEAD)

run:
	@go run _examples/basic.go

commit:
	/Users/ackers/.shell/commit "$(PROJECT) $(VERSION) $(SHA1) $(NOW)"
