# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

[english](https://github.com/lizongying/go-crawler/README.md)

[中文](https://github.com/lizongying/go-crawler/README_CN.md)

[go-crawler](https://github.com/lizongying/go-crawler)

## Usage

[baidu-baike](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_baike_spider)

[baidu-image](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_image_spider)

[baidu-tieba](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_tieba_spider)

[baidu-zhidao](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/baidu_zhidao_spider)

[bnu](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/bnu_spider)

[douban](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/douban_spider)

[feixiaohao](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/feixiaohao_spider)

[levi](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/levi_spider)

[nike](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/nike_spider)

[youtube](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/youtube_spider)

[zdic](https://github.com/lizongying/go-crawler-example/tree/main/cmd/zdic_spider)

[zhihu](https://github.com/lizongying/go-crawler-example/tree/main/internal/spiders/zhihu_spider)

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
go get -u github.com/lizongying/go-crawler@42d155f
make

```

## Docker build

```shell
docker build -f ./cmd/baidu_baike_spider/Dockerfile -t crawler/baidu-baike-spider:latest . 

```

```shell
docker run -d crawler/baidu-baike-spider:latest spider -c example.yml -n baidu-baike -f Test -m once

```