package request

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"net/http/cookiejar"

	log "github.com/sirupsen/logrus"
)

type HandResponse interface {
	HandleResponse()
}

type Request struct {
	url      string
	host     string
	path     string
	scheme   string
	method   string
	headers  map[string]interface{}
	rawQuery string
	timeout  time.Duration

	totalRequest int
	parallel     int
	debug        bool

	contentType string
	body        io.ReadWriter

	req    *http.Request
	cookie map[string]*cookiejar.Jar

	HandResp HandResponse
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

func (r *Request) BuildRequest() *Request {
	urlpath := &url.URL{
		User: nil,
		Host: r.host,
	}
	clientRequest, err := http.NewRequest(r.method, urlpath.String(), r.body)
	if err != nil {
		log.Info("build request failed")
	}
	r.req = clientRequest
	return r
}

func (r *Request) Get() *Request {
	r.method = "GET"
	return r
}

func (r *Request) Post() *Request {
	r.method = "POST"
	return r
}

func (r *Request) Put() *Request {
	r.method = "PUT"
	return r
}

func (r *Request) Delete() *Request {
	r.method = "DELETE"
	return r
}

func (r *Request) Head() *Request {
	r.method = "HEAD"
	return r
}

func (r *Request) Patch() *Request {
	r.method = "PATCH"
	return r
}

func (r *Request) Timeout(timeout time.Duration) *Request {
	r.timeout = timeout
	return r
}

func (r *Request) Do() *Request {
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
