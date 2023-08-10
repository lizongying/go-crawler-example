package main

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/url"
	"regexp"
)

type Spider struct {
	pkg.Spider
	logger               pkg.Logger
	collectionBaiduBaike string
	reItem               *regexp.Regexp
}

func (s *Spider) ParseDetail(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraDetail
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))

	content := response.BodyText()
	if content == "" {
		err = errors.New("content empty")
		s.logger.Error(err)
		return
	}

	data := DataWord{
		Id:      extra.Keyword,
		Keyword: extra.Keyword,
		Content: content,
	}
	if err = s.YieldItem(ctx, items.NewItemMongo(s.collectionBaiduBaike, true).
		SetUniqueKey(extra.Keyword).
		SetId(extra.Keyword).
		SetData(&data)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseIndex(ctx pkg.Context, response pkg.Response) (err error) {
	links := response.AllLink()

	for _, v := range links {
		r := s.reItem.FindStringSubmatch(v.Path)
		if len(r) == 2 {
			decodedString, e := url.QueryUnescape(r[1])
			if e != nil {
				continue
			}
			if err = s.YieldRequest(ctx, request.NewRequest().
				SetExtra(&ExtraDetail{
					Keyword: decodedString,
				}).
				SetCallBack(s.ParseDetail)); err != nil {
				s.logger.Error(err)
				continue
			}
		} else {
			if err = s.YieldRequest(ctx, request.NewRequest().
				SetUrl(v.String()).
				SetCallBack(s.ParseIndex)); err != nil {
				s.logger.Error(err)
				continue
			}
		}
	}

	return
}

// Test go run cmd/baiduBaikeSpider/*.go -c dev.yml -n baidu-baike -m prod
func (s *Spider) Test(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraDetail{
			Keyword: "周口店遗址",
		}).
		SetCallBack(s.ParseDetail)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestIndex go run cmd/baiduBaikeSpider/*.go -c dev.yml -n baidu-baike -f TestIndex -m prod
func (s *Spider) TestIndex(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://baike.baidu.com/").
		SetCallBack(s.ParseIndex)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:               baseSpider,
		logger:               baseSpider.GetLogger(),
		collectionBaiduBaike: "baidu_baike",
		reItem:               regexp.MustCompile(`/item/([^/]+)`),
	}
	spider.WithOptions(
		pkg.WithName("baidu-baike"),
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithMongoPipeline(),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
