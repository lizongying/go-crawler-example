package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/levi_spider"
	"github.com/lizongying/go-crawler-example/internal/spiders/nike_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
)

// go run cmd/example_spider/*.go -c example.yml -n levi -m once
func main() {
	app.NewApp(
		//baidu_baike_spider.NewSpider,
		//baidu_image_spider.NewSpider,
		//baidu_tieba_spider.NewSpider,
		//baidu_zhidao_spider.NewSpider,
		//bnu_spider.NewSpider,
		//douban_spider.NewSpider,
		//feixiaohao_spider.NewSpider,
		levi_spider.NewSpider,
		nike_spider.NewSpider,
		//youtube_spider.NewSpider,
		//zdic_spider.NewSpider,
		//zhihu_spider.NewSpider,
	).Run(
		pkg.WithMockServerRoutes(mock_servers.NewRouteOk),
	)
}
