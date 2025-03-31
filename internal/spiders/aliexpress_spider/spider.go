package aliexpress_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"time"
)

type Spider struct {
	pkg.Spider
	logger                     pkg.Logger
	collectionAliexpressSearch string
	fingerprint                string
}

func (s *Spider) ParseSearch(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraSearch
	response.MustUnmarshalExtra(&extra)

	js := response.MustJson()
	for k, v := range js.ManySelector(`data.result.mods.itemList.content`) {
		data := DataDetail{
			Id:   fmt.Sprintf("%d_%d", extra.Page, k),
			Json: v.String(),
		}
		s.MustYieldItem(ctx, items.NewItemJsonl(s.collectionAliexpressSearch).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}

	if extra.Page < 3 {
		if err = s.YieldRequest(ctx, request.NewRequest().
			SetExtra(&ExtraSearch{
				Page: extra.Page + 1,
			}).
			SetFingerprint(s.fingerprint).
			SetCallBack(s.ParseSearch)); err != nil {
			s.logger.Error(err)
		}
	}

	return
}

// TestSearch go run cmd/aliexpress_spider/*.go -c dev.yml -n aliexpress -f TestSearch -m once
func (s *Spider) TestSearch(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraSearch{
			Page: 1,
		}).
		SetFingerprint(s.fingerprint).
		SetCallBack(s.ParseSearch)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:                     baseSpider,
		logger:                     baseSpider.GetLogger(),
		collectionAliexpressSearch: "aliexpress_search",
		fingerprint:                "771,49195-49199-49196-49200-52393-52392-49161-49171-49162-49172-49170-4865-4866-4867,5-10-11-13-65281-23-18-43-51,29-23-24-25,0",
	}
	spider.WithOptions(
		pkg.WithName("aliexpress"),
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithInterval(time.Second*10),
		pkg.WithRetryMaxTimes(1),
	)

	return
}
