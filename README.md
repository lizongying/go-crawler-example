# go-crawler-example

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

## Feature

* 为了方便开发调试，增加了本地httpserver，在<code>-m dev</code>模式下会默认启用。可以自定义route，仅需要实现<code>
  pkg.Route</code>，然后在spider中通过<code>GetDevServer().AddRoutes(...pkg.Route)</code>注册到devServer即可。
* 编写中间件的时候需要注意，name不能重复。注册的时候order不能重复。注意不要忘记<code>nextRequest()</code>/<code>
  nextResponse()</code>
* 在配置文件中可以配置全局request参数，在具体request中可以覆盖此配置

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
* request.interval: Request interval(seconds). If set to 0, it is the default interval(1). If set to a negative number,
  it
  is 0.
* request.timeout: Request timeout(seconds)
* request.ok_http_codes: Request ok httpcodes
* request.retry_max_times: Request retry max times
* request.http_proto: Request http proto
* dev_addr: dev httpserver addr

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
go get -u github.com/lizongying/go-crawler@a594453 
```