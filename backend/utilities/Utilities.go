package utilities

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/mail"
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

//A simple function, short for the built in mail.parse, this function is redundant but it is comfortable to use i gues.
func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
