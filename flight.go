package goweb_hsts

import "net/http"

// A single request "flight"
type flight struct {
	rw  http.ResponseWriter
	req *IncomingRequest
	cfg handlerConfig
}

type handlerConfig struct {
	Handler Handler
}

// Result is the result of writing an HTTP response
// use ResponseWriter methods to obtain it

type Result struct {
}
