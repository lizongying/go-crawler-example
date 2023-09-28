# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

[english](https://github.com/lizongying/go-crawler/README.md)

[中文](https://github.com/lizongying/go-crawler/README_CN.md)

[go-crawler](https://github.com/lizongying/go-crawler)

## Usage

[baidu-baike](https://github.com/lizongying/go-crawler-example/tree/main/cmd/baiduBaikeSpider)

[baidu-tieba](https://github.com/lizongying/go-crawler-example/tree/main/cmd/baiduTiebaSpider)

[baidu-zhidao](https://github.com/lizongying/go-crawler-example/tree/main/cmd/baiduZhidaoSpider)

[bnu](https://github.com/lizongying/go-crawler-example/tree/main/cmd/bnuSpider)

[douban](https://github.com/lizongying/go-crawler-example/tree/main/cmd/doubanSpider)

[feixiaohao](https://github.com/lizongying/go-crawler-example/tree/main/cmd/feixiaohaoSpider)

[nike](https://github.com/lizongying/go-crawler-example/tree/main/cmd/nikeSpider)

[youtube](https://github.com/lizongying/go-crawler-example/tree/main/cmd/youtubeSpider)

[zdic](https://github.com/lizongying/go-crawler-example/tree/main/cmd/zdicSpider)

[zhihu](https://github.com/lizongying/go-crawler-example/tree/main/cmd/zhihuSpider)

### clone

```shell
git clone git@github.com:lizongying/go-crawler-example.git crawler

```

### build

```shell
cd crawler
make
```

### run

```shell
./releases/bnuSpider -c dev.yml -n bnu -f Test -m once
```

### update package

```shell
go get -u github.com/lizongying/go-crawler@latest
go get -u github.com/lizongying/go-crawler@db55a76
```

## Docker build

```shell
docker build -f ./cmd/testSpider/Dockerfile -t crawler/baidu-baike-spider:latest . 
```

```shell
docker run -d crawler/baidu-baike-spider:latest spider -c example.yml -n baidu-baike -f Test -m once
```