package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/baidu_image_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(baidu_image_spider.NewSpider).Run()
}
