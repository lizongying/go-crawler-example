package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/bnu_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(bnu_spider.NewSpider).Run()
}
