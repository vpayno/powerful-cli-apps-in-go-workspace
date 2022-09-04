#
# Makefile
#

.DEFAULT_GOAL := all

.PHONY: all run build version clean prepare gotest cover annotate bench check vet ineffassign lint gocyclo gocritic golangci-lint misspell vendor tidy gosec

BIN_FILE=./build/changeme

all: clean annotate check

check: vet ineffassign lint gocyclo gocritic golangci-lint misspell

build: clean prepare
	goreleaser build --rm-dist --single-target || goreleaser build --rm-dist --single-target --snapshot
	@printf "\n"

build-all:
	@printf "Building for every OS and Platform\n\n"
	time goreleaser build --rm-dist || time goreleaser build --rm-dist --snapshot
	@printf "\n"

run: build prepare
	go run ./...
	@printf "\n"

install: build
	./scripts/go-install
	@printf "\n"

clean:
	go clean
	rm -fv .coverage.*
	git restore ./cmd/*/.version.txt
	rm -fv ./cli-*
	@printf "\n"

vendor:
	go mod vendor
	@printf "\n"

tidy:
	go mod tidy
	@printf "\n"

prepare:
	go generate ./...
	@printf "\n"

gotest: clean prepare
	gotest -v -covermode=count -coverprofile=./reports/.coverage.out -cover ./...
	@printf "\n"

cover: gotest
	gocov convert ./reports/.coverage.out | gocov report
	@printf "\n"

annotate: cover
	gocov convert ./reports/.coverage.out | gocov annotate -ceiling=100 -color -
	@printf "\n"

badges: annotate
	./scripts/go-badges-coverage

bench: clean prepare
	go test --run=xxx --bench . --benchmem  ./...
	@printf "\n"

vet: prepare
	go vet ./...
	@printf "\n"

ineffassign: prepare
	ineffassign ./...
	@printf "\n"

lint: prepare
	revive ./...
	@printf "\n"

gocyclo: prepare
	gocyclo .
	@printf "\n"

gocritic: prepare
	gocritic check -enableAll ./...
	@printf "\n"

golangci-lint: prepare
	golangci-lint run --out-format line-number ./...
	@printf "\n"

misspell: prepare
	misspell -error .
	@printf "\n"

gosec: prepare
	gosec ./...
	@printf "\n"
