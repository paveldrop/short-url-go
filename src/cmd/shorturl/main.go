package main

import (
	"fmt"
	"log"
	"net/http"
	db "short-url-go/dboperations"
	g_short "short-url-go/genshort"
	g_url "short-url-go/geturl"
	redir "short-url-go/redirect"
)

func main() {
	_ = redir.RedirectFromShort()
	log.Println("Server started on port :8084")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
	}
	link, err := g_url.GetUrl()
	if err != nil {
		fmt.Println(err)
		return
	}

	shortLink, _ := g_short.ShortURL(link)
	linkToAdd := &db.Link{
		ShortURL: shortLink,
		FullURL:  link,
	}
	db.AddLink(linkToAdd)
	db.GetShortURl(link)

}
