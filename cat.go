package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//go:embed cat1.gif
var cat1GIF []byte

//go:embed cat2.gif
var cat2GIF []byte

//go:embed cat.html
var catHTML string

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/cat/1.html", http.StatusFound)
	})
	http.HandleFunc("/cat/1.html", func(w http.ResponseWriter, r *http.Request) {
		html := strings.ReplaceAll(catHTML, "/cat/X.gif", "/cat/1.gif")
		html = strings.ReplaceAll(html, "/cat/X.html", "/cat/2.html")
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(html)))
		w.Write([]byte(html))
	})
	http.HandleFunc("/cat/2.html", func(w http.ResponseWriter, r *http.Request) {
		html := strings.ReplaceAll(catHTML, "/cat/X.gif", "/cat/2.gif")
		html = strings.ReplaceAll(html, "/cat/X.html", "/cat/1.html")
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(html)))
		w.Write([]byte(html))
	})
	http.HandleFunc("/cat/1.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		w.Header().Set("Content-Length", strconv.Itoa(len(cat1GIF)))
		w.Write(cat1GIF)
	})
	http.HandleFunc("/cat/2.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		w.Header().Set("Content-Length", strconv.Itoa(len(cat2GIF)))
		w.Write(cat2GIF)
	})

	fmt.Println("Starting web server at http://0.0.0.0:8081/")
	http.ListenAndServe(":8081", nil)

}
