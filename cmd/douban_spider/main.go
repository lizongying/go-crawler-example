package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/douban_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(douban_spider.NewSpider).Run()
}
