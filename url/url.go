package url

import (
	"errors"
	"fmt"
	"strings"
)

// Parses rawurl into a URL structure
func Parse(rawurl string) (*URL, error) {

	i := strings.Index(rawurl, "://")

	if i < 0 {
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

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
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
