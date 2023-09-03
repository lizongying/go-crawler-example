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

func (s *Spider) ParseQuestion(_ pkg.Context, response pkg.Response) (err error) {
	var dataPosts DataQuestions
	response.MustUnmarshalData(&dataPosts)
	utils.DumpJsonPretty(dataPosts.Data)
	return
}

// TestQuestion go run cmd/baiduZhidaoSpider/*.go -c example.yml -n baidu-zhidao -f TestQuestion -m once
func (s *Spider) TestQuestion(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://zhidao.baidu.com/question/1678301174738526227").
		SetCallBack(s.ParseQuestion))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("baidu-zhidao"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
