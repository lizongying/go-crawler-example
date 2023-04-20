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
* log.long_file: if set true, will log full file path.
* log.level: DEBUG/INFO/WARN/ERROR
* proxy.example: proxy

clone

```shell
git clone git@github.com:lizongying/go-crawler-example.git
cd go-crawler-example

```

build

```shell
make
```

run

```shell
./releases/youtubeSpider -c example.yml -f FromKeyword -m prod
```

update package

```shell
go get -u github.com/lizongying/go-crawler@a1d1a9a
```