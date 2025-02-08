package redirect

import (
	"html/template"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", loadTemplate)
	_ = RedirectFromShort()
	log.Println("Server started on port :8084")
	if err := http.ListenAndServe(":8084", mux); err != nil {
		log.Fatal(err)
	}
}

func loadTemplate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	template, err := template.ParseFiles("../../template/home.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
