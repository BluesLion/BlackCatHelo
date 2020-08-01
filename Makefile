export GOPATH := $(shell pwd)

all: blackcathelo

blackcathelo:
	find src/blackcathelo -name "*.go" -exec go fmt {} \;
	go install blackcathelo

test:
	go test -count=1 ./...

clean:
	rm -fr bin lib pkg
