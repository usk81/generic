package generic

import (
	"net/url"
	"reflect"
	"testing"
)

type testURLStruct struct {
	URL       URL `json:"url"`
	PATH      URL `json:"path"`
	NullValue URL `json:"null_value"`
}

type testURLResult struct {
	in         string
	uri        URL
	hostname   string
	abs        bool
	port       string
	queries    url.Values
	requestURI string
}

const (
	testURLString = "https://foobar.test:8080?q=fizzbizz"
	rootPath      = "/"
)

var (
	urltests = []testURLResult{
		// no path
		testURLResult{
			in: "http://www.google.com",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "www.google.com",
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// path
		testURLResult{
			in: "http://www.google.com/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "www.google.com",
					Path:   rootPath,
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// path with hex escaping
		testURLResult{
			in: "http://www.google.com/file%20one%26two",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:  "http",
					Host:    "www.google.com",
					Path:    "/file one&two",
					RawPath: "/file%20one%26two",
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/file%20one%26two",
		},
		// user
		testURLResult{
			in: "ftp://webmaster@www.google.com/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "ftp",
					User:   url.User("webmaster"),
					Host:   "www.google.com",
					Path:   rootPath,
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// escape sequence in username
		testURLResult{
			in: "ftp://john%20doe@www.google.com/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "ftp",
					User:   url.User("john doe"),
					Host:   "www.google.com",
					Path:   rootPath,
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// empty query
		testURLResult{
			in: "http://www.google.com/?",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:     "http",
					Host:       "www.google.com",
					Path:       "/",
					ForceQuery: true,
				},
			},
			hostname:   "www.google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/?",
		},
		// query ending in question mark
		// golang.org/issue/14573
		testURLResult{
			in: "http://www.google.com/?foo=bar?",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/",
					RawQuery: "foo=bar?",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"foo": []string{"bar?"},
			},
			requestURI: "/?foo=bar?",
		},
		// query
		testURLResult{
			in: "http://www.google.com/?q=go+language",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/",
					RawQuery: "q=go+language",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "/?q=go+language",
		},
		// query with hex escaping: NOT parsed
		testURLResult{
			in: "http://www.google.com/?q=go%20language",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/",
					RawQuery: "q=go%20language",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "/?q=go%20language",
		},
		// %20 outside query
		testURLResult{
			in: "http://www.google.com/a%20b?q=c+d",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/a b",
					RawQuery: "q=c+d",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"c d"},
			},
			requestURI: "/a%20b?q=c+d",
		},
		// path without leading /, so no parsing
		testURLResult{
			in: "http:www.google.com/?q=go+language",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Opaque:   "www.google.com/",
					RawQuery: "q=go+language",
				},
			},
			hostname: "",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "www.google.com/?q=go+language",
		},
		// path without leading /, so no parsing
		testURLResult{
			in: "http:%2f%2fwww.google.com/?q=go+language",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Opaque:   "%2f%2fwww.google.com/",
					RawQuery: "q=go+language",
				},
			},
			hostname: "",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "%2f%2fwww.google.com/?q=go+language",
		},
		// non-authority with path
		testURLResult{
			in: "mailto:/webmaster@golang.org",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "mailto",
					Path:   "/webmaster@golang.org",
				},
			},
			hostname:   "",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/webmaster@golang.org",
		},
		// non-authority
		testURLResult{
			in: "mailto:webmaster@golang.org",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "mailto",
					Opaque: "webmaster@golang.org",
				},
			},
			hostname:   "",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "webmaster@golang.org",
		},
		// unescaped :// in query should not create a scheme
		testURLResult{
			in: "/foo?query=http://bad",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Path:     "/foo",
					RawQuery: "query=http://bad",
				},
			},
			hostname: "",
			abs:      false,
			port:     "",
			queries: url.Values{
				"query": []string{"http://bad"},
			},
			requestURI: "/foo?query=http://bad",
		},
		// leading // without scheme should create an authority
		testURLResult{
			in: "//foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Host: "foo",
				},
			},
			hostname:   "foo",
			abs:        false,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// leading // without scheme, with userinfo, path, and query
		testURLResult{
			in: "//user@foo/path?a=b",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					User:     url.User("user"),
					Host:     "foo",
					Path:     "/path",
					RawQuery: "a=b",
				},
			},
			hostname: "foo",
			abs:      false,
			port:     "",
			queries: url.Values{
				"a": []string{"b"},
			},
			requestURI: "/path?a=b",
		},
		// Three leading slashes isn't an authority, but doesn't return an error.
		// (We can't return an error, as this code is also used via
		// ServeHTTP -> ReadRequest -> Parse, which is arguably a
		// different URL parsing context, but currently shares the
		// same codepath)
		testURLResult{
			in: "///threeslashes",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Path: "///threeslashes",
				},
			},
			hostname:   "",
			abs:        false,
			port:       "",
			queries:    url.Values{},
			requestURI: "///threeslashes",
		},
		testURLResult{
			in: "http://user:password@google.com",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					User:   url.UserPassword("user", "password"),
					Host:   "google.com",
				},
			},
			hostname:   "google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// unescaped @ in username should not confuse host
		testURLResult{
			in: "http://j@ne:password@google.com",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					User:   url.UserPassword("j@ne", "password"),
					Host:   "google.com",
				},
			},
			hostname:   "google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// unescaped @ in password should not confuse host
		testURLResult{
			in: "http://jane:p@ssword@google.com",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					User:   url.UserPassword("jane", "p@ssword"),
					Host:   "google.com",
				},
			},
			hostname:   "google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		testURLResult{
			in: "http://j@ne:password@google.com/p@th?q=@go",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					User:     url.UserPassword("j@ne", "password"),
					Host:     "google.com",
					Path:     "/p@th",
					RawQuery: "q=@go",
				},
			},
			hostname: "google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"@go"},
			},
			requestURI: "/p@th?q=@go",
		},
		testURLResult{
			in: "http://www.google.com/?q=go+language#foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/",
					RawQuery: "q=go+language",
					Fragment: "foo",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "/?q=go+language",
		},
		testURLResult{
			in: "http://www.google.com/?q=go+language#foo%26bar",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "www.google.com",
					Path:     "/",
					RawQuery: "q=go+language",
					Fragment: "foo&bar",
				},
			},
			hostname: "www.google.com",
			abs:      true,
			port:     "",
			queries: url.Values{
				"q": []string{"go language"},
			},
			requestURI: "/?q=go+language",
		},
		testURLResult{
			in: "file:///home/adg/rabbits",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "file",
					Host:   "",
					Path:   "/home/adg/rabbits",
				},
			},
			hostname:   "",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/home/adg/rabbits",
		},
		// "Windows" paths are no exception to the rule.
		// See golang.org/issue/6027, especially comment #9.
		testURLResult{
			in: "file:///C:/FooBar/Baz.txt",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "file",
					Host:   "",
					Path:   "/C:/FooBar/Baz.txt",
				},
			},
			hostname:   "",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/C:/FooBar/Baz.txt",
		},
		// case-insensitive scheme
		testURLResult{
			in: "MaIlTo:webmaster@golang.org",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "mailto",
					Opaque: "webmaster@golang.org",
				},
			},
			hostname:   "",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "webmaster@golang.org",
		},
		// Relative path
		testURLResult{
			in: "a/b/c",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Path: "/a/b/c",
				},
			},
			hostname:   "",
			abs:        false,
			port:       "",
			queries:    url.Values{},
			requestURI: "/a/b/c",
		},
		// escaped '?' in username and password
		testURLResult{
			in: "http://%3Fam:pa%3Fsword@google.com",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					User:   url.UserPassword("?am", "pa?sword"),
					Host:   "google.com",
				},
			},
			hostname:   "google.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host subcomponent; IPv4 address in RFC 3986
		testURLResult{
			in: "http://192.168.0.1/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "192.168.0.1",
					Path:   "/",
				},
			},
			hostname:   "192.168.0.1",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host and port subcomponents; IPv4 address in RFC 3986
		testURLResult{
			in: "http://192.168.0.1:8080/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "192.168.0.1:8080",
					Path:   "/",
				},
			},
			hostname:   "192.168.0.1",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host subcomponent; IPv6 address in RFC 3986
		testURLResult{
			in: "http://[fe80::1]/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1]",
					Path:   "/",
				},
			},
			hostname:   "fe80::1",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host and port subcomponents; IPv6 address in RFC 3986
		testURLResult{
			in: "http://[fe80::1]:8080/",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1]:8080",
					Path:   "/",
				},
			},
			hostname:   "fe80::1",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host subcomponent; IPv6 address with zone identifier in RFC 6874
		testURLResult{
			in: "http://[fe80::1%25en0]/", // alphanum zone identifier
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1%en0]",
					Path:   "/",
				},
			},
			hostname:   "fe80::1%en0",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host and port subcomponents; IPv6 address with zone identifier in RFC 6874
		testURLResult{
			in: "http://[fe80::1%25en0]:8080/", // alphanum zone identifier
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1%en0]:8080",
					Path:   "/",
				},
			},
			hostname:   "fe80::1%en0",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host subcomponent; IPv6 address with zone identifier in RFC 6874
		testURLResult{
			in: "http://[fe80::1%25%65%6e%301-._~]/", // percent-encoded+unreserved zone identifier
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1%en01-._~]",
					Path:   "/",
				},
			},
			hostname:   "fe80::1%en01-._~",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// host and port subcomponents; IPv6 address with zone identifier in RFC 6874
		testURLResult{
			in: "http://[fe80::1%25%65%6e%301-._~]:8080/", // percent-encoded+unreserved zone identifier
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[fe80::1%en01-._~]:8080",
					Path:   "/",
				},
			},
			hostname:   "fe80::1%en01-._~",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		// alternate escapings of path survive round trip
		testURLResult{
			in: "http://rest.rsc.io/foo%2fbar/baz%2Fquux?alt=media",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "http",
					Host:     "rest.rsc.io",
					Path:     "/foo/bar/baz/quux",
					RawPath:  "/foo%2fbar/baz%2Fquux",
					RawQuery: "alt=media",
				},
			},
			hostname: "rest.rsc.io",
			abs:      true,
			port:     "",
			queries: url.Values{
				"alt": []string{"media"},
			},
			requestURI: "/foo%2fbar/baz%2Fquux?alt=media",
		},
		// issue 12036
		testURLResult{
			in: "mysql://a,b,c/bar",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "mysql",
					Host:   "a,b,c",
					Path:   "/bar",
				},
			},
			hostname:   "a,b,c",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/bar",
		},
		// worst case host, still round trips
		testURLResult{
			in: "scheme://!$&'()*+,;=hello!:port/path",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "scheme",
					Host:   "!$&'()*+,;=hello!:port",
					Path:   "/path",
				},
			},
			hostname:   "!$&'()*+,;=hello!",
			abs:        true,
			port:       "port",
			queries:    url.Values{},
			requestURI: "/path",
		},
		// worst case path, still round trips
		testURLResult{
			in: "http://host/!$&'()*+,;=:@[hello]",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:  "http",
					Host:    "host",
					Path:    "/!$&'()*+,;=:@[hello]",
					RawPath: "/!$&'()*+,;=:@[hello]",
				},
			},
			hostname:   "host",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/!$&'()*+,;=:@[hello]",
		},
		// golang.org/issue/5684
		testURLResult{
			in: "http://example.com/oid/[order_id]",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:  "http",
					Host:    "example.com",
					Path:    "/oid/[order_id]",
					RawPath: "/oid/[order_id]",
				},
			},
			hostname:   "example.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/oid/[order_id]",
		},
		// golang.org/issue/12200 (colon with empty port)
		testURLResult{
			in: "http://192.168.0.2:8080/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "192.168.0.2:8080",
					Path:   "/foo",
				},
			},
			hostname:   "192.168.0.2",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		testURLResult{
			in: "http://192.168.0.2:/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "192.168.0.2:",
					Path:   "/foo",
				},
			},
			hostname:   "192.168.0.2",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		// Malformed IPv6 but still accepted.
		testURLResult{
			in: "http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:8080/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac:8080",
					Path:   "/foo",
				},
			},
			hostname:   "2b01",
			abs:        true,
			port:       "e34:ef40:7730:8e70:5aff:fefe:edac:8080",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		// Malformed IPv6 but still accepted.
		testURLResult{
			in: "http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac:",
					Path:   "/foo",
				},
			},
			hostname:   "2b01",
			abs:        true,
			port:       "e34:ef40:7730:8e70:5aff:fefe:edac:",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		testURLResult{
			in: "http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:8080/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:8080",
					Path:   "/foo",
				},
			},
			hostname:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
			abs:        true,
			port:       "8080",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		testURLResult{
			in: "http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:",
					Path:   "/foo",
				},
			},
			hostname:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		// golang.org/issue/7991 and golang.org/issue/12719 (non-ascii %-encoded in host)
		testURLResult{
			in: "http://hello.世界.com/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "hello.世界.com",
					Path:   "/foo",
				},
			},
			hostname:   "hello.世界.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		testURLResult{
			in: "http://hello.%e4%b8%96%e7%95%8c.com/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "hello.世界.com",
					Path:   "/foo",
				},
			},
			hostname:   "hello.世界.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		testURLResult{
			in: "http://hello.%E4%B8%96%E7%95%8C.com/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "hello.世界.com",
					Path:   "/foo",
				},
			},
			hostname:   "hello.世界.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		// golang.org/issue/10433 (path beginning with //)
		testURLResult{
			in: "http://example.com//foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "//foo",
				},
			},
			hostname:   "example.com",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "//foo",
		},
		// test that we can reparse the host names we accept.
		testURLResult{
			in: "myscheme://authority<\"hi\">/foo",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "myscheme",
					Host:   "authority<\"hi\">",
					Path:   "/foo",
				},
			},
			hostname:   "authority<\"hi\">",
			abs:        true,
			port:       "",
			queries:    url.Values{},
			requestURI: "/foo",
		},
		// spaces in hosts are disallowed but escaped spaces in IPv6 scope IDs are grudgingly OK.
		// This happens on Windows.
		// golang.org/issue/14002
		testURLResult{
			in: "tcp://[2020::2020:20:2020:2020%25Windows%20Loves%20Spaces]:2020",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme: "tcp",
					Host:   "[2020::2020:20:2020:2020%Windows Loves Spaces]:2020",
				},
			},
			hostname:   "2020::2020:20:2020:2020%Windows Loves Spaces",
			abs:        true,
			port:       "2020",
			queries:    url.Values{},
			requestURI: rootPath,
		},
		testURLResult{
			in: "magnet:?xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "magnet",
					Host:     "",
					Path:     "",
					RawQuery: "xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
				},
			},
			hostname: "",
			abs:      true,
			port:     "",
			queries: url.Values{
				"xt": []string{"urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a"},
				"dn": []string{""},
			},
			requestURI: "/?xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
		},
		testURLResult{
			in: "mailto:?subject=hi",
			uri: URL{
				ValidFlag: true,
				url: &url.URL{
					Scheme:   "mailto",
					Host:     "",
					Path:     "",
					RawQuery: "subject=hi",
				},
			},
			hostname: "",
			abs:      true,
			port:     "",
			queries: url.Values{
				"subject": []string{"hi"},
			},
			requestURI: "/?subject=hi",
		},
	}

	// testURLResult{
	// 			in: "validFlag false",
	// 			uri: URL{
	// 				ValidFlag: false,
	// 				url: &url.URL{
	// 					Scheme:   "http",
	// 					Host:     "www.google.com",
	// 					Path:     "/",
	// 					RawQuery: "q=go+language",
	// 					Fragment: "foo",
	// 				},
	// 			},
	// 			hostname:   "",
	// 			abs:        false,
	// 			port:       "",
	// 			queries:    url.Values{},
	// 			requestURI: "",
	// 		},
	// 		testURLResult{
	// 			in: "nil",
	// 			uri: URL{
	// 				ValidFlag: true,
	// 				url:       nil,
	// 			},
	// 			hostname:   "",
	// 			abs:        false,
	// 			port:       "",
	// 			queries:    url.Values{},
	// 			requestURI: "",
	// 		},
)

