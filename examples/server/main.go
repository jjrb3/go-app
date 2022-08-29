package main

import "net/http"

func main() {
	http.HandleFunc("/", home)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
