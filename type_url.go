package generic

import (
	"database/sql/driver"
	"encoding/json"
	"net/url"
)

// URL is generic url.URL type structure
type URL struct {
	ValidFlag
	url *url.URL
}

// MarshalURL return generic.URL converting of request data
func MarshalURL(x interface{}) (URL, error) {
	v := URL{}
	err := v.Scan(x)
	return v, err
}

// Value implements the driver Valuer interface.
func (v URL) Value() (driver.Value, error) {
	if !v.Valid() || v.url == nil {
		return nil, nil
	}
	return v.url.String(), nil
}

// Scan implements the sql.Scanner interface.
func (v *URL) Scan(x interface{}) (err error) {
	v.url, v.ValidFlag, err = asURL(x)
	return
}

// Weak returns *url.URL, but if String.ValidFlag is false, returns nil.
func (v URL) Weak() interface{} {
	return v.URL()
}

// Set sets a specified value.
func (v *URL) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// String implements the Stringer interface.
func (v URL) String() string {
	if !v.Valid() || v.url == nil {
		return ""
	}
	return v.url.String()
}

// URL returns *url.URL, but if String.ValidFlag is false, returns nil.
func (v URL) URL() *url.URL {
	if !v.Valid() || v.url == nil {
		return nil
	}
	return v.url
}

// MarshalJSON implements the json.Marshaler interface.
func (v URL) MarshalJSON() ([]byte, error) {
	if !v.Valid() || v.url == nil {
		return nullBytes, nil
	}
	return json.Marshal(v.url.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *URL) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}

// EscapedPath returns the escaped form of v.url.Path.
// In general there are multiple possible escaped forms of any path.
//
// EscapedPath returns v.url.RawPath when it is a valid escaping of v.url.Path.
// Otherwise EscapedPath ignores v.url.RawPath and computes an escaped form on its own.
// The String and RequestURI methods use EscapedPath to construct their results.
// In general, code should call EscapedPath instead of reading v.url.RawPath directly.
func (v URL) EscapedPath() string {
	if !v.Valid() || v.url == nil {
		return ""
	}
	return v.url.EscapedPath()
}

// Hostname returns v.url.Host, without any port number.
//
// If Host is an IPv6 literal with a port number, Hostname returns the IPv6 literal without the square brackets.
// IPv6 literals may include a zone identifier.
func (v URL) Hostname() string {
	if !v.Valid() || v.url == nil {
		return ""
	}
	return v.url.Hostname()
}

// IsAbs reports whether the URL is absolute. Absolute means that it has a non-empty scheme.
func (v URL) IsAbs() bool {
	if !v.Valid() || v.url == nil {
		return false
	}
	return v.url.IsAbs()
}

// Port returns the port part of u.Host, without the leading colon.
// If u.Host doesn't contain a port, Port returns an empty string.
func (v URL) Port() string {
	if !v.Valid() || v.url == nil {
		return ""
	}
	return v.url.Port()
}

// Query parses RawQuery and returns the corresponding values.
// It silently discards malformed value pairs. To check errors use ParseQuery.
func (v URL) Query() url.Values {
	if !v.Valid() || v.url == nil {
		return url.Values{}
	}
	return v.url.Query()
}

// Parse parses a URL in the context of the receiver.
// The provided URL may be relative or absolute.
// Parse returns nil, err on parse failure, otherwise its return value is the same as ResolveReference.
func (v URL) Parse(ref string) (result URL, err error) {
	if v.url == nil {
		u := url.URL{}
		v.url, err = u.Parse(ref)
	} else {
		v.url, err = v.url.Parse(ref)
	}
	if err != nil {
		v = URL{}
		return v, err
	}
	v.ValidFlag = true
	return v, err
}

// RequestURI returns the encoded path?query or opaque?query string that would be used in an HTTP request for v.
func (v URL) RequestURI() string {
	if !v.Valid() || v.url == nil {
		return ""
	}
	return v.url.RequestURI()
}

// ResolveReference resolves a URI reference to an absolute URI from an absolute base URI, per RFC 3986 Section 5.2.
// The URI reference may be relative or absolute. ResolveReference always returns a new URL instance, even if the returned URL is identical to either the base or reference.
// If ref is an absolute URL, then ResolveReference ignores base and returns a copy of ref.
func (v URL) ResolveReference(ref *url.URL) URL {
	if ref == nil {
		v.ValidFlag = false
		v.url = nil
	} else {
		v.ValidFlag = true
		v.url = v.url.ResolveReference(ref)
	}
	return v
}
