package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/baidu_tieba_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(baidu_tieba_spider.NewSpider).Run()
}
