# go-crawler-example

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

## Feature

* 为了方便开发调试，增加了本地devServer，在<code>-m dev</code>模式下会默认启用。可以自定义route，仅需要实现<code>
  pkg.Route</code>，然后在spider中通过<code>AddDevServerRoutes(...pkg.Route)</code>注册到devServer即可。
* 中间件的order不能重复。编写的时候不要忘记<code>nextRequest()</code>/<code>nextResponse()</code>/<code>nextItem()</code>
* 本框架舍弃了pipeline概念，功能合并到middleware。在很多情况下，功能会有交叉，合并后会更方便，同时编写也更简单。
* middleware包括框架内置、自定义公共（internal/middlewares）和自定义爬虫内（和爬虫同module）。
* 框架内置middleware和默认order，建议自定义ProcessItem的中间件order大于140
    * stats:100
    * filter:110
    * retry:120
    * http:130
    * dump:140 在debug模式下打印item.data
    * csv
    * jsonlines
    * mongo
    * mysql
    * kafka
* 在配置文件中可以配置全局request参数，在具体request中可以覆盖此配置
* 解析模块
    * query选择器. response.Query() [go-query](https://github.com/lizongying/go-query)
    * xpath选择器. response.Xpath() [go-xpath](https://github.com/lizongying/go-xpath)
    * gjson. response.Json()
* 代理
    * 可以自行搭建隧道代理 [go-proxy](https://github.com/lizongying/go-proxy)

## Usage

### args

* -c config file. must set it.
* -f start func. default Test.
* -a args. json string.
* -m mode. default test. prod? dev? or another something.

### config

* mongo.example.uri: mongo uri
* mongo.example.database: mongo database
* log.filename: Log file path. You can replace {name} with ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
* proxy.example: proxy
* request.concurrency: Number of concurrency
* request.interval: Request interval(Millisecond). If set to 0, it is the default interval(1). If set to a negative
  number,
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
go get -u github.com/lizongying/go-crawler@eda774b 
```