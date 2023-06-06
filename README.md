# go-crawler-example

go-crawler的爬虫示例。

[go-crawler-example](https://github.com/lizongying/go-crawler-example)
[go-crawler](https://github.com/lizongying/go-crawler)

## Usage

* 为了方便开发调试，增加了本地devServer，在`-m dev`模式下会默认启用。可以自定义route，仅需要实现`pkg.Route`
  ，然后在spider中通过`AddDevServerRoutes(...pkg.Route)`注册到devServer即可。
    * 支持http和https，可以设置`dev_server`。如http`http://localhost:8081`，https`https://localhost:8081`。
    * 默认显示ja3指纹。
    * 可以通过tls工具`tls`生成服务器的私钥和证书。
* 需要返回的item需要实现Item（可以组合ItemUnimplemented）
    * UniqueKey 唯一键，不会保存到数据库。可以用作过滤等其他用途。
    * Id 保存主键
    * Data 完整数据
    * 内置item：ItemCsv、ItemJsonl、ItemMongo、ItemMysql、ItemKafka
* 中间件的order不能重复。编写的时候不要忘记`nextRequest()`/`nextResponse()`/`nextItem()`
* 本框架舍弃了pipeline概念，功能合并到middleware。在很多情况下，功能会有交叉，合并后会更方便，同时编写也更简单。
* middleware包括框架内置、自定义公共（internal/middlewares）和自定义爬虫内（和爬虫同module）。
* 框架内置middleware，自定义middleware请参照以下order进行配置。
    * stats:100
    * device:101
        * 修改request设备信息。修改header和tls信息，暂时只支持user-agent随机切换。需要设置`SetPlatforms`和`SetBrowsers`
          限定设备范围。默认不启用。
        * 启用方法。`spider.SetMiddleware(middlewares.NewDeviceMiddleware, 101)`
        * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
        * Browsers: Chrome/Edge/Safari/FireFox
    * filter:110
        * 过滤重复请求。默认支持的是item保存成功后才会进入去重队列，防止出现请求失败后再次请求却被过滤的问题。所以当请求速度大于保存速度的时候可能会有请求不被过滤的情况。
    * retry:120
        * 如果请求出错，会进行重试。默认启用，`RetryMaxTimes=3`
    * referer:130
        * 通过设置`referrer_policy`采用不同的referer策略，
        * DefaultReferrerPolicy。会加入请求来源，默认
        * NoReferrerPolicy。不加入请求来源
    * http:140
    * dump:150
        * 在debug模式下打印item.data
        * 默认启用
    * csv
        * 保存结果到csv文件。
        * 需在在ItemCsv中设置`FileName`，保存的文件名称，不包含.csv
        * 启用方法。`spider.SetMiddleware(middlewares.NewCsvMiddleware, 151)`
    * jsonlines
        * 保存结果到jsonlines文件。
        * 需在在ItemJsonl中设置`FileName`，保存的文件名称，不包含.jsonl
        * 启用方法。`spider.SetMiddleware(middlewares.NewJsonlinesMiddleware, 152)`
    * mongo
        * 保存结果到mongo。
        * 需在在ItemMongo中设置`Collection`，保存的collection
        * 启用方法。`spider.SetMiddleware(middlewares.NewMongoMiddleware, 153)`
    * mysql
        * 保存结果到mysql。
        * 需在在ItemMysql中设置`table`，保存的table
        * 启用方法。`spider.SetMiddleware(middlewares.NewMysqlMiddleware, 154)`
    * kafka
        * 保存结果到kafka。
        * 需在在ItemKafka中设置`Topic`，保存的topic
        * 启用方法。`spider.SetMiddleware(middlewares.NewKafkaMiddleware, 155)`
* 在配置文件中可以配置全局request参数，在具体request中可以覆盖此配置
* 解析模块
    * query选择器 [go-query](https://github.com/lizongying/go-query)
        * ```response.Query()```
    * xpath选择器 [go-xpath](https://github.com/lizongying/go-xpath)
        * ```response.Xpath()```
    * gjson
        * ```response.Json()```
    * re选择器 [go-re](https://github.com/lizongying/go-re)
        * ```response.Re()```
* 代理
    * 可以自行搭建隧道代理 [go-proxy](https://github.com/lizongying/go-proxy)
      。这是一个随机切换的隧道代理，调用方无感知，方便使用。后期会加入一些其他的调用方式，比如维持原来的代理地址。

### args

* -c config file. must set it. 配置文件，必须配置。
* -f start func. default `Test`. 入口方法，默认`Test`。
* -a args. json string. 额外的参数，用于入口方法调用，非必须项。
* -m mode. default `test`. 启动模式，如`dev`,`prod`等，默认`test`

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
* dev_server: devServer。如http`http://localhost:8081`，https`https://localhost:8081`。

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
./releases/bnuSpider -c dev.yml -f Test -m dev
```

update package

```shell
go get -u github.com/lizongying/go-crawler@18aa07d
```