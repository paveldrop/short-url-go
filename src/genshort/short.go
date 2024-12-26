package genshort

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	db "short-url-go/dboperations"
	"strconv"
)

const (
	sugar string = "sUgAr"
)

func ShortURL(url string) string {

	shortURL := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(shortURL[:6])
	randomNum := strconv.Itoa(rand.Int())
	result, err := db.ValidateShortInBD(shortURL[:6])
	if err != nil {
		fmt.Printf("error in valid short link: %v", err)
	}
	if result {
		return shortURL[:6]
	} else {
		shortURL = base64.StdEncoding.EncodeToString([]byte(shortURL + randomNum))
		fmt.Printf("gen new short link\n%s", shortURL[:6])
		return shortURL[:6]
	}
}
