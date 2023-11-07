package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/zhihu_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(zhihu_spider.NewSpider).Run()
}
