package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GenerateReferenceNumber(prefix string) (string, error) {
	const max = 1000000

	value, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s-%s-%06d",
		prefix,
		time.Now().Format("20060102"),
		value.Int64(),
	), nil
}