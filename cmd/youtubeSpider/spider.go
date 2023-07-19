package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Spider struct {
	pkg.Spider
	logger                pkg.Logger
	collectionYoutubeUser string
	proxyEnable           bool

	apiKey          string
	initialDataRe   *regexp.Regexp
	apiKeyRe        *regexp.Regexp
	emailRe         *regexp.Regexp
	urlRe           *regexp.Regexp
	floatRe         *regexp.Regexp
	intRe           *regexp.Regexp
	publishedTimeRe *regexp.Regexp
}

func (s *Spider) ParseSearch(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraSearch
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("Search", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	r := s.initialDataRe.FindSubmatch(response.GetBodyBytes())
	if len(r) != 2 {
		err = errors.New("not find content")
		s.logger.Error(err)
		return
	}
	var respSearch RespSearch
	err = json.Unmarshal(r[1], &respSearch)
	if err != nil {
		s.logger.Error(err)
		return
	}
	token := ""
	for _, v := range respSearch.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		continuationCommand := v.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand
		if continuationCommand.Request == "CONTINUATION_REQUEST_TYPE_SEARCH" {
			token = continuationCommand.Token
		} else {
			for _, v1 := range v.ItemSectionRenderer.Contents {
				if v1.VideoRenderer.VideoID == "" {
					continue
				}

				runs := v1.VideoRenderer.OwnerText.Runs
				if len(runs) < 1 {
					s.logger.Error("runs err")
					continue
				}
				id := strings.TrimPrefix(runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL, "/@")
				e := s.YieldRequest(ctx, request.NewRequest().
					SetExtra(&ExtraVideos{
						KeyWord:  extra.Keyword,
						Id:       id,
						Key:      runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
						UserName: runs[0].Text,
					}).
					SetCallBack(s.ParseVideos).
					SetProxyEnable(&s.proxyEnable).
					SetUniqueKey(id))
				if e != nil {
					s.logger.Error(e)
					continue
				}
			}
		}
	}

	r = s.apiKeyRe.FindSubmatch(response.GetBodyBytes())
	if len(r) != 2 {
		err = errors.New("not find api-key")
		s.logger.Error(err)
		return
	}

	s.apiKey = string(r[1])

	if extra.MaxPage > 0 && extra.Page >= extra.MaxPage {
		s.logger.Info("max page")
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraSearchApi{
			Keyword:       extra.Keyword,
			Sp:            extra.Sp,
			Page:          extra.Page + 1,
			MaxPage:       extra.MaxPage,
			NextPageToken: token,
		}).
		SetCallBack(s.ParseSearchApi).
		SetProxyEnable(&s.proxyEnable))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseSearchApi(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraSearchApi
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("SearchApi", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	var respSearch RespSearchApi
	err = response.UnmarshalBody(&respSearch)
	if err != nil {
		s.logger.Error(err)
		return
	}

	token := ""
	onResponseReceivedCommands := respSearch.OnResponseReceivedCommands
	if len(onResponseReceivedCommands) < 1 {
		err = errors.New("onResponseReceivedCommands err")
		s.logger.Error(err)
		return
	}

	for _, v := range onResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems {
		continuationCommand := v.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand
		if continuationCommand.Request == "CONTINUATION_REQUEST_TYPE_SEARCH" {
			token = continuationCommand.Token
		} else {
			for _, v1 := range v.ItemSectionRenderer.Contents {
				if v1.VideoRenderer.VideoID == "" {
					continue
				}

				runs := v1.VideoRenderer.OwnerText.Runs
				if len(runs) < 1 {
					s.logger.Error("runs err")
					continue
				}
				id := strings.TrimPrefix(runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL, "/@")
				e := s.YieldRequest(ctx, request.NewRequest().
					SetExtra(&ExtraVideos{
						KeyWord:  extra.Keyword,
						Id:       id,
						Key:      runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
						UserName: runs[0].Text,
					}).
					SetCallBack(s.ParseVideos).
					SetUniqueKey(id).
					SetProxyEnable(&s.proxyEnable))
				if e != nil {
					s.logger.Error(e)
					continue
				}
			}
		}
	}

	if token != "" {
		if extra.MaxPage > 0 && extra.Page >= extra.MaxPage {
			s.logger.Info("max page")
			return
		}
		err = s.YieldRequest(ctx, request.NewRequest().
			SetExtra(&ExtraSearchApi{
				Keyword:       extra.Keyword,
				Sp:            extra.Sp,
				Page:          extra.Page + 1,
				MaxPage:       extra.MaxPage,
				NextPageToken: token,
			}).
			SetCallBack(s.ParseSearchApi).
			SetProxyEnable(&s.proxyEnable))
		if err != nil {
			s.logger.Error(err)
			return
		}
	}

	return
}

func (s *Spider) ParseVideos(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraVideos
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("Videos", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	r := s.initialDataRe.FindSubmatch(response.GetBodyBytes())
	if len(r) != 2 {
		err = errors.New("not find content")
		s.logger.Error(err)
		return
	}
	var respVideos RespVideos
	err = json.Unmarshal(r[1], &respVideos)
	if err != nil {
		s.logger.Error(err)
		return
	}

	viewAvg := 0
	viewTotal := 0
	for _, v := range respVideos.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		if v.TabRenderer.Title != "Videos" {
			continue
		}

		i := 0
		for _, v1 := range v.TabRenderer.Content.RichGridRenderer.Contents {
			video := v1.RichItemRenderer.Content.VideoRenderer

			videoID := video.VideoID
			if videoID == "" {
				continue
			}

			viewCountText := video.ViewCountText.SimpleText
			viewCount := 0
			if viewCountText != "" && viewCountText != "No views" {
				viewCountInt, e := strconv.Atoi(strings.Join(s.intRe.FindAllString(viewCountText, -1), ""))
				if e != nil {
					s.logger.Error(e, "viewCount", viewCountText)
					continue
				}
				viewCount = viewCountInt
			}

			t := time.Now().Unix()
			publishedTime := s.publishedTimeRe.FindStringSubmatch(video.PublishedTimeText.SimpleText)
			if len(publishedTime) == 3 {
				i1, _ := strconv.Atoi(publishedTime[1])
				switch publishedTime[2] {
				case "year":
					t -= int64(i1 * 60 * 60 * 24 * 30 * 365)
				case "month":
					t -= int64(i1 * 60 * 60 * 24 * 30)
				case "week":
					t -= int64(i1 * 60 * 60 * 24 * 7)
				case "day":
					t -= int64(i1 * 60 * 60 * 24)
				case "hour":
					t -= int64(i1 * 60 * 60)
				case "minute":
					t -= int64(i1 * 60)
				case "second":
					t -= int64(i1)
				default:
				}
			}

			i++
			viewTotal += viewCount
			viewAvg = viewTotal / i
		}
	}

	subscriber := respVideos.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText
	index := strings.Index(subscriber, " ")
	followers := 0
	if index > 0 {
		followersText := subscriber[0:index]
		followers64, e := strconv.ParseFloat(strings.Join(s.floatRe.FindAllString(followersText, -1), ""), 64)
		if e != nil {
			s.logger.Error(e, "followers64", subscriber)
		}
		if strings.HasSuffix(followersText, "T") {
			followers = int(followers64 * 1000 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "G") {
			followers = int(followers64 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "M") {
			followers = int(followers64 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "K") {
			followers = int(followers64 * 1000)
		} else {
			followers = int(followers64)
		}
	}

	description := strings.TrimSpace(respVideos.Metadata.ChannelMetadataRenderer.Description)
	email := ""
	emails := s.emailRe.FindAllString(description, -1)
	if len(emails) > 0 {
		email = emails[0]
	}

	link := ""
	urls := s.urlRe.FindAllString(description, -1)
	if len(urls) > 0 {
		link = urls[0]
	}

	data := DataUser{
		Id:          extra.Id,
		UserName:    extra.UserName,
		Description: description,
		Link:        link,
		Email:       email,
		Followers:   followers,
		ViewAvg:     viewAvg,
		Keyword:     extra.KeyWord,
	}
	err = s.YieldItem(ctx, items.NewItemMongo(s.collectionYoutubeUser, true).
		SetUniqueKey(extra.Id).
		SetId(extra.Id).
		SetData(&data))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) Test(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraVideos{
			Id: "sierramarie",
		}).
		SetCallBack(s.ParseVideos).
		SetProxyEnable(&s.proxyEnable))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) FromKeyword(ctx context.Context, _ string) (err error) {
	for _, v := range []string{
		"veja",
	} {
		e := s.YieldRequest(ctx, request.NewRequest().
			SetExtra(&ExtraSearch{
				Keyword: v,
				Sp:      Video,
				Page:    1,
				MaxPage: 2,
			}).
			SetCallBack(s.ParseSearch).
			SetProxyEnable(&s.proxyEnable))
		if e != nil {
			s.logger.Error(e)
		}
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	spider = &Spider{
		Spider:                baseSpider,
		logger:                baseSpider.GetLogger(),
		collectionYoutubeUser: "youtube_user",
		proxyEnable:           true,

		apiKey:          "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8",
		initialDataRe:   regexp.MustCompile(`ytInitialData = (.+);</script>`),
		apiKeyRe:        regexp.MustCompile(`"INNERTUBE_API_KEY":"([^"]+)`),
		emailRe:         regexp.MustCompile(`(\w+[-+.]*\w+@\w+[-.]*\w+\.\w+[-.]*\w+)`),
		urlRe:           regexp.MustCompile(`(?i)\b((?:https?://|www\d{0,3}[.]|[a-z0-9.-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s\` + "`" + `!()\[\]{};:'".,<>?«»“”‘’]))`),
		floatRe:         regexp.MustCompile(`[\d.]`),
		intRe:           regexp.MustCompile(`\d`),
		publishedTimeRe: regexp.MustCompile(`(\d+)\s*(year|month|week|day|hour|minute|second)`),
	}
	spider.SetName("youtube")

	return
}

func main() {
	app.NewApp(NewSpider,
		pkg.WithTimeout(time.Second*30),
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithMongoPipeline(),
	).Run()
}
