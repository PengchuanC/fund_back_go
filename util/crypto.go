package util

import (
	"golang.org/x/crypto/scrypt"
)

const (
	Salt = "fund_back_go"
)

func Encode(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(Salt), 16384, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return string(dk)
}