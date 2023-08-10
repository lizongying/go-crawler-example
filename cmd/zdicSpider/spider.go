package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"strings"
	"time"
)

type Spider struct {
	pkg.Spider
	logger             pkg.Logger
	collectionZdicWord string
}

func (s *Spider) ParseCategory(ctx pkg.Context, response pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	for _, v := range x.FindStrMany(`//dt/a[@class="pck"]/@title`) {
		if e := s.YieldRequest(ctx, request.NewRequest().
			SetUrl(fmt.Sprintf("https://www.zdic.net/zd/py/py/?py=%s", v)).
			SetCallBack(s.ParseList)); e != nil {
			s.logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseList(ctx pkg.Context, response pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	for _, v := range x.FindStrMany(`//a/@href`) {
		if e := s.YieldRequest(ctx, request.NewRequest().
			SetUrl(fmt.Sprintf("https://www.zdic.net%s", v)).
			SetCallBack(s.ParseDetail)); e != nil {
			s.logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseDetail(ctx pkg.Context, response pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	fan := x.FindStrOne(`//span[text()="繁体"]/../a/text()`)
	id := response.GetUrl()[strings.LastIndex(response.GetUrl(), "/")+1:]
	data := DataWord{
		Id:  id,
		Fan: fan,
	}
	if err = s.YieldItem(ctx, items.NewItemMongo(s.collectionZdicWord, true).
		SetId(id).
		SetData(&data)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// Test go run cmd/zdicSpider/*.go -c dev.yml -n zdic -m prod
func (s *Spider) Test(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("https://www.zdic.net%s", "/hans/汉")).
		SetCallBack(s.ParseDetail)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// FromCategory go run cmd/zdicSpider/*.go -c dev.yml -n zdic -f FromCategory -m prod
func (s *Spider) FromCategory(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.zdic.net/zd/py/").
		SetCallBack(s.ParseCategory)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:             baseSpider,
		logger:             baseSpider.GetLogger(),
		collectionZdicWord: "zdic_word",
	}
	spider.WithOptions(
		pkg.WithName("zdic"),
		pkg.WithRetryMaxTimes(100),
		pkg.WithInterval(time.Second),
		pkg.WithTimeout(time.Minute),
		pkg.WithMongoPipeline(),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
