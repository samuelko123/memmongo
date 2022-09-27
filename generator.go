package memongo

import (
	"crypto/rand"

	"github.com/pkg/errors"
)

func generateDBName(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.Wrapf(err, "cannot generate database name")
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}
