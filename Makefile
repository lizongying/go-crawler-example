.PHONY: all

all: tidy youtubeSpider zdicSpider

module := github.com/lizongying/go-crawler

tidy:
	go mod tidy

bnuSpider:
	go vet ./cmd/bnuSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=bnu" -o ./releases/bnuSpider ./cmd/bnuSpider

youtubeSpider:
	go vet ./cmd/youtubeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube" -o ./releases/youtubeSpider ./cmd/youtubeSpider

zdicSpider:
	go vet ./cmd/zdicSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zdic" -o ./releases/zdicSpider ./cmd/zdicSpider