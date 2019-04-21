package random

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}

	return builder.String()
}
