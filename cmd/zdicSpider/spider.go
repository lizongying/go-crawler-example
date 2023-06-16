package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"strings"
	"time"
)

type Spider struct {
	pkg.Spider
	logger             pkg.Logger
	collectionZdicWord string
}

func (s *Spider) ParseCategory(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	items := x.FindStrMany(`//dt/a[@class="pck"]/@title`)
	for _, v := range items {
		e := s.YieldRequest(ctx, &pkg.Request{
			Http: pkg.Http{
				Url: fmt.Sprintf("https://www.zdic.net/zd/py/py/?py=%s", v),
			},
			CallBack: s.ParseList,
		})
		if e != nil {
			s.logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseList(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	items := x.FindStrMany(`//a/@href`)
	for _, v := range items {
		e := s.YieldRequest(ctx, &pkg.Request{
			Http: pkg.Http{
				Url: fmt.Sprintf("https://www.zdic.net%s", v),
			},
			CallBack: s.ParseDetail,
		})
		if e != nil {
			s.logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseDetail(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	fan := x.FindStrOne(`//span[text()="繁体"]/../a/text()`)
	id := response.Request.Url[strings.LastIndex(response.Request.Url, "/")+1:]
	data := DataWord{
		Id:  id,
		Fan: fan,
	}
	item := pkg.ItemMongo{
		Update:     true,
		Collection: s.collectionZdicWord,
		ItemUnimplemented: pkg.ItemUnimplemented{
			Id:   id,
			Data: &data,
		},
	}
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return
}

// Test go run cmd/zdicSpider/* -c dev.yml -m prod
func (s *Spider) Test(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, &pkg.Request{
		Http: pkg.Http{
			Url: fmt.Sprintf("https://www.zdic.net%s", "/hans/汉"),
		},
		CallBack: s.ParseDetail,
	})
	return
}

// FromCategory go run cmd/zdicSpider/* -c dev.yml -f FromCategory -m prod
func (s *Spider) FromCategory(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, &pkg.Request{
		Http: pkg.Http{
			Url: "https://www.zdic.net/zd/py/",
		},
		CallBack: s.ParseCategory,
	})

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	spider = &Spider{
		Spider:             baseSpider,
		logger:             baseSpider.GetLogger(),
		collectionZdicWord: "zdic_word",
	}
	spider.SetName("zdic")
	spider.SetTimeout(time.Minute)
	spider.SetInterval(200)
	spider.SetRetryMaxTimes(100)
	spider.SetPipeline(new(pipelines.MongoPipeline), 141)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
