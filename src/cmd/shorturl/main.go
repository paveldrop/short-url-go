package main

import (
	"fmt"
	db "short-url-go/dboperations"
	g_url "short-url-go/geturl"
)

func main() {
	link, err := g_url.GetUrl()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(link)
	}
	linkToAdd := &db.Link{
		ShortURL: link,
		FullURL:  link,
	}
	db.AddLink(linkToAdd)
	db.PrintDB()
}
