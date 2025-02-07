package redirect

import (
	// "net/http"

	"log"
	"net/http"
	db "short-url-go/dboperations"
)

func RedirectFromShort() error {
	links, err := db.GetAllShortLinks()
	if err != nil {
		return err
	}
	for i := 0; i < len(links); i++ {
		AddLink(links[i])
	}
	return nil

}

func AddLink(link db.Link) {
	http.HandleFunc("/"+link.ShortURL, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, link.FullURL, http.StatusFound)
	})
}

func StartServer() {
	_ = RedirectFromShort()
	log.Println("Server started on port :8084")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
	}
}

// func AddLinkToHandle(link db.Link) {
// 	http.HandleFunc("/"+link.ShortURL, func(w http.ResponseWriter, r *http.Request) {
// 		http.Redirect(w, r, link.FullURL, http.StatusFound)
// 	})
// }
