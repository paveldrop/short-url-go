package redirect

import (
	// "net/http"

	"net/http"
	db "short-url-go/dboperations"
)

func RedirectFromShort() error {
	// http.HandleFunc("")
	// links :=
	links, err := db.GetAllShortLinks()
	if err != nil {
		return err
	}
	for i := 0; i < len(links); i++ {
		startAllLinks(links[i])
	}
	return nil

}

func startAllLinks(link db.Link) {
	http.HandleFunc("/"+link.ShortURL, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, link.FullURL, http.StatusFound)
	})
}

/*
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ya.ru", http.StatusFound)
})

log.Println("Запуск сервера на порту :8080")
if err := http.ListenAndServe(":8080", nil); err != nil {
	log.Fatal(err)
}
*/
