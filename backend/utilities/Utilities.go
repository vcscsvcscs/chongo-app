package utilities

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

//Checks if file on path exists or not
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

//Md 5 string hasher for quick use, this is not safe for passwords because md5, but it is absolutely fine for session tokens.
func Md(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
