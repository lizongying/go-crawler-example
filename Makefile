.PHONY: all

all: tidy baiduBaikeSpider baiduTiebaSpider baiduZhidaoSpider bnuSpider doubanSpider feixiaohaoSpider nikeSpider youtubeSpider zdicSpider zhihuSpider

module := github.com/lizongying/go-crawler

shell:
	@echo 'SHELL='$(SHELL)

tidy:
	go mod tidy

baiduBaikeSpider:
	go vet ./cmd/baiduBaikeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baiduBaike" -o ./releases/baiduBaikeSpider ./cmd/baiduBaikeSpider

baiduTiebaSpider:
	go vet ./cmd/baiduTiebaSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baiduTieba" -o ./releases/baiduTiebaSpider ./cmd/baiduTiebaSpider

baiduZhidaoSpider:
	go vet ./cmd/baiduZhidaoSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baiduZhidao" -o ./releases/baiduZhidaoSpider ./cmd/baiduZhidaoSpider

bnuSpider:
	go vet ./cmd/bnuSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=bnu" -o ./releases/bnuSpider ./cmd/bnuSpider

doubanSpider:
	go vet ./cmd/doubanSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=douban" -o ./releases/doubanSpider ./cmd/doubanSpider

feixiaohaoSpider:
	go vet ./cmd/feixiaohaoSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=feixiaohao" -o ./releases/feixiaohaoSpider ./cmd/feixiaohaoSpider

nikeSpider:
	go vet ./cmd/nikeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=nike" -o ./releases/nikeSpider ./cmd/nikeSpider

youtubeSpider:
	go vet ./cmd/youtubeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube" -o ./releases/youtubeSpider ./cmd/youtubeSpider

zdicSpider:
	go vet ./cmd/zdicSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zdic" -o ./releases/zdicSpider ./cmd/zdicSpider

zhihuSpider:
	go vet ./cmd/zhihuSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zhihu" -o ./releases/zhihuSpider ./cmd/zhihuSpider