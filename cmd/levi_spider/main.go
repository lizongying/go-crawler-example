package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/levi_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(levi_spider.NewSpider).Run()
}
