package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/nike_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(nike_spider.NewSpider).Run()
}
