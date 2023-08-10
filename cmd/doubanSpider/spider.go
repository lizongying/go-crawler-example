package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"strconv"
)

type DataTop250 struct {
	Id      int    `bson:"_id" json:"id"`
	Img     string `json:"img"`
	Title   string `json:"title"`
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
	Quote   string `json:"quote"`
}

type Spider struct {
	pkg.Spider
	logger           pkg.Logger
	collectionTop250 string
}

func (s *Spider) ParseTop250(ctx pkg.Context, response pkg.Response) (err error) {
	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	for _, i := range x.FindNodeMany(`//div[@class="item"]`) {
		data := DataTop250{
			Id:      i.FindIntOne(`.//em/text()`),
			Img:     i.FindStrOne(`.//img/@src`),
			Title:   i.FindStrOne(`.//span[@class="title"]/text()`),
			Rating:  i.FindStrOne(`.//span[@class="rating_num"]/text()`),
			Comment: i.FindStrOne(`.//div[@class="star"]/span[last()]/text()`),
			Quote:   i.FindStrOne(`.//span[@class="inq"]/text()`),
		}
		if err = s.YieldItem(ctx, items.NewItemMongo(s.collectionTop250, true).
			SetUniqueKey(strconv.Itoa(data.Id)).
			SetId(data.Id).
			SetData(&data)); err != nil {
			s.logger.Error(err)
			continue
		}
	}
	return
}

// Test go run cmd/doubanSpider/*.go -c dev.yml -n douban-movie -f Test -m prod
func (s *Spider) Test(ctx pkg.Context, _ string) (err error) {
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("https://movie.douban.com/top250?start=%d", i*25)
		if err = s.YieldRequest(ctx, request.NewRequest().
			SetUrl(url).
			SetCallBack(s.ParseTop250)); err != nil {
			s.logger.Error(err)
			continue
		}
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:           baseSpider,
		logger:           baseSpider.GetLogger(),
		collectionTop250: "douban_movie_top250",
	}
	spider.WithOptions(
		pkg.WithName("douban-movie"),
		pkg.WithMongoPipeline(),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
