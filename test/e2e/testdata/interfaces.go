package testdata

import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
}

func ReadData(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func CheckInterface(i interface{}) bool {
	return i != nil
}
