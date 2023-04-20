package main

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

type YoutubeMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger
}

func (m *YoutubeMiddleware) GetName() string {
	return "youtube"
}

func (m *YoutubeMiddleware) ProcessRequest(_ context.Context, r *pkg.Request) (request *pkg.Request, response *pkg.Response, err error) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	return
}

func NewYoutubeMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &YoutubeMiddleware{
		logger: logger,
	}
	return
}
