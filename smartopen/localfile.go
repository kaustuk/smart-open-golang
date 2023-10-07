package smartopen

import (
	"io"
	"os"
)

func openLocalFile(Path string) (io.Reader, error) {
	f, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	return io.Reader(f), nil
}
