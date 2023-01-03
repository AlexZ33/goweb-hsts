package goweb_hsts

import (
	"errors"
	"fmt"
	"net/http"
	"net/textproto"
)

// Header represents the key-value pairs in an HTTP header.
// The keys will be in canonical form, as returned by textproto.CanonicalMIMEHeaderKey.
type Header struct {
	wrapped http.Header
	claimed map[string]bool
}

// NewHeader creates a new Header.
func NewHeader(h http.Header) Header {
	if h == nil {
		h = http.Header{}
	}
	return Header{
		wrapped: h,
		claimed: map[string]bool{},
	}
}

func (h Header) Claim(name string) (set func([]string)) {
	name = textproto.CanonicalMIMEHeaderKey(name)
	if err := h.writableHeader(name); err != nil {
		panic(err)
	}
	h.claimed[name] = true
	return func(v []string) {
		if v == nil {
			return
		}
		h.wrapped[name] = v
	}
}

// IsClaimed reports whether the provided header is already claimed. The name is
// first canonicalized using textproto.CanonicalMIMEHeaderKey. The Set-Cookie header
// is treated as claimed.
func (h Header) IsClaimed(name string) bool {
	name = textproto.CanonicalMIMEHeaderKey(name)
	err := h.writableHeader(name)
	return err != nil
}

// writableHeader assumes that the given name already has been canonicalized
// using textproto.CanonicalMIMEHeaderKey.
func (h Header) writableHeader(name string) error {
	if name == "Set-Cookie" {
		return errors.New("can't write to Set-Cookie header")
	}
	if h.claimed[name] {
		return fmt.Errorf("claimed header: %s", name)
	}
	return nil
}

// Set sets the header with the given name to the given value.
// The name is first canonicalized using textproto.CanonicalMIMEHeaderKey.
// This method first removes all other values associated with this
// header before setting the new value. It panics when applied on claimed headers
// or on the Set-Cookie header.

func (h Header) Set(name, value string) {
	name = textproto.CanonicalMIMEHeaderKey(name)
	if err := h.writableHeader(name); err != nil {
		panic(err)
	}
	h.wrapped.Set(name, value)
}

// Add adds a new header with the given name and the given value to
// the collection of headers. The name is first cannonicalized using textproto.CanonicalMIMEHeaderKey
// It panics when applied on claimed headers or on the Set-Cookie header.
func (h Header) Add(name, value string) {
	name = textproto.CanonicalMIMEHeaderKey(name)
	if err := h.writableHeader(name); err != nil {
		panic(err)
	}
	h.wrapped.Add(name, value)
}

// Del deletes all headers with the given name. The name is first canonicalized
// using textproto.CanonicalMIMEHeaderKey. It panics when applied on claimed headers
// or on the Set-Cookie header.
func (h Header) Del(name string) {
	name = textproto.CanonicalMIMEHeaderKey(name)
	if err := h.writableHeader(name); err != nil {
		panic(err)
	}
	h.wrapped.Del(name)
}

// Get returns the value of the first header with the given name.
// The name is first canonicalized using textproto.CanonicalMIMEHeaderKey.
// If no header exists with the given name then "" is returned.
func (h Header) Get(name string) string {
	return h.wrapped.Get(name)
}

// Values returns all the values of all the headers with the given name.
// The name is first canonicalized using textproto.CanonicalMIMEHeaderKey.
// The values are returned in the same order as they were sent in the request.
// The values are returned as a copy of the original slice of strings in
// the internal header map. This is to prevent modification of the original
// slice. If no header exists with the given name then an empty slice is
// returned.
func (h Header) Values(name string) []string {
	v := h.wrapped.Values(name)
	clone := make([]string, len(v))
	copy(clone, v)
	return clone
}

// addCookie adds the cookie provided as a Set-Cookie header in the header
// collection. If the cookie is nil or cookie.Name() is invalid, no header is
// added and an error is returned. This is the only method that can modify the
// Set-Cookie header. If other methods try to modify the header they will return
// errors.
func (h Header) addCookie(c *Cookie) error {
	v := c.String()
	if v == "" {
		return errors.New("invalid cookie name")
	}
	h.wrapped.Add("Set-Cookie", v)
	return nil
}
