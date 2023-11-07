package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/feixiaohao_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(feixiaohao_spider.NewSpider).Run()
}
