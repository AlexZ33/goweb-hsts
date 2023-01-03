package goweb_hsts

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type IncomingRequest struct {
	Header Header
	// TLS is set just like this TLS field of the net/http.Request. For more information
	// see https://pkg.go.dev/net/http?tab=doc#Request.
	TLS *tls.ConnectionState
	req *http.Request

	postParseOnce      *sync.Once
	multipartParseOnce *sync.Once
}

//func NewIncomingRequest(req *http.Request) *IncomingRequest {
//	if req == nil {
//		return nil
//	}
//	req = req.WithContext(context.WithValue(req.Context(),
//		flight))
//}

func (r *IncomingRequest) WithStrippedURLPrefix(prefix string) (*IncomingRequest, error) {
	req := rawRequest(r)
	if !strings.HasPrefix(req.URL.Path, prefix) {
		return nil, fmt.Errorf("Path %q doesn't have prefix %q", req.URL.Path, prefix)
	}
	if req.URL.RawPath != "" && !strings.HasPrefix(req.URL.RawPath, prefix) {
		return nil, fmt.Errorf("RawPath %q doesn't have prefix %q", req.URL.RawPath, prefix)
	}
	req2 := new(http.Request)
	*req2 = *req
	req2.URL = new(url.URL)
	*req2.URL = *req.URL
	req2.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
	req2.URL.RawPath = strings.TrimPrefix(req.URL.RawPath, prefix)

	r2 := new(IncomingRequest)
	*r2 = *r
	r2.req = req2

	return r2, nil
}

func rawRequest(r *IncomingRequest) *http.Request {
	return r.req
}
