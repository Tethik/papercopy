default: | single-build

single-build:
	goreleaser build --single-target --snapshot --clean

build:
	goreleaser --snapshot --clean

# Used for manually releasing, normally running in github actions.
manual-release:
ifneq ($(shell git symbolic-ref --short HEAD),main)
	$(error Not on main branch)
endif
	goreleaser --clean

test:
	go test ./...

.PHONY: lint test release single-build