package request

import (
	"context"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type Request struct {
	url      string
	host     string
	path     string
	scheme   string
	method   string
	headers  map[string][]string
	rawQuery string
	timeout  int64

	totalRequest int
	parallel     int
	debug        bool

	bodyType string
	body     io.ReadWriter

	req *http.Request
}

func (r *Request) Scheme(scheme string) *Request {
	r.scheme = scheme
	return r
}

func (r *Request) BodyType(bodyType string) *Request {
	r.bodyType = bodyType
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

func (r *Request) ParseUrl(urlstr string) *Request {
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

func (r *Request) Build() *Request {
	client := http.NewRequest
	return r
}

func (r *Request) Timeout() *Request {
	return r
}

func (r *Request) Do() *Request {
	return r
}

func (r *Request) HandleResponse() *Request {
	return r
}

func (r *Request) Context() context.Context {
	return nil
}
