package main

import (
	"fmt"

	"github.com/kaustuk/smart-open-golang/smartopen"
)

func main() {
	fmt.Println("hello from main")

	f, err := smartopen.Open("test-file/testfile")

	// f, err := smartopen.Open("s3://test-kaustuk-bucket/testfile")

	if err != nil {
		fmt.Printf("error: %#v\n", err)
		return
	}

	b := make([]byte, 100)

	f.Read(b)

	fmt.Printf("value: \n%s", b)

	f.Close()
}
