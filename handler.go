package goweb_hsts

// Handler responds to an HTTP request.
type Handler interface {
	ServerHTTP(writer ResponseWriter, request *IncomingRequest) Result
}

// HandlerFunc is used to convert a function into a Handler.
type HandlerFunc func(ResponseWriter, *IncomingRequest) Result

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *IncomingRequest) Result {
	return f(w, r)
}
