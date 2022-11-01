package rhttp

import (
	"context"
	"time"

	"clean-wb/basic/log"

	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Config struct {
	TimeoutSeconds int
	RetryCount     int
	Trace          bool
}

func NewResty(c *Config) (client *resty.Client) {
	client = resty.New()
	if c.Trace {
		client.OnBeforeRequest(traceHookBefore)
		client.OnAfterResponse(traceHookAfter)
		client.OnError(traceHookError)
	}

	client.SetTimeout(time.Duration(c.TimeoutSeconds) * time.Second)
	client.SetRetryCount(c.RetryCount)
	return
}

func NewRequest(client *resty.Client, ctx context.Context) *resty.Request {
	return client.R().SetContext(ctx)
}

func traceHookBefore(c *resty.Client, req *resty.Request) error {
	if req.Attempt > 1 {
		return nil
	}

	operationName := "resty:" + req.Method + ":" + req.URL

	span, ctx := opentracing.StartSpanFromContext(req.Context(), operationName)

	ext.HTTPUrl.Set(span, req.URL)
	ext.HTTPMethod.Set(span, req.Method)

	//Todo... append req params to span
	//toStringHeader, _ := jsoniter.MarshalToString(req.Header)
	//toStringQuery, _ := jsoniter.MarshalToString(req.QueryParam)
	//toStringBody, _ := jsoniter.MarshalToString(req.Body)
	//span.LogKV(
	//	"header", toStringHeader,
	//	"req_body", toStringBody,
	//	"req_params", toStringQuery,
	//)

	req.SetContext(ctx)

	err := span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		log.Err(err).Msg("resty track hook before inject span to req header failed")
	}

	return nil
}

func traceHookAfter(c *resty.Client, resp *resty.Response) error {
	ctx := resp.Request.Context()
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	defer span.Finish()
	ext.HTTPStatusCode.Set(span, uint16(resp.StatusCode()))
	if resp.IsError() {
		ext.Error.Set(span, true)
		span.LogKV("error", resp.String())
	}
	//Todo append response data ...
	//span.LogKV(
	//	"response", resp.String(),
	//)

	return nil
}

func traceHookError(req *resty.Request, err error) {
	ctx := req.Context()
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	defer span.Finish()
	ext.Error.Set(span, true)
	span.SetTag("error.message", err.Error())
	if v, ok := err.(*resty.ResponseError); ok {
		// v.Response contains the last response from the server
		// v.Err contains the original error
		resp := v.Response
		ext.HTTPStatusCode.Set(span, uint16(resp.StatusCode()))
		//Todo append response data ...
		//span.LogKV(
		//	"response", resp.String(),
		//)
	}
	// Log the error, increment a metric, etc...
}
