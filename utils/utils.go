package utils

import (
	"fmt"
	"crypto/md5"
)

func Md5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
