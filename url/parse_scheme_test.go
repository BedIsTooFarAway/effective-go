package url_test

import (
	"testing"

	"github.com/inancgumus/effective-go/ch04/url"
)

func TestParseScheme(t *testing.T) {
	const (
		rawurl     = "https://foo.com/go"
		wantScheme = "https"
		wantRest   = "foo.com/go"
		wantok     = true
	)

	scheme, rest, ok := url.ParseScheme(rawurl)
	if ok != wantok {
		t.Errorf("ParseSchema(%q) ok =  %q, want %q", rawurl, ok, wantok)
	}
	if scheme != wantScheme {

		t.Errorf("ParseSchema(%q) schema =  %q, want %q", rawurl, ok, wantScheme)
	}
	if rest != wantRest {

		t.Errorf("ParseSchema(%q) schema =  %q, want %q", rawurl, ok, wantRest)
	}
}
