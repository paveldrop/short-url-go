package main

import (
	"fmt"
	g_url "short-url-go/geturl"
)

func main() {
	link, err := g_url.GetUrl()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(link)
	}
}
