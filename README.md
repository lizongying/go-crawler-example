# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)
[go-crawler](https://github.com/lizongying/go-crawler)
[english](https://github.com/lizongying/go-crawler/README.md)
[中文](https://github.com/lizongying/go-crawler/README_CN.md)

### clone

```shell
git clone git@github.com:lizongying/go-crawler-example.git crawler
cd crawler

```

### build

```shell
make
```

### run

```shell
./releases/bnuSpider -c dev.yml -n bnu -f Test -m dev
```

### update package

```shell
go get -u github.com/lizongying/go-crawler@6fc1b1c
```

## Docker build

```shell
docker build -f ./cmd/testSpider/Dockerfile -t crawler/baidu-baike-spider:latest . 
```

```shell
docker run -d crawler/baidu-baike-spider:latest spider -c example.yml -n test -f Test -m dev
```