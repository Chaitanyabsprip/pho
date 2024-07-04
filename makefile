.DEFAULT_GOAL:=build
SOURCES := $(shell find . -type f -name '*.go')
CLI_PKG_PATH=./cmd/pho

build: ${CLI_PKG_PATH}/main.go
	@mkdir bin 2> /dev/null
	@go build -o ./bin ${CLI_PKG_PATH}

note:
	@mkdir bin 2> /dev/null
	@go build -o ./bin ${CLI_PKG_PATH}

clean:
	@rm -rd ./bin

install:
	@go install ${CLI_PKG_PATH}

uninstall: clean
	@rm "$(which pho)"

watch: 
	@fd -tf -E '*.go' | entr sh -c 'clear; printf "making..."; ${MAKE} install; printf made'

