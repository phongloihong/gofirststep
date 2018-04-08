package main

import (
	"io"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", serveImage)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./"))))
	appengine.Main()
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/go-logo.png"> <h1>Hello man</h1>`)
}
