package main

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseImage(ctx pkg.Context, _ pkg.Response) (err error) {
	s.MustYieldItem(ctx, items.NewItemJsonl("baidu_image").
		SetData(&DataImage{}).
		SetImagesRequest([]pkg.Request{
			request.NewRequest().SetUrl("https://www.baidu.com/img/pcyayunhuikaimushidoodle_35c0ef27c30a077f2e46ddb5db1993ef.gif"),
		}))
	return
}

// TestImage go run cmd/baiduImageSpider/*.go -c example.yml -n baidu-image -f TestImage -m once
func (s *Spider) TestImage(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.baidu.com").
		SetCallBack(s.ParseImage))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("baidu-image"),
		pkg.WithJsonLinesPipeline(),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
