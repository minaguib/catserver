package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
)

//go:embed cat.gif
var catGIF []byte

//go:embed cat.html
var catHTML []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(catHTML)))
		w.Write(catHTML)
	})
	http.HandleFunc("/cat.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		w.Header().Set("Content-Length", strconv.Itoa(len(catGIF)))
		w.Write(catGIF)
	})
	fmt.Println("Starting web server at http://0.0.0.0:8081/")
	http.ListenAndServe(":8081", nil)
}
