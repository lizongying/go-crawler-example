# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

[go-crawler](https://github.com/lizongying/go-crawler)

[english](https://github.com/lizongying/go-crawler/blob/main/README.md)

[中文](https://github.com/lizongying/go-crawler/blob/main/README_CN.md)

## Usage

[baidu-baike](./internal/spiders/baidu_baike_spider)

[baidu-image](./internal/spiders/baidu_image_spider)

[baidu-tieba](./internal/spiders/baidu_tieba_spider)

[baidu-zhidao](./internal/spiders/baidu_zhidao_spider)

[bnu](./internal/spiders/bnu_spider)

[douban](./internal/spiders/douban_spider)

[feixiaohao](./internal/spiders/feixiaohao_spider)

[levi](./internal/spiders/levi_spider)

[nike](./internal/spiders/nike_spider)

[youtube](./internal/spiders/youtube_spider)

[zdic](./internal/spiders/zdic_spider)

[zhihu](./internal/spiders/zhihu_spider)

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
go get -u github.com/lizongying/go-crawler@7c81a24
make

```

## Docker image

[ghcr.io](https://github.com/lizongying/go-crawler-example/pkgs/container/go-crawler-example)

[hub.docker.com](https://hub.docker.com/r/lizongying/go-crawler-example)

### build

```shell
# cross platform
docker buildx create --use
docker buildx inspect --bootstrap

# https://github.com/docker/buildx/issues/290

# for amd64
docker buildx build --platform linux/amd64 -f ./cmd/example_spider/Dockerfile -t lizongying/go-crawler-example/example-spider:latest .

# for arm64(mac m1)
docker buildx build --platform linux/arm64 -f ./cmd/example_spider/Dockerfile -t lizongying/go-crawler-example/example-spider:latest . --load

```

### run

```shell
# once
docker run -d lizongying/go-crawler-example/example-spider:latest -c example.yml -n levi -f TestList -m once

# manual
docker run -p 8090:8090 -d lizongying/go-crawler-example/example-spider:latest -c example.yml

```