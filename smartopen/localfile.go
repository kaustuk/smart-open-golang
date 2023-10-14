package smartopen

import (
	"io"
	"os"
)

func openLocalFile(Path string) (io.ReadCloser, error) {
	f, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
