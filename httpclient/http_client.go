package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/dajinkuang/dlog"
	"github.com/dajinkuang/errors"
	"github.com/dajinkuang/util"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	*http.Client
}

func NewHttpClient(timeout time.Duration) (httpClient *HttpClient) {
	httpClient = &HttpClient{
		util.HttpClient,
	}
	if timeout > 0 {
		httpClient.Timeout = timeout
	} else {
		httpClient.Timeout = 5 * time.Second
	}
	return
}

func NewTlsHttpClient(timeout time.Duration, cfg *tls.Config) (httpClient *HttpClient) {
	httpClient = NewHttpClient(timeout)
	httpClient.Transport.(*http.Transport).TLSClientConfig = cfg
	return
}

func (clt *HttpClient) Do(ctx context.Context, request *http.Request) (responseHeader http.Header, responseBody []byte, err error) {
	if ctx == nil {
		err = errors.New("ctx is nil")
		return
	}
	if request == nil {
		err = errors.New("request is nil")
		return
	}
	defer func() {
		if request != nil && request.Body != nil {
			request.Body.Close()
		}
	}()
	var (
		reqBytes   []byte
		e          error
		reqBodyStr string
	)
	if request != nil && request.Body != nil {
		reqBytes, e = ioutil.ReadAll(request.Body)
		if e != nil {
			err = errors.Wrap(e)
			return
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(reqBytes))
		reqBodyStr = string(reqBytes)
	}
	var start = time.Now()
	resp, e := clt.Client.Do(request)
	dur := time.Since(start) / time.Millisecond // ms
	var (
		respHeader http.Header
		respBytes  []byte
	)
	if resp == nil {
		err = errors.Wrap(e, "http do response is nil")
		dlog.Error("log_desc", "HttpClient.Do",
			"http_method", request.Method,
			"request-url", request.URL.Scheme+"//"+request.URL.Host+request.URL.Path+request.URL.RawQuery,
			"request-header", request.Header,
			"request-body", reqBodyStr,
			"response-header", "",
			"response-body", reqBodyStr,
			"time(ms)", dur,
			"error", err)
		return
	} else {
		respHeader = resp.Header
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	var respBodyStr string
	if resp != nil && resp.Body != nil {
		respBytes, e = ioutil.ReadAll(resp.Body)
		respBodyStr = string(respBytes)
		if e != nil {
			err = errors.Wrap(e)
			dlog.Error("log_desc", "HttpClient.Do",
				"HttpMethod", request.Method,
				"request-url", request.URL.Scheme+"//"+request.URL.Host+request.URL.Path+request.URL.RawQuery,
				"request-header", request.Header,
				"request-body", reqBodyStr,
				"response-header", respHeader,
				"response-body", respBodyStr,
				"time(ms)", dur,
				"error", err)
			return
		}
	}
	dlog.Info("log_desc", "HttpClient.Do",
		"HttpMethod", request.Method,
		"request-url", request.URL.Scheme+"//"+request.URL.Host+request.URL.Path+request.URL.RawQuery,
		"request-header", request.Header,
		"request-body", reqBodyStr,
		"response-header", respHeader,
		"response-body", respBodyStr,
		"time(ms)", dur)
	responseHeader = respHeader
	responseBody = respBytes
	return
}
