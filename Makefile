NAME = dash-light
PWD := $(MKPATH:%/Makefile=%)

help:
	@echo "Usage:"
	@echo "    make <target>"
	@echo
	@echo "Available targets: "
	@echo "    build                - performs a full build of the project (clean install check)"
	@echo "    compile				- creates a binary in bin directory of GOPATH"
	@echo "    check                - performs all verification tasks in the project"
	@echo "    coverage             - print a coverage report to terminal"
	@echo "    clean                - deletes the project vendor directory."
	@echo "    install              - download all dependencies"
	@echo "    lint                 - ensure code is standards compliant"
	@echo "    test            		- run tests"
	@echo


build:	clean install compile check

check:	test

clean :
	cd "$(PWD)"
	rm -rf vendor

compile:
	go install ./$(NAME)

coverage:
	echo 'mode: atomic' > coverage.txt && go list $(shell glide novendor) | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

fmt:
	go fmt ./...

test:
	go test -v $(shell glide novendor)

race:
	go test -race -v $(shell glide novendor)

install:
	glide install

default: help
