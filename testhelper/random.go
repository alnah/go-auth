package testhelper

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string of the specified length.
func RandomString(n uint) string {
	var stringBuilder strings.Builder
	k := len(alphabet)

	for i := uint(0); i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}

	return stringBuilder.String()
}
