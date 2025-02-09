package redirect

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	db "short-url-go/dboperations"
)

type PageData struct {
	ShortURL string
}

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
	mux.HandleFunc("/create-link", createLinkHandler)
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
	var pageData PageData
	if r.Method == "GET" && r.URL.Query().Get("short") != "" {
		pageData.ShortURL = r.URL.Query().Get("short")
	}
	template, err := template.ParseFiles("../../template/home.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = template.Execute(w, pageData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func createLinkHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()

	name := r.FormValue("name")
	if name == "" {
		fmt.Fprint(w, "Please provide a valid URL.")
		return
	}
	pageData := PageData{
		ShortURL: "test", // reform to generate link
	}

	// log.Printf("Received URL: %s\n", name)

	// shortURL := generateShortURL(name)

	// save shortlink in DB

	// add redirect to shortlink

	template, err := template.ParseFiles("../../template/home.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = template.Execute(w, pageData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// fmt.Fprintf(w, "Your short URL is: %s", "shortURL")
}
