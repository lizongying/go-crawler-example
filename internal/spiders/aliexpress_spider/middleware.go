package aliexpress_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlSearch string
	header    map[string]string
	body      string
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	switch request.GetExtraName() {
	case "ExtraSearch":
		var extra ExtraSearch
		if err = request.UnmarshalExtra(&extra); err != nil {
			m.logger.Error(err)
			return
		}

		page := extra.Page
		request.SetBodyStr(fmt.Sprintf(m.body, page))
		request.SetUrl(m.urlSearch)
	}
	request.SetHeaders(m.header)
	return
}

func (m *Middleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.urlSearch = "https://www.aliexpress.com/fn/search-pc/index"
	m.header = map[string]string{
		"content-type": "application/json;charset=UTF-8",
		"cookie":       "_m_h5_tk_enc=47f9905d6f2907c9e57822683f71db27; x_router_us_f=x_alimid=6134255603; sgcookie=E100yKJwqgFwh+0tgzFQLAQrc5VGx4kHtHDFr5I7D1nZOE1Cv2U8uWaYbhix4C+uTspsgD91P2uMowA2FhrO5ftSrjgwkl9cYr3EeQMOqsGwHOI=; xman_f=/DvgoGBKUCliGcNVuHgYkhlEOM9IzfXILJ4mWpIAB0qS3sKwNo3vAhha9sGuQWOyd01l2uqC2Xl5dRbTPY1cFRWKtjh662iASljmGimknO2EcDviQnO2XO3eyXYccslvOgngyYj26rZu32KhsuMU3XMijaPTm46cxtJGKvCebbeuCLMic5RQf+rzqRTnsrytDnBBJx+EWT97AzALbEV8c7fyNKVBNeQQ1brIEcNqTSWziKckqseqKl1usnltPagNBE2GCvk/GMuRJAfuOeIs3NjH9X4uoxGD+lL02NRWGCyjQO8plXa/AqrNI55SaEKehOFSXdkI5UTjUztVcs40MPw4Ry4RWftcHCz9gAkFyQVY2lRHYRXjlQoK4PgcySE+0q8bP2jd75yUMTFl4IE702KoPXUsNJfmclQ794UMmeXrTtn9q8oXfQ==; aep_usuc_f=site=glo&c_tp=MYR&x_alimid=6134255603&re_sns=google&isb=y&region=MY&b_locale=en_US&ae_u_p_s=2;",
	}
	m.body = `{"pageVersion":"7ece9c0cc9cf2052db74f0d1b26b7033","target":"root","data":{"page":%d,"g":"y","SearchText":"shoes","origin":"y"},"eventName":"onChange","dependency":[]}`
	return m
}
