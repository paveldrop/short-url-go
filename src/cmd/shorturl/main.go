package main

import (
	"fmt"
	db "short-url-go/dboperations"
	g_short "short-url-go/genshort"
	g_url "short-url-go/geturl"
	redir "short-url-go/redirect"
)

func main() {

	go redir.StartServer()
	go operateWithLink()

	select {}
}

func operateWithLink() {
	for {
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
}
