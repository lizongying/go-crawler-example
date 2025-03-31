package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/aliexpress_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(aliexpress_spider.NewSpider).Run()
}
