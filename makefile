BINARY=luffy
FORMAT=go fmt $$(go list ./... | grep -v /vendor/)
TESTS=go test $$(go list ./... | grep -v /vendor/) -race -cover

build:
	${FORMAT}
	${TESTS}
	go build -o ${BINARY}

install:
	${FORMAT}
	${TESTS}
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)


clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest
