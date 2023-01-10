package randstr

import (
	"math/rand"
	"time"
)

const (
	Alphanumeric        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Numeric      string = "0123456789"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Charset(charset string, length int) string {
	var b = make([]byte, length)

	random.Read(b)

	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}

	return string(b)
}
