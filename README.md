# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

[go-crawler](https://github.com/lizongying/go-crawler)

[english](https://github.com/lizongying/go-crawler/blob/main/README.md)

[中文](https://github.com/lizongying/go-crawler/blob/main/README_CN.md)

## Usage

[baidu-baike](/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_baike_spider)

[baidu-image](/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_image_spider)

[baidu-tieba](/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_tieba_spider)

[baidu-zhidao](/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_zhidao_spider)

[bnu](/lizongying/go-crawler-example/tree/main/internal/spiders/bnu_spider)

[douban](/lizongying/go-crawler-example/tree/main/internal/spiders/douban_spider)

[feixiaohao](/lizongying/go-crawler-example/tree/main/internal/spiders/feixiaohao_spider)

[levi](/lizongying/go-crawler-example/tree/main/internal/spiders/levi_spider)

[nike](/lizongying/go-crawler-example/tree/main/internal/spiders/nike_spider)

[youtube](/lizongying/go-crawler-example/tree/main/internal/spiders/youtube_spider)

[zdic](/lizongying/go-crawler-example/tree/main/cmd/zdic_spider)

[zhihu](/lizongying/go-crawler-example/tree/main/internal/spiders/zhihu_spider)

### clone

```shell
git clone git@github.com:lizongying/go-crawler-example.git my-crawler
cd my-crawler

```

### develop

```shell
go run cmd/multi_spider/*.go -c example.yml -n test1 -m once

```

### build

```shell
make multi_spider
./releases/multi_spider -c prod.yml -n test1 -m once

```

### update go-crawler

```shell
go get -u github.com/lizongying/go-crawler@ef7431d
make

```

## Docker image

[ghcr.io](https://github.com/lizongying/go-crawler-example/pkgs/container/go-crawler-example)

[hub.docker.com](https://hub.docker.com/r/lizongying/go-crawler-example)

### build

```shell
# cross platform
docker buildx create --use

# for linux
docker buildx build --platform linux/amd64 -f ./cmd/example_spider/Dockerfile -t lizongying/go-crawler-example/example-spider:amd64 . --load

# for mac m1
docker buildx build --platform linux/arm64 -f ./cmd/example_spider/Dockerfile -t lizongying/go-crawler-example/example-spider:arm64 . --load

```

### run

```shell
# once
docker run -d lizongying/go-crawler-example/example-spider:arm64 -c example.yml -n levi -f TestList -m once

# manual
docker run -p 8090:8090 -d lizongying/go-crawler-example/example-spider:arm64 -c example.yml

```