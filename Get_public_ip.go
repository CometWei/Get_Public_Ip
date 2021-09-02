package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomasen/realip"
)

func main() {
	http.HandleFunc("/", getip)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	clientIP := realip.FromRequest(r)
	log.Println("GET from", clientIP)
	fmt.Fprintln(w, clientIP)
}
