package main

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger            pkg.Logger
	collectionBnu8105 string
}

func (s *Spider) ParseFind(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraFind
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("Find", utils.JsonStr(extra))

	var respFind RespFind
	err = response.UnmarshalBody(&respFind)
	if err != nil {
		s.logger.Error(err)
		return
	}

	for _, v := range respFind.Data.List {
		for _, v1 := range v.Data {
			e := s.YieldRequest(ctx, request.NewRequest().
				SetExtra(&ExtraSearch{
					Word: v1.Hanzi,
				}).
				SetCallBack(s.ParseSearch).
				SetUniqueKey(v1.Hanzi))
			if e != nil {
				s.logger.Error(e)
			}
		}
	}

	return
}

func (s *Spider) ParseSearch(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraSearch
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("Search", utils.JsonStr(extra))

	var respSearch RespSearch
	err = response.UnmarshalBody(&respSearch)
	if err != nil {
		s.logger.Error(err)
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

	err = s.YieldItem(ctx, items.NewItemMongo(s.collectionBnu8105, true).
		SetUniqueKey(extra.Word).
		SetId(respSearch.Data.Hanzi.ID).
		SetData(&data))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// Test go run cmd/bnuSpider/*.go -c dev.yml -n bnu -m prod
func (s *Spider) Test(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraSearch{
			Word: "丰",
		}).
		SetCallBack(s.ParseSearch))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// FromFind go run cmd/bnuSpider/*.go -c dev.yml -n bnu -f FromFind -m prod
func (s *Spider) FromFind(ctx pkg.Context, _ string) (err error) {
	for _, v := range []string{
		"1",
		//"2",
		//"3",
		//"4",
		//"5",
	} {
		e := s.YieldRequest(ctx, request.NewRequest().
			SetExtra(&ExtraFind{
				Bishun: v,
			}).
			SetCallBack(s.ParseFind))
		if e != nil {
			s.logger.Error(e)
		}
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	spider = &Spider{
		Spider:            baseSpider,
		logger:            baseSpider.GetLogger(),
		collectionBnu8105: "bnu_8105",
	}
	spider.WithOptions(
		pkg.WithName("bnu"),
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithMongoPipeline(),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
