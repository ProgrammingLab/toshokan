package util

import (
	rand "crypto/rand"
	"math/big"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString generates random string
func GenerateRandomString(length int) (string, error) {
	res := make([]byte, length)

	max := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		res[i] = chars[n.Int64()]
	}

	return string(res), nil
}
