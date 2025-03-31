.PHONY: all

all: tidy multi_spider baidu_baike_spider baidu_image_spider baidu_tieba_spider baidu_zhidao_spider bnu_spider douban_spider feixiaohao_spider levi_spider nike_spider youtube_spider zdic_spider zhihu_spider example_spider

module := $(shell head -n 1 go.mod)
module := $(subst module ,,${module})
branch := $(shell git rev-parse --abbrev-ref HEAD)
commit := $(shell git rev-parse --short HEAD)
commit_time := $(shell git log -1 --format=%ct)

info:
	@echo 'SHELL='$(SHELL)
	@echo 'module='$(module)
	@echo 'branch='$(branch)
	@echo 'commit='$(commit)
	@echo 'commit_time='$(commit_time)

tidy:
	go mod tidy

multi_spider:
	go vet ./cmd/multi_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=multi_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/multi_spider ./cmd/multi_spider

baidu_baike_spider:
	go vet ./cmd/baidu_baike_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_baike_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/baidu_baike_spider ./cmd/baidu_baike_spider

baidu_image_spider:
	go vet ./cmd/baidu_image_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_image_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/baidu_image_spider ./cmd/baidu_image_spider

baidu_tieba_spider:
	go vet ./cmd/baidu_tieba_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_tieba_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/baidu_tieba_spider ./cmd/baidu_tieba_spider

baidu_zhidao_spider:
	go vet ./cmd/baidu_zhidao_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=baidu_zhidao_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/baidu_zhidao_spider ./cmd/baidu_zhidao_spider

bnu_spider:
	go vet ./cmd/bnu_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=bnu_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/bnu_spider ./cmd/bnu_spider

douban_spider:
	go vet ./cmd/douban_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=douban_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/douban_spider ./cmd/douban_spider

feixiaohao_spider:
	go vet ./cmd/feixiaohao_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=feixiaohao_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/feixiaohao_spider ./cmd/feixiaohao_spider

levi_spider:
	go vet ./cmd/levi_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=levi_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/levi_spider ./cmd/levi_spider

nike_spider:
	go vet ./cmd/nike_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=nike_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/nike_spider ./cmd/nike_spider

youtube_spider:
	go vet ./cmd/youtube_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=youtube_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/youtube_spider ./cmd/youtube_spider

zdic_spider:
	go vet ./cmd/zdic_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zdic_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/zdic_spider ./cmd/zdic_spider

zhihu_spider:
	go vet ./cmd/zhihu_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=zhihu_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/zhihu_spider ./cmd/zhihu_spider

example_spider:
	go vet ./cmd/example_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=example_spider -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/example_spider ./cmd/example_spider
