package main

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-json/gjson"
)

type Spider struct {
	pkg.Spider
	logger             pkg.Logger
	collectionCategory string
	collectionDetail   string
}

func (s *Spider) ParseList(ctx pkg.Context, response pkg.Response) (err error) {
	js := response.MustJson()
	for _, v := range js.ManySelector(`data.products.products`) {
		data := DataDetail{
			Id:       v.One(`id`).String(),
			Title:    v.One(`title`).String(),
			Subtitle: v.One(`subtitle`).String(),
		}
		s.MustYieldItem(ctx, items.NewItemMongo(s.collectionDetail, true).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}
	next := js.One(`data.products.pages.next`).String()
	if next != "" {
		s.MustYieldRequest(ctx, request.NewRequest().
			SetUrl("https://api.nike.com/cic/browse/v2").
			SetQuery("queryid", "products").
			SetQuery("anonymousId", "861455C184756D93D923B7EB72650BFC").
			SetQuery("country", "us").
			SetQuery("endpoint", next).
			SetQuery("language", "en").
			SetQuery("localizedRangeStr", "{lowestPrice} — {highestPrice}").
			SetProxy("http://127.0.0.1:7890").
			SetCallBack(s.ParseList))
	}
	return
}

func (s *Spider) ParseIndex(ctx pkg.Context, response pkg.Response) (err error) {
	jsonStr := response.MustRe().One(`<script id="__NEXT_DATA__" type="application/json">(.+?)</script>`).String()
	js, err := gjson.NewSelectorFromStr(jsonStr)
	if err != nil {
		return err
	}

	for _, v := range js.ManySelector(`props.pageProps.initialState.Wall.facetNav.categories`) {
		data := DataCategory{
			Id:    v.One(`attributeId`).String(),
			Name:  v.One(`alternateName`).String(),
			Count: v.One(`resultCount`).Int(),
		}
		s.MustYieldItem(ctx, items.NewItemMongo(s.collectionCategory, true).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}

	for _, v := range js.ManySelector(`props.pageProps.initialState.Wall.products`) {
		data := DataDetail{
			Id:       v.One(`id`).String(),
			Title:    v.One(`title`).String(),
			Subtitle: v.One(`subtitle`).String(),
		}
		s.MustYieldItem(ctx, items.NewItemMongo(s.collectionDetail, true).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}

	next := js.One(`props.pageProps.initialState.Wall.pageData.next`).String()
	if next != "" {
		s.MustYieldRequest(ctx, request.NewRequest().
			SetUrl("https://api.nike.com/cic/browse/v2").
			SetQuery("queryid", "products").
			SetQuery("anonymousId", "861455C184756D93D923B7EB72650BFC").
			SetQuery("country", "us").
			SetQuery("endpoint", next).
			SetQuery("language", "en").
			SetQuery("localizedRangeStr", "{lowestPrice} — {highestPrice}").
			SetProxy("http://127.0.0.1:7890").
			SetCallBack(s.ParseList))
	}
	return
}

// TestList go run cmd/nikeSpider/*.go -c dev.yml -n nike -f TestList -m once
func (s *Spider) TestList(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://api.nike.com/cic/browse/v2").
		SetQuery("queryid", "products").
		SetQuery("anonymousId", "861455C184756D93D923B7EB72650BFC").
		SetQuery("country", "us").
		SetQuery("endpoint", "/product_feed/rollup_threads/v2?filter=marketplace%28US%29\u0026filter=language%28en%29\u0026filter=employeePrice%28true%29\u0026anchor=24\u0026consumerChannelId=d9a5bc42-4b9c-4976-858a-f159cf99c647\u0026count=24").
		SetQuery("language", "en").
		SetQuery("localizedRangeStr", "{lowestPrice} — {highestPrice}").
		SetProxy("http://127.0.0.1:7890").
		SetCallBack(s.ParseList))
	return
}

// TestIndex go run cmd/nikeSpider/*.go -c dev.yml -n nike -f TestIndex -m once
func (s *Spider) TestIndex(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.nike.com/w").
		SetProxy("http://127.0.0.1:7890").
		SetCallBack(s.ParseIndex))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:             baseSpider,
		logger:             baseSpider.GetLogger(),
		collectionCategory: "nike_category",
		collectionDetail:   "nike_detail",
	}
	spider.WithOptions(
		pkg.WithName("nike"),
		pkg.WithMongoPipeline(),
		pkg.WithRetryMaxTimes(0),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
