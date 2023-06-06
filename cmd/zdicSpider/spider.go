package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/spider"
	"strings"
	"time"
)

type Spider struct {
	*spider.BaseSpider

	collectionZdicWord string
}

func (s *Spider) ParseCategory(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.Logger.Error(err)
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
			s.Logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseList(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.Logger.Error(err)
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
			s.Logger.Error(e)
			continue
		}
	}

	return
}

func (s *Spider) ParseDetail(ctx context.Context, response *pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.Logger.Error(err)
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
		s.Logger.Error(err)
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

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}

	baseSpider.Name = "zdic"
	baseSpider.Timeout = time.Minute
	baseSpider.Interval = 200
	baseSpider.RetryMaxTimes = 100
	baseSpider.
		SetMiddleware(middlewares.NewMongoMiddleware, 141)
	spider = &Spider{
		BaseSpider:         baseSpider,
		collectionZdicWord: "zdic_word",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
