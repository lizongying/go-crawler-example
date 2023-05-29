package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/url"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlFind   string
	urlSearch string
	aes       *utils.Aes
}

func (m *Middleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request
	_, ok := request.Extra.(*ExtraFind)
	if ok {
		extra := request.Extra.(*ExtraFind)
		request.Url = fmt.Sprintf(m.urlFind, extra.Bishun)
	}
	_, ok = request.Extra.(*ExtraSearch)
	if ok {
		extra := request.Extra.(*ExtraSearch)
		request.Method = "POST"
		request.Url = m.urlSearch
		e, _ := m.aes.Encrypt([]byte(extra.Word))
		b := fmt.Sprintf(`ziFuJiId=%s&jstjId=%s&content=%s`, url.QueryEscape("zEm7A9LuQRXiTpuujAASv5ZkY8o5AP8y4FDl5qAte9PfHuy7vpDo6e6AzRRCBEKm"), url.QueryEscape("WagkdUR2Niv2c+IxZAl5V2sIf1yADd9a+TvoJFx0sd1dWfwAszERW4dywPjrLMOF"),
			url.QueryEscape(e),
		)
		request.BodyStr = b
		request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		request.SetHeader("Accept", "application/json, text/plain, */*")
	}

	err = c.NextRequest()
	return
}

func (m *Middleware) ProcessResponse(c *pkg.Context) (err error) {
	response := c.Response

	err = c.NextResponse()
	e, _ := m.aes.Decrypt(string(response.BodyBytes))
	response.BodyBytes = e

	return
}

func (m *Middleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	m.urlSearch = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/hanzi/search"
	m.urlFind = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/bishun/find?content=%s&zifujiId=49c12ccb-35cc-437b-af4a-3fe126df8fca"
	return m
}

func NewMiddleware() pkg.Middleware {
	return &Middleware{}
}
