package request

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"net/http/cookiejar"
	"net/http/httptrace"

	log "github.com/sirupsen/logrus"
)

type HandResponse interface {
	HandleResponse()
}

type Request struct {
	url string

	host   string
	path   string
	scheme string
	method string

	rawQuery     string
	timeout      time.Duration
	headers      http.Header
	totalRequest int
	parallel     int
	debug        bool

	contentType string
	body        io.ReadWriter

	req    *http.Request
	cookie map[string]*cookiejar.Jar

	HandResp   HandResponse
	transport  http.Transport
	Httptrace  *httptrace.ClientTrace
	HttpClient *http.Client
}

func NewRequest() *Request {
	return nil
}

func (r *Request) Scheme(scheme string) *Request {
	r.scheme = scheme
	return r
}

func (r *Request) ContentType(contentType string) *Request {
	r.contentType = contentType
	return r
}

func (r *Request) RawQuery(rawQuery string) *Request {
	r.rawQuery = rawQuery
	return r
}

func (r *Request) IsDebug() bool {
	return r.debug
}

func (r *Request) EnableDebug() *Request {
	r.debug = true
	return r
}

func (r *Request) parseUrl(urlstr string) *Request {
	urlbody, err := url.Parse(urlstr)
	if err != nil {
		log.Errorf("parse url failed : ", err.Error())
		return r
	}
	r.path = urlbody.Path
	r.scheme = urlbody.Scheme
	r.host = urlbody.Host
	r.rawQuery = urlbody.RawQuery
	return r
}

func (r *Request) Url(utlstr string) *Request {
	r.url = utlstr
	return r.parseUrl(utlstr)
}

func (r *Request) Host(host string) *Request {
	r.host = host
	return r
}

func (r *Request) Path(path string) *Request {
	r.path = path
	return r
}

func (r *Request) Body(body io.ReadWriter) *Request {
	r.body = body
	return r
}

func (r *Request) buildRequest() *Request {
	urlpath := &url.URL{
		User: nil,
		Host: r.host,
	}
	switch r.method {
	case "GET":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Error("build request failed")
			panic("build request failed")
		}
		r.req = clientRequest
	case "POST":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Info("build request failed")
		}
		r.req = clientRequest
	case "DELETE":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Info("build request failed")
		}
		r.req = clientRequest
	case "PUT":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Info("build request failed")
		}
		r.req = clientRequest
	case "HEAD":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Info("build request failed")
		}
		r.req = clientRequest
	case "OPTIONS":
		clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
		if err != nil {
			log.Info("build request failed")
		}
		r.req = clientRequest
	default:
		panic("unsupport http method")
	}
	return r
}

func (r *Request) Get() *Request {
	r.method = "GET"
	return r.buildRequest()
}

func (r *Request) Post() *Request {
	r.method = "POST"
	return r.buildRequest()
}

func (r *Request) Put() *Request {
	r.method = "PUT"
	return r.buildRequest()
}

func (r *Request) Delete() *Request {
	r.method = "DELETE"
	return r.buildRequest()
}

func (r *Request) Head() *Request {
	r.method = "HEAD"
	return r.buildRequest()
}

func (r *Request) OPTIONS() *Request {
	r.method = "OPTIONS"
	return r.buildRequest()
}

func (r *Request) Timeout(timeout time.Duration) *Request {
	r.timeout = timeout
	return r
}

func (r *Request) Do() *Request {
	resp, err := r.HttpClient.Do(r.req)
	return r
}

func (r *Request) HandleResponse() {
	return
}

func (r *Request) Context(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	var newctx context.Context
	if r.timeout != 0 {
		newctx, _ = context.WithTimeout(ctx, r.timeout)
		return newctx
	}
	newctx = context.Background()
	return newctx
}
