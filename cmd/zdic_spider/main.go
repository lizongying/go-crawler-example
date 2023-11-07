package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/zdic_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(zdic_spider.NewSpider).Run()
}
