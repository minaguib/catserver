package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed cat.gif
var cat_gif []byte

//go:embed cat.html
var cat_html []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(cat_html)
	})
	http.HandleFunc("/cat.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		w.Write(cat_gif)
	})
	fmt.Println("Starting web server at http://0.0.0.0:8081/")
	http.ListenAndServe(":8081", nil)
}
