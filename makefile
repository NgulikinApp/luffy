BINARY=luffy

build: test
	go build -o ${BINARY}

test:
	./test-cover.sh

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

.PHONY: unittest test
