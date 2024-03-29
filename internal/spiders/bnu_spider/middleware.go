package bnu_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
	"net/url"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlFind   string
	urlSearch string
	aes       *utils.Aes
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	switch request.GetExtraName() {
	case "ExtraFind":
		var extraFind ExtraFind
		err = request.UnmarshalExtra(&extraFind)
		if err != nil {
			m.logger.Error(err)
			return
		}
		request.SetUrl(fmt.Sprintf(m.urlFind, extraFind.Bishun))
	case "ExtraSearch":
		var extraSearch ExtraSearch
		err = request.UnmarshalExtra(&extraSearch)
		if err != nil {
			m.logger.Error(err)
			return
		}
		request.SetMethod(http.MethodPost)
		request.SetUrl(m.urlSearch)
		encryptedStr, _ := m.aes.Encrypt([]byte(extraSearch.Word))
		b := fmt.Sprintf(`ziFuJiId=%s&jstjId=%s&content=%s`, url.QueryEscape("zEm7A9LuQRXiTpuujAASv5ZkY8o5AP8y4FDl5qAte9PfHuy7vpDo6e6AzRRCBEKm"), url.QueryEscape("WagkdUR2Niv2c+IxZAl5V2sIf1yADd9a+TvoJFx0sd1dWfwAszERW4dywPjrLMOF"),
			url.QueryEscape(encryptedStr),
		)
		request.SetBodyStr(b)
		request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		request.SetHeader("Accept", "application/json, text/plain, */*")
	}

	return
}

func (m *Middleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
	var bodyBytes []byte
	bodyBytes, err = m.aes.Decrypt(response.BodyStr())
	if err != nil {
		m.logger.Error(err)
		return
	}
	response.SetBodyBytes(bodyBytes)
	return
}

func (m *Middleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.urlSearch = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/hanzi/search"
	m.urlFind = "https://qxk.bnu.edu.cn/qxkapi/gjqxk/bishun/find?content=%s&zifujiId=49c12ccb-35cc-437b-af4a-3fe126df8fca"
	m.aes, _ = utils.NewAes([]byte("crzjmwlcmgylxtyl"), utils.ECB)
	return m
}
