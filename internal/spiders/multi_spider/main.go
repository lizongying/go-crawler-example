package main

import (
	"github.com/lizongying/go-crawler-example/internal/spiders/test1_spider"
	"github.com/lizongying/go-crawler-example/internal/spiders/test2_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
)

// go run cmd/multi_spider/*.go -c example.yml -n test1 -m once
func main() {
	app.NewApp(
		test1_spider.NewSpider,
		test2_spider.NewSpider,
	).Run(
		pkg.WithMockServerRoutes(mock_servers.NewRouteOk),
	)
}
