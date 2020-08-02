export GOPATH := $(shell pwd)

all: pre blackcathelo

pre:
	@[ -d "github.com/bwmarrin/discordgo" ] || go get github.com/bwmarrin/discordgo

blackcathelo:
	find src/blackcathelo -name "*.go" -exec go fmt {} \;
	go install blackcathelo

test:
	go test -count=1 blackcathelo
	go test -count=1 lib/...

clean:
	sudo rm -fr bin lib pkg
