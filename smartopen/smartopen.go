package smartopen

import "io"

func Open(path string) (io.Reader, error) {
	return openLocalFile(path)
}
