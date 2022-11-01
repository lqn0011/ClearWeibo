package svc

import (
	"context"
	"strconv"

	"clean-wb/app/config"
	"clean-wb/basic/log"
	"clean-wb/basic/rhttp"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	emptyErr = errors.New("No more blog")
)

type ServiceContext struct {
	Config config.Config
	client *resty.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
		client: rhttp.NewResty(c.Resty),
	}

	if c.Debug {
		svc.client.SetDebug(true)
	}

	return svc
}

func (srv *ServiceContext) Close() {
}

func (srv *ServiceContext) GetMyBlog(ctx context.Context, page int64) (data MymblogData, err error) {
	var (
		cookie     = srv.Config.Cookie
		uid        = srv.Config.Uid
		xXsrfToken = srv.Config.XXsrfToken
		//traceParent = "xxx"
	)
	client := srv.client

	restyReq := rhttp.NewRequest(client, ctx)
	restyReq.SetHeader("authority", "weibo.com")
	restyReq.SetHeader("accept", "application/json, text/plain, */*")
	restyReq.SetHeader("accept-language", "zh-CN,zh;q=0.9")
	restyReq.SetHeader("cache-control", "no-cache")
	restyReq.SetHeader("client-version", "v2.36.13")
	restyReq.SetHeader("pragma", "no-cache")
	restyReq.SetHeader("referer", "https://weibo.com/"+strconv.FormatInt(uid, 10)+"?is_all=1")
	restyReq.SetHeader("sec-ch-ua", `"Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	restyReq.SetHeader("sec-ch-ua-mobile", "?0")
	restyReq.SetHeader("sec-ch-ua-platform", `"macOS"`)
	restyReq.SetHeader("sec-fetch-dest", "empty")
	restyReq.SetHeader("sec-fetch-mode", "cors")
	restyReq.SetHeader("sec-fetch-site", "same-origin")
	restyReq.SetHeader("server-version", "v2022.10.28.2")
	restyReq.SetHeader("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.24`)
	restyReq.SetHeader("x-requested-with", "XMLHttpRequest")
	restyReq.SetHeader("cookie", cookie)
	restyReq.SetHeader("x-xsrf-token", xXsrfToken)
	//restyReq.SetHeader("traceparent", traceParent)

	restyResp, err := restyReq.
		SetQueryParam("uid", strconv.FormatInt(uid, 10)).
		SetQueryParam("page", strconv.FormatInt(page, 10)).
		SetQueryParam("feature", "0").
		Get(MyBlogUrl)
	if err != nil {
		return
	}

	resp := MymblogResp{}
	err = jsoniter.Unmarshal(restyResp.Body(), &resp)
	if err != nil {
		return
	}
	return resp.Data, nil
}

func (srv *ServiceContext) DelBlog(ctx context.Context, id string) (err error) {
	var (
		xXsrfToken = srv.Config.XXsrfToken
		cookie     = srv.Config.Cookie
		uid        = srv.Config.Uid
		//traceParent = "xxx"
	)
	client := srv.client

	restyReq := rhttp.NewRequest(client, ctx)
	restyReq.SetHeader("authority", "weibo.com")
	restyReq.SetHeader("accept", "application/json, text/plain, */*")
	restyReq.SetHeader("accept-language", "zh-CN,zh;q=0.9")
	restyReq.SetHeader("cache-control", "no-cache")
	restyReq.SetHeader("client-version", "v2.36.13")
	restyReq.SetHeader("content-type", "application/json;charset=UTF-8")
	restyReq.SetHeader("origin", "https://weibo.com")
	restyReq.SetHeader("pragma", "no-cache")
	restyReq.SetHeader("referer", "https://weibo.com/"+strconv.FormatInt(uid, 10)+"?is_all=1")
	restyReq.SetHeader("sec-ch-ua", `"Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	restyReq.SetHeader("sec-ch-ua-mobile", "?0")
	restyReq.SetHeader("sec-ch-ua-platform", `"macOS"`)
	restyReq.SetHeader("sec-fetch-dest", "empty")
	restyReq.SetHeader("sec-fetch-mode", "cors")
	restyReq.SetHeader("sec-fetch-site", "same-origin")
	restyReq.SetHeader("server-version", "v2022.10.28.2")
	restyReq.SetHeader("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.24`)
	restyReq.SetHeader("x-requested-with", "XMLHttpRequest")
	restyReq.SetHeader("cookie", cookie)
	restyReq.SetHeader("x-xsrf-token", xXsrfToken)
	//restyReq.SetHeader("traceparent", traceParent)

	jsonBody := `{"id":"` + id + `"}`
	restyResp, err := restyReq.
		SetBody(jsonBody).
		Post(BlogDestroyUrl)
	if err != nil {
		return
	}

	resp := DelErrResp{}
	err = jsoniter.Unmarshal(restyResp.Body(), &resp)
	if err == nil {
		if resp.Ok == 0 {
			log.Error().Str("id", id).Str("err_msg", resp.Message).Msg("Del blog failed.")
		}
	}

	return
}

func (srv *ServiceContext) CleanAll(ctx context.Context) (err error) {
	var cnt = 0
	for {
		err = srv.cleanOnePage(ctx)
		if err == emptyErr {
			log.Info().Msg("Got empty error.")
			return
		}
		cnt++
		if cnt > 900 {
			log.Error().Msg("Reach counter max!")
			return
		}
	}
}

func (srv *ServiceContext) cleanOnePage(ctx context.Context) (err error) {
	data, err := srv.GetMyBlog(ctx, 1)
	if err != nil {
		log.Err(err).Msg("Get my blog failed.")
		return
	}

	if data.Total == 0 {
		return emptyErr
	}

	log.Info().Int("total", data.Total).Msg("Got one page.")
	for _, d := range data.List {
		var id string
		if len(d.OriMid) > 0 {
			id = d.OriMid
			log.Info().Str("id", d.OriMid).Msg("Deleting OriMid.")
		} else {
			//if d.Mid != "-1" {
			//	id = d.Mid
			//	log.Info().Str("id", d.Mid).Msg("Deleting Mid.")
			//} else {
			//	id = strconv.FormatInt(d.ID, 10)
			//	log.Info().Str("id", id).Msg("Deleting Id.")
			//}
			id = strconv.FormatInt(d.ID, 10)
			log.Info().Str("id", id).Msg("Deleting Id.")
		}
		err = srv.DelBlog(ctx, id)
		if err != nil {
			log.Err(err).Str("id", id).Msg("del blog failed.")
		}
	}

	return
}
