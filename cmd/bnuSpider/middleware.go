package main

import (
	"context"
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

func (m *Middleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	var extraFind ExtraFind
	e := request.GetExtra(&extraFind)
	if e == nil {
		request.Url = fmt.Sprintf(m.urlFind, extraFind.Bishun)
	}
	var extraSearch ExtraSearch
	e = request.GetExtra(&extraSearch)
	if e == nil {
		request.Method = "POST"
		request.SetUrl(m.urlSearch)
		encryptedStr, _ := m.aes.Encrypt([]byte(extraSearch.Word))
		b := fmt.Sprintf(`ziFuJiId=%s&jstjId=%s&content=%s`, url.QueryEscape("zEm7A9LuQRXiTpuujAASv5ZkY8o5AP8y4FDl5qAte9PfHuy7vpDo6e6AzRRCBEKm"), url.QueryEscape("WagkdUR2Niv2c+IxZAl5V2sIf1yADd9a+TvoJFx0sd1dWfwAszERW4dywPjrLMOF"),
			url.QueryEscape(encryptedStr),
		)
		request.BodyStr = b
		request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		request.SetHeader("Accept", "application/json, text/plain, */*")
	}
	return
}

func (m *Middleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	response.BodyBytes, _ = m.aes.Decrypt(string(response.BodyBytes))
	return
}

func (m *Middleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	m.logger = crawler.GetLogger()
	m.urlSearch = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/hanzi/search"
	m.urlFind = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/bishun/find?content=%s&zifujiId=49c12ccb-35cc-437b-af4a-3fe126df8fca"
	m.aes, _ = utils.NewAes([]byte("crzjmwlcmgylxtyl"), utils.ECB)
	return m
}
