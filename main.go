package main

import (
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Start the server
	http.ListenAndServe(":8000", nil)
}
