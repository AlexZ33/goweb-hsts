package goweb_hsts

type ResponseWriter interface {
	ResponseHeadersWriter
}

type ResponseHeadersWriter interface {
	// Header returns the collection of headers that will be set
	// on the response. Headers must be set before writing a
	// response (e.g. Write, WriteTemplate).
	Header() Header
}
