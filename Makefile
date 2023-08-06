.PHONY: all

all: shell tidy baiduBaikeSpider bnuSpider youtubeSpider zdicSpider

module := github.com/lizongying/go-crawler

shell:
	@echo 'SHELL='$(SHELL)

tidy:
	go mod tidy

baiduBaikeSpider:
	go vet ./cmd/baiduBaikeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baiduBaike" -o ./releases/baiduBaikeSpider ./cmd/baiduBaikeSpider

bnuSpider:
	go vet ./cmd/bnuSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=bnu" -o ./releases/bnuSpider ./cmd/bnuSpider

youtubeSpider:
	go vet ./cmd/youtubeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube" -o ./releases/youtubeSpider ./cmd/youtubeSpider

zdicSpider:
	go vet ./cmd/zdicSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zdic" -o ./releases/zdicSpider ./cmd/zdicSpider

zhihuSpider:
	go vet ./cmd/zhihuSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zhihu" -o ./releases/zhihuSpider ./cmd/zhihuSpider