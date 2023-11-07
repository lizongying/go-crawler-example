.PHONY: all

all: tidy multi_spider baidu_baike_spider baidu_image_spider baidu_tieba_spider baidu_zhidao_spider bnu_spider douban_spider feixiaohao_spider levi_spider nike_spider youtube_spider zdic_spider zhihu_spider

module := github.com/lizongying/go-crawler

shell:
	@echo 'SHELL='$(SHELL)

tidy:
	go mod tidy

multi_spider:
	go vet ./cmd/multi_spider
	go build -ldflags "-s -w" -o ./releases/multi_spider ./cmd/multi_spider

baidu_baike_spider:
	go vet ./cmd/baidu_baike_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_baike_spider" -o ./releases/baidu_baike_spider ./cmd/baidu_baike_spider

baidu_image_spider:
	go vet ./cmd/baidu_image_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_image_spider" -o ./releases/baidu_image_spider ./cmd/baidu_image_spider

baidu_tieba_spider:
	go vet ./cmd/baidu_tieba_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_tieba_spider" -o ./releases/baidu_tieba_spider ./cmd/baidu_tieba_spider

baidu_zhidao_spider:
	go vet ./cmd/baidu_zhidao_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_zhidao_spider" -o ./releases/baidu_zhidao_spider ./cmd/baidu_zhidao_spider

bnu_spider:
	go vet ./cmd/bnu_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=bnu_spider" -o ./releases/bnu_spider ./cmd/bnu_spider

douban_spider:
	go vet ./cmd/douban_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=douban_spider" -o ./releases/douban_spider ./cmd/douban_spider

feixiaohao_spider:
	go vet ./cmd/feixiaohao_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=feixiaohao_spider" -o ./releases/feixiaohao_spider ./cmd/feixiaohao_spider

levi_spider:
	go vet ./cmd/levi_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=levi_spider" -o ./releases/levi_spider ./cmd/levi_spider

nike_spider:
	go vet ./cmd/nike_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=nike_spider" -o ./releases/nike_spider ./cmd/nike_spider

youtube_spider:
	go vet ./cmd/youtube_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube_spider" -o ./releases/youtube_spider ./cmd/youtube_spider

zdic_spider:
	go vet ./cmd/zdic_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zdic_spider" -o ./releases/zdic_spider ./cmd/zdic_spider

zhihu_spider:
	go vet ./cmd/zhihu_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zhihu_spider" -o ./releases/zhihu_spider ./cmd/zhihu_spider