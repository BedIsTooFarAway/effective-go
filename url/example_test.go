package url_test

import (
	"fmt"
	"log"

	"github.com/inancgumus/effective-go/ch04/url"
)

func ExampleURL() {
	u, err := url.Parse("http://foo.com:80/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Path = "nogo"

	fmt.Println(u)
	// Output:
	// https://foo.com:80/nogo
}

func ExampleURL_fields() {
	u, err := url.Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u)

	// Output:
	// https
	// foo.com
	// go
	// https://foo.com/go
}
