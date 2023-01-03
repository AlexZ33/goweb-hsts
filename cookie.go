package goweb_hsts

import (
	"net/http"
	"sync"
)

// A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
// HTTP response or the Cookie header of an HTTP request.
// See https://tools.ietf.org/html/rfc6265 for details.
type Cookie struct {
	wrapped *http.Cookie
}

// NewCookie creates a new Cookie with safe default settings.
// Those safe defaults are:
// - Secure: true (if the framework is not in dev mode)
// - HttpOnly: true
// - SameSite: Lax
// For more info about all the options, see:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
func NewCookie(name, value string) *Cookie {
	var devMu sync.RWMutex
	var isLocalDev bool
	devMu.RLock()
	defer devMu.RUnlock()
	return &Cookie{
		&http.Cookie{
			Name:     name,
			Value:    value,
			Secure:   !isLocalDev,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
	}
}

func (c *Cookie) String() string {
	return c.wrapped.String()
}
