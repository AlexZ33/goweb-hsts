package goweb_hsts

import (
	"net/http"
	"time"
)

// isHTTPS
// params: acceptXForwardedProtoHeader  Decides whether to accept the X-Forwarded-Proto header as proof of SSL.
func isHTTPS(r *http.Request, acceptXForwardedProtoHeader bool) bool {
	// Added by common load balancers which do TLS offloading
	if acceptXForwardedProtoHeader && r.Header.Get("X-Forwarded-Proto") == "https" {
		return true
	}
	// If the X-Forwarded-Proto was set upstream as HTTP, then the request came in without TLS.
	if acceptXForwardedProtoHeader && r.Header.Get("X-Forwarded-Proto") == "http" {
		return false
	}

	// Set by some middleware
	if r.URL.Scheme == "https" {
		return true
	}
	// Set when the Go server is running HTTPS itself
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return true
	}
	return false
}

func createHeaderValueNew(maxAge time.Duration, sendPreloadDirective bool) string {
	
}
