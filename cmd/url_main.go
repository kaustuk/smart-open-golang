package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("hello from main")

	// a, _ := url.Parse("s3://test-kaustuk-bucket/testfile")

	a, _ := url.Parse("/kaustuk/a")

	fmt.Printf("%s\n", a.Scheme)
	fmt.Printf("%s\n", a.Host)
	fmt.Printf("%#v\n", a.Path)
}
