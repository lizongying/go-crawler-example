.PHONY: all

all:  youtubeSpider


youtubeSpider:
	go mod tidy
	go vet ./cmd/youtubeSpider
	go build -ldflags "-s -w" -o ./releases/youtubeSpider ./cmd/youtubeSpider