func TestMarshalURL(t *testing.T) {
	v, _ := url.Parse(testURLString)

	expected := URL{
		ValidFlag: true,
		url:       v,
	}

	tests := []struct {
		name    string
		args    interface{}
		want    URL
		wantErr bool
	}{
		{
			name:    "URL",
			args:    v,
			want:    expected,
			wantErr: false,
		},
		{
			name:    "string",
			args:    testURLString,
			want:    expected,
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    10000,
			want:    URL{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalURL(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL_Value(t *testing.T) {
	t.Skip()
}

func TestURL_Scan(t *testing.T) {
	t.Skip()
}

func TestURL_Weak(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.Weak(); got != tt.uri.url {
				t.Errorf("URL.URL() = %v, want %v", got, tt.uri.url)
			}
		})
	}
}

// func TestURL_Set(t *testing.T) {
// 	t.Skip()
// }

func TestURL_String(t *testing.T) {
	u1, _ := url.Parse(testURLString)

	s := "/foobar?fizz=bizz"
	u2, _ := url.Parse(s)

	tests := []struct {
		name   string
		fields URL
		want   string
	}{
		{
			name: "valid url",
			fields: URL{
				ValidFlag: true,
				url:       u1,
			},
			want: testURLString,
		},
		{
			name: "ValidFlag false",
			fields: URL{
				ValidFlag: false,
				url:       u1,
			},
			want: "",
		},
		{
			name: "only path",
			fields: URL{
				ValidFlag: true,
				url:       u2,
			},
			want: s,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.String(); got != tt.want {
				t.Errorf("URL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL_URL(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.URL(); got != tt.uri.url {
				t.Errorf("URL.URL() = %v, want %v", got, tt.uri.url)
			}
		})
	}
}

func TestURL_MarshalJSON_UnmarshalJSON(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			bs, err := tt.uri.MarshalJSON()
			if err != nil {
				t.Errorf("URL.MarshalJSON() error = %v, not expected", err)
				return
			}
			v := URL{}
			err = v.UnmarshalJSON(bs)
			if err != nil {
				t.Errorf("URL.UnmarshalJSON() error = %v, not expected", err)
				return
			}
			if v.String() != tt.uri.String() {
				t.Errorf("URL.UnmarshalJSON() = %v, want %v", v, tt.uri)
			}
		})
	}
}

