package youtube_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"net/url"
)

const Video = "EgIQAQ%253D%253D"

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlSearch    string
	urlSearchApi string
	urlVideos    string

	apiKey string
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	switch request.GetExtraName() {
	case "ExtraSearch":
		var extraSearch ExtraSearch
		err = request.UnmarshalExtra(&extraSearch)
		if err != nil {
			m.logger.Error(err)
			return
		}
		keyword := url.QueryEscape(extraSearch.Keyword)
		request.SetUrl(fmt.Sprintf(m.urlSearch, keyword))
		if extraSearch.Sp == Video {
			request.SetUrl(fmt.Sprintf(m.urlSearch, keyword) + fmt.Sprintf("&sp=%s", Video))
		}
	case "ExtraSearchApi":
		var extraSearchApi ExtraSearchApi
		err = request.UnmarshalExtra(&extraSearchApi)
		if err != nil {
			m.logger.Error(err)
			return
		}
		request.SetMethod(http.MethodPost)
		request.SetUrl(fmt.Sprintf(m.urlSearchApi, m.apiKey))
		request.SetBodyStr(fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"continuation":"%s"}`, extraSearchApi.NextPageToken))
	case "ExtraVideos":
		var extraVideos ExtraVideos
		err = request.UnmarshalExtra(&extraVideos)
		if err != nil {
			m.logger.Error(err)
			return
		}
		request.SetUrl(fmt.Sprintf(m.urlVideos, extraVideos.Id))
	}

	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	return
}

func (m *Middleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.urlSearch = "https://www.youtube.com/results?search_query=%s"
	m.urlSearchApi = "https://www.youtube.com/youtubei/v1/search?key=%s"
	m.urlVideos = "https://www.youtube.com/@%s/videos"
	m.apiKey = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
	return m
}
