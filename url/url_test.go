package url

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var hostTests = map[string]struct {
	in       string // URL.host field
	hostname string
	port     string // Expected port
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "foo.com:", hostname: "foo.com", port: ""},
	"without port":    {in: "foo.com", hostname: "foo.com", port: ""},
	"ip with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
	"ip without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
	"with empty host": {in: "", hostname: "", port: ""},
}

func BenchmarkURLString(b *testing.B) {
	var benchmarks = []*URL{
		{Scheme: "https"},
		{Scheme: "https", Host: "foo.com"},
		{Scheme: "https", Host: "foo.com", Path: "go"},
	}
	for _, u := range benchmarks {
		b.Run(u.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				u.String()
			}
		})
	}
}

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	want := &URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}

	got, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Parse(%q) mismatch (-want +got):\n%s", rawurl, diff)
	}
}

func TestParseInvalidURLs(t *testing.T) {
	tests := map[string]string{
		"missing scheme": "foo.com",
		"empty scheme":   "://foo.com",
	}
	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(in); err == nil {
				t.Errorf("Parse(%q)=nil; want an error", in)
			}
		})
	}
}

func TestParseScheme(t *testing.T) {
	const rawurl = "https://goo.com/go"
	scheme, _, ok := parseScheme(rawurl)
	if !ok {
		t.Fatalf("parseScheme(%q) invalid scheme", rawurl)
	}
	if got, want := scheme, "https"; got != want {
		t.Errorf("parseScheme(%q)=%q, want %q", rawurl, got, want)
	}
}

// Test that URL parses port correcly
func TestUrlHost(t *testing.T) {

	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("Port/%q/%q", name, tt.in), testPort(tt.in, tt.port))
	}

	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("HostName/%q/%q", name, tt.in), testHostName(tt.in, tt.hostname))
	}
}

func testPort(in, wantPort string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		u := &URL{Host: in}
		if got := u.Port(); got != wantPort {
			t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
		}
	}

}

func testHostName(in, wantHostName string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		u := &URL{Host: in}
		if got := u.HostName(); got != wantHostName {
			t.Errorf("for host %q; got %q; want %q", in, got, wantHostName)
		}
	}

}

func TestURLString(t *testing.T) {
	tests := map[string]struct {
		url  *URL
		want string
	}{
		"nil url":   {url: nil, want: ""},
		"empty url": {url: &URL{}, want: ""},
		"scheme":    {url: &URL{Scheme: "https"}, want: "https://"},
		"host": {
			url:  &URL{Scheme: "https", Host: "foo.com"},
			want: "https://foo.com",
		},
		"path": {
			url:  &URL{Scheme: "https", Host: "foo.com", Path: "go"},
			want: "https://foo.com/go",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if g, w := tt.url, tt.want; g.String() != w {
				t.Errorf("url: %#v\ngot: %q\nwant: %q", g, g, w)
			}
		})
	}
}
