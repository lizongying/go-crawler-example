package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider

	collectionBnu8105 string
}

func (s *Spider) ParseFind(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraFind)
	s.Logger.Info("Find", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	var respFind RespFind
	err = json.Unmarshal(response.BodyBytes, &respFind)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	for _, v := range respFind.Data.List {
		for _, v1 := range v.Data {
			e := s.YieldRequest(&pkg.Request{
				UniqueKey: v1.Hanzi,
				Extra: &ExtraSearch{
					Word: v1.Hanzi,
				},
				CallBack: s.ParseSearch,
			})
			if e != nil {
				s.Logger.Error(e)
				continue
			}
		}
	}

	return
}

func (s *Spider) ParseSearch(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraSearch)
	s.Logger.Info("Search", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	var respSearch RespSearch
	err = json.Unmarshal(response.BodyBytes, &respSearch)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	data := DataWord{
		Id:        respSearch.Data.Hanzi.ID,
		Fantizi:   respSearch.Data.Fantizi,
		Hanzi:     respSearch.Data.Hanzi,
		Pinyins:   respSearch.Data.Pinyins,
		ShowWrite: respSearch.Data.ShowWrite,
		Zhengzi:   respSearch.Data.Zhengzi,
		ZfjID:     respSearch.Data.ZfjID,
		Zfj:       respSearch.Data.Zfj,
	}
	item := pkg.ItemMongo{
		Update:     true,
		Collection: s.collectionBnu8105,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: extra.Word,
			Id:        respSearch.Data.Hanzi.ID,
			Data:      &data,
		},
	}
	err = s.YieldItem(&item)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	return
}

// Test go run cmd/bnuSpider/* -c dev.yml -m prod
func (s *Spider) Test(_ context.Context, _ string) (err error) {
	err = s.YieldRequest(&pkg.Request{
		Extra: &ExtraSearch{
			Word: "丰",
		},
		CallBack: s.ParseSearch,
	})
	return
}

// FromFind go run cmd/bnuSpider/* -c dev.yml -f FromFind -m prod
func (s *Spider) FromFind(_ context.Context, _ string) (err error) {
	for _, v := range []string{
		"1",
		//"2",
		//"3",
		//"4",
		//"5",
	} {
		err = s.YieldRequest(&pkg.Request{
			Extra: &ExtraFind{
				Bishun: v,
			},
			CallBack: s.ParseFind,
		})
	}

	return
}

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}

	baseSpider.Name = "bnu"
	baseSpider.
		SetMiddleware(NewMiddleware(logger), 90).
		SetMiddleware(middlewares.NewMongoMiddleware(logger, baseSpider.MongoDb), 141)
	spider = &Spider{
		BaseSpider:        baseSpider,
		collectionBnu8105: "bnu_8105",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
