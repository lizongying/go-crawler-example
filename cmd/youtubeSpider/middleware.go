package main

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
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

func (m *Middleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	var extraSearch ExtraSearch
	e := request.GetExtra(&extraSearch)
	if e == nil {
		keyword := url.QueryEscape(extraSearch.Keyword)
		request.SetUrl(fmt.Sprintf(m.urlSearch, keyword))
		if extraSearch.Sp == Video {
			request.SetUrl(fmt.Sprintf(m.urlSearch, keyword) + fmt.Sprintf("&sp=%s", Video))
		}
	}
	var extraSearchApi ExtraSearchApi
	e = request.GetExtra(&extraSearchApi)
	if e == nil {
		request.Method = "POST"
		request.SetUrl(fmt.Sprintf(m.urlSearchApi, m.apiKey))
		request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"continuation":"%s"}`, extraSearchApi.NextPageToken)
	}
	var extraVideos ExtraVideos
	e = request.GetExtra(&extraVideos)
	if e == nil {
		request.SetUrl(fmt.Sprintf(m.urlVideos, extraVideos.Id))
	}
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	return
}

func (m *Middleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	m.logger = crawler.GetLogger()
	m.urlSearch = "https://www.youtube.com/results?search_query=%s"
	m.urlSearchApi = "https://www.youtube.com/youtubei/v1/search?key=%s"
	m.urlVideos = "https://www.youtube.com/@%s/videos"
	m.apiKey = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
	return m
}
