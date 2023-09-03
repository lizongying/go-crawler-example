package main

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParsePost(_ pkg.Context, response pkg.Response) (err error) {
	var dataPosts DataPosts
	response.MustUnmarshalData(&dataPosts)
	utils.DumpJsonPretty(dataPosts.Data)
	return
}

// TestPost go run cmd/baiduTiebaSpider/*.go -c example.yml -n baidu-tieba -f TestPost -m once
func (s *Spider) TestPost(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://tieba.baidu.com/p/8568462599?pn=2&ajax=1&t=1693665791192").
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetCallBack(s.ParsePost))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("baidu-tieba"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
