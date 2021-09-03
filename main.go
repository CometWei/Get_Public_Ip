package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
	"github.com/tomasen/realip"
)

func main() {
	port := (":" + os.Getenv("PORT"))
	http.HandleFunc("/", getip)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	clientIP := realip.FromRequest(r)
	log.Println("GET from", clientIP)
	fmt.Fprintln(w, clientIP)

	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(clientIP)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, " ", record.Country.Names["en"])
	fmt.Fprintln(w, " ", record.City.Names["zh-TW"])
	fmt.Fprintln(w, " ", record.Subdivisions[0].Names["en"])
	//	fmt.Fprintln(w, " ", record.Location.Latitude, record.Location.Longitude)
}
