# go-crawler-example

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

## Usage

### args

* -c config file. must set it.
* -f start func. default Test.
* -m mode. default test. prod? dev? or another something.

### config

* mongo.example.uri: mongo uri
* mongo.example.database: mongo database
* log.filename: Log file path. You can replace {name} with ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
* proxy.example: proxy
* request.concurrency: Number of concurrency
* request.interval: Request interval(seconds). If set to 0, it is the default interval. If set to a negative number, it
  is 0.
* request.timeout: Request timeout(seconds)

clone

```shell
git clone git@github.com:lizongying/go-crawler-example.git go-spiders
cd go-spiders

```

build

```shell
make
```

run

```shell
./releases/youtubeSpider -c dev.yml -f FromKeyword -m prod
```

update package

```shell
go get -u github.com/lizongying/go-crawler@d372e54 
```