func TestURL_UnmarshalJSON(t *testing.T) {
	t.Skip()
}

func TestURL_EscapedPath(t *testing.T) {
	t.Skip()
}

func TestURL_Hostname(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.Hostname(); got != tt.hostname {
				t.Errorf("URL.Hostname() = %v, want %v", got, tt.hostname)
			}
		})
	}
}

func TestURL_IsAbs(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.IsAbs(); got != tt.abs {
				t.Errorf("URL.IsAbs() = %v, want %v", got, tt.abs)
			}
		})
	}
}

func TestURL_Port(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.Port(); got != tt.port {
				t.Errorf("URL.Port() = %v, want %v", got, tt.port)
			}
		})
	}
}

func TestURL_Query(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.Query(); !reflect.DeepEqual(got, tt.queries) {
				t.Errorf("URL.Query() = %#v, want %#v", got, tt.queries)
			}
		})
	}
}

func TestURL_Parse(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			v := URL{}
			gotResult, err := v.Parse(tt.in)
			if err != nil {
				t.Errorf("URL.Parse() error = %v, not expected", err)
				if gotResult != v {
					t.Errorf("URL.Parse() = %v, want %v", gotResult, URL{})
				}
				return
			}
			if !reflect.DeepEqual(gotResult, tt.uri) {
				t.Errorf("URL.Parse() = %v, want %v", gotResult, tt.uri)
			}
		})
	}
}

func TestURL_RequestURI(t *testing.T) {
	for _, tt := range urltests {
		t.Run(tt.in, func(t *testing.T) {
			if got := tt.uri.RequestURI(); got != tt.requestURI {
				t.Errorf("URL.RequestURI() = %v, want %v", got, tt.requestURI)
			}
		})
	}
}

func TestURL_ResolveReference(t *testing.T) {
	u, _ := url.Parse(testURLString)
	ud, _ := url.Parse("https://notfound.test?q1=a&q2=b")
	v := URL{
		ValidFlag: false,
		url:       ud,
	}
	tests := []struct {
		name string
		args *url.URL
		want URL
	}{
		{
			name: "valid URL",
			args: u,
			want: URL{
				ValidFlag: true,
				url:       u,
			},
		},
		{
			name: "nil",
			args: nil,
			want: URL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := v.ResolveReference(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL.ResolveReference() = %v, want %v", got, tt.want)
			}
		})
	}
}
