package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/baidu_zhidao_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(baidu_zhidao_spider.NewSpider).Run()
}
