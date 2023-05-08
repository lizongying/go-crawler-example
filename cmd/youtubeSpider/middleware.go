package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/url"
)

const Video = "EgIQAQ%253D%253D"

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	urlSearch    string
	urlSearchApi string
	urlVideos    string

	apiKey string
}

func (m *Middleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request
	_, ok := request.Extra.(*ExtraSearch)
	if ok {
		extra := request.Extra.(*ExtraSearch)
		keyword := url.QueryEscape(extra.Keyword)
		request.Url = fmt.Sprintf(m.urlSearch, keyword)
		if extra.Sp == Video {
			request.Url += fmt.Sprintf("&sp=%s", Video)
		}
	}
	_, ok = request.Extra.(*ExtraSearchApi)
	if ok {
		extra := request.Extra.(*ExtraSearchApi)
		request.Method = "POST"
		request.Url = fmt.Sprintf(m.urlSearchApi, m.apiKey)
		request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"continuation":"%s"}`, extra.NextPageToken)
	}
	_, ok = request.Extra.(*ExtraVideos)
	if ok {
		extra := request.Extra.(*ExtraVideos)
		request.Url = fmt.Sprintf(m.urlVideos, extra.Id)
	}
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	err = c.NextRequest()
	return
}

func NewMiddleware(logger *logger.Logger) (middleware pkg.Middleware) {
	middleware = &Middleware{
		logger: logger,

		urlSearch:    "https://www.youtube.com/results?search_query=%s",
		urlSearchApi: "https://www.youtube.com/youtubei/v1/search?key=%s",
		urlVideos:    "https://www.youtube.com/@%s/videos",

		apiKey: "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8",
	}
	return
}
