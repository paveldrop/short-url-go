package genshort

import (
	"crypto/rand"
	"math/big"
)

const (
	length   int    = 6
	alphabet string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ShortURL(url string) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(62))
		if err != nil {
			return "", err
		}
		index := int(random.Int64())
		result[i] = alphabet[index]
	}
	// fmt.Println(string(result))
	return string(result), nil
}
