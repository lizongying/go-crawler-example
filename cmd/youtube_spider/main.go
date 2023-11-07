package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/youtube_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(youtube_spider.NewSpider).Run()
}
