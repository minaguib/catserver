package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

//go:embed cat1.gif
var cat1GIF []byte

//go:embed cat2.gif
var cat2GIF []byte

//go:embed cat.html
var catHTML []byte

func main() {

	var catN uint32

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(catHTML)))
		w.Write(catHTML)
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
	http.HandleFunc("/cat.gif", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddUint32(&catN, 1)
		path := fmt.Sprintf("/cat/%d.gif", (n%2)+1)
		http.Redirect(w, r, path, http.StatusFound)
	})

	fmt.Println("Starting web server at http://0.0.0.0:8081/")
	http.ListenAndServe(":8081", nil)

}
