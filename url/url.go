package url

import (
	"errors"
	"fmt"
	"strings"
)

// Parses rawurl into a URL structure
func Parse(rawurl string) (*URL, error) {

	i := strings.Index(rawurl, "://")

	if i < 1 {
		return nil, errors.New("missing scheme")
	}
	scheme, rest := rawurl[:i], rawurl[i+3:]
	host, path := rest, ""

	if i := strings.Index(rest, "/"); i > 0 {
		host, path = rest[:i], rest[i+1:]
	}

	return &URL{scheme, host, path}, nil
	// return errors.New("malformed url")
}

func parseScheme(rawurl string) (scheme, rest string, ok bool) {
	return split(rawurl, "://", 1)
}

func split(s, sep string, n int) (a, b string, ok bool) {
	i := strings.Index(s, sep)
	if i < n {
		return "", "", false
	}
	return s[:i], s[i+len(sep):], true
}
func (u *URL) String() string {
	if u == nil {
		return ""
	}
	var s strings.Builder
	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteString("/")
		s.WriteString(p)
	}
	return s.String()
}

func (u *URL) testString() string {
	if u == nil {
		return ""
	}

	return fmt.Sprintf("scheme=%q, host=%q, path=%q", u.Scheme, u.Host, u.Path)
}

func (u *URL) HostName() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return u.Host
	}
	return u.Host[:i]
}

func (u *URL) Port() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return ""
	}
	return u.Host[i+1:]
}

// URL represents a parsed URL
type URL struct {
	// https://foo.com/go
	Scheme string // https
	Host   string // foo.com
	Path   string // go
}
