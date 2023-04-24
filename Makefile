.PHONY: all

all: tidy youtubeSpider

module := github.com/lizongying/go-crawler

tidy:
	go mod tidy

youtubeSpider:
	go vet ./cmd/youtubeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube" -o ./releases/youtubeSpider ./cmd/youtubeSpider
