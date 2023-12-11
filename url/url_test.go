package url

import (
	"fmt"
	"testing"
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
}

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}
	// test Scheme
	if got, want := u.Scheme, "https"; want != got {
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawurl, got, want)
	}
	// test Host
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	}

	// test Path
	if got, want := u.Path, "go"; want != got {
		t.Errorf("Parse(%q).Path = %q; want %q", rawurl, got, want)
	}
}

func TestParseInvalidURLs(t *testing.T) {
	tests := map[string]string{
		"missing scheme": "foo.com",
	}
	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(in); err == nil {
				t.Errorf("Parse(%q)=nil; want an error", in)
			}
		})
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
	u := &URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}
	got, want := u.String(), "https://foo.com/go"
	if got != want {
		t.Errorf("%#v.String()\ngot %q\nwant %q", u, got, want)
	}
}
