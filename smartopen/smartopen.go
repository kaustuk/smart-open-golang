package smartopen

import (
	"fmt"
	"io"
	"net/url"
)

func Open(path string) (io.ReadCloser, error) {

	u, err := url.Parse(path)

	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "s3":
		return smartOpenS3Obj.openFile(u.Host, u.Path[1:])
	case "":
		return openLocalFile(path)
	default:
		fmt.Println("unknown object time")
		return nil, fmt.Errorf("unknown file path passed")
	}

	// return smartOpenS3Obj.openFile("test-kaustuk-bucket", "testfile")
	// return openLocalFile(path)
}
