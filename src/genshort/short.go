package genshort

import (
	"encoding/base64"
	"fmt"
	db "short-url-go/dboperations"
)

func ShortURL(url string) string {
	shortURL := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(shortURL[:6])
	result, err := db.ValidateShortInBD(shortURL[6:])
	if err != nil {
		fmt.Printf("error in valid short link: %v", err)
	}
	if result {
		return shortURL[:6]
	} else {
		fmt.Printf("gen new short link")
		return ""
	}
}
