package util

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"code.byted.org/ad/gromore/lib/utils/enhancelog"
	"code.byted.org/gopkg/env"
	"github.com/bitly/go-simplejson"
)

var HTTPUtilIns = NewHTTPUtil()

type (
	IHTTPUtil interface {
		APIPost(ctx context.Context, url string, params map[string]interface{}, headers map[string]string) (*simplejson.Json, error)
	}

	HTTPUtil struct{}
)

func NewHTTPUtil() IHTTPUtil {
	return &HTTPUtil{}
}

func (h *HTTPUtil) APIPost(ctx context.Context, url string, params map[string]interface{}, headers map[string]string) (*simplejson.Json, error) {
	return APIPost(ctx, url, params, headers)
}

func Post(ctx context.Context, url string, params interface{}, headers map[string]string) (*http.Response, error) {
	param, err := json.Marshal(params)
	if err != nil {
		enhancelog.CtxError(ctx, "json Marshal failed: %s", err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(param)))
	if err != nil {
		enhancelog.CtxError(ctx, "new request failed: %s", err)
		return nil, err
	}
	enhancelog.CtxInfo(ctx, "post with url:[%s], params:[%+v]", url, params)
	return doBaseRequest(ctx, req, headers)
}

func GetResponseBody(ctx context.Context, resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		enhancelog.CtxError(ctx, "[doRequest] Failed to read body for err=%+v", err)
		return nil, err
	}
	_ = resp.Body.Close()
	return body, nil
}

// APIPost 发起HTTP POST请求
func APIPost(ctx context.Context, url string, params map[string]interface{}, headers map[string]string) (*simplejson.Json, error) {
	param, err := json.Marshal(params)
	if err != nil {
		enhancelog.CtxError(ctx, "json Marshal failed: %s", err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(param)))
	if err != nil {
		enhancelog.CtxError(ctx, "new request failed: %s", err)
		return nil, err
	}
	enhancelog.CtxInfo(ctx, "post with url:[%s], params:[%+v]", url, params)
	return doRequest(ctx, req, headers)
}

// APIGet 发起HTTP GET请求
func APIGet(ctx context.Context, url string, params map[string]string, headers map[string]string) (*simplejson.Json, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		enhancelog.CtxError(ctx, "[ApiGet] Failed to create new request for err=%+v", err)
		return nil, err
	}
	query := req.URL.Query()
	if len(params) > 0 {
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	enhancelog.CtxInfo(ctx, "[ApiGet] Print info,  url=%+v, params=%+v", url, params)
	return doRequest(ctx, req, headers)
}

func doRequest(ctx context.Context, req *http.Request, headers map[string]string) (*simplejson.Json, error) {
	resp, err := doBaseRequest(ctx, req, headers)
	if err != nil {
		enhancelog.CtxError(ctx, "[doRequest] Failed to read body for err=%+v", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		enhancelog.CtxError(ctx, "[doRequest] Failed to read body for err=%+v", err)
		return nil, err
	}
	_ = resp.Body.Close()
	enhancelog.CtxInfo(ctx, "[doRequest] Original response = %+v", string(body))
	result, err := simplejson.NewJson(body)
	if err != nil {
		enhancelog.CtxError(ctx, "[doRequest] Failed to parse body to json for err=%+v, resp=%+v", err, resp)
		return nil, err
	}
	return result, nil
}

func doBaseRequest(ctx context.Context, req *http.Request, headers map[string]string) (*http.Response, error) {
	req.Header.Add("X-Tt-LOGID", ctx.Value("K_LOGID").(string))
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	if req.Header.Get("X-Tt-Env") == "" && os.Getenv("SERVICE_ENV") != "" {
		req.Header.Add("X-Tt-Env", os.Getenv("SERVICE_ENV"))
	}

	if env.IsPPE() {
		req.Header.Add("X-Use-PPE", "1")
	}
	enhancelog.CtxInfo(ctx, "[doRequest] Print info, url=%+v, method=%+v, headers=%+v", req.URL, req.Method, req.Header)

	client := &http.Client{}
	return client.Do(req)
}
