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
	http.HandleFunc("/teapot", teapot)
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
	if val, ok := record.Country.Names["en"]; ok {
		fmt.Fprint(w, val)
	}

	if len(record.Subdivisions) > 1 {
		if val, ok := record.Subdivisions[0].Names["en"]; ok {
			fmt.Fprintln(w, " ", val)
		}
	}

	//fmt.Fprint(w, record.Country.Names["zh-CN"])
	//fmt.Fprintln(w, " ",record.City.Names["zh-CN"])
	//fmt.Fprintln(w, " ", record.Location.Latitude, record.Location.Longitude)
}

func teapot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("<pre>"))
	w.Write([]byte("                                              /\n"))
	w.Write([]byte("                                             /\n"))
	w.Write([]byte("                             xxX###xx       /\n"))
	w.Write([]byte("                              ::XXX        /\n"))
	w.Write([]byte("                       xxXX::::::###XXXXXx/#####\n"))
	w.Write([]byte("                  :::XXXXX::::::XXXXXXXXX/    ####\n"))
	w.Write([]byte("       xXXX//::::::://///////:::::::::::/#####    #         ##########\n"))
	w.Write([]byte("    XXXXXX//:::::://///xXXXXXXXXXXXXXXX/#    #######      ###   ###\n"))
	w.Write([]byte("   XXXX        :://///XXXXXXXXX######X/#######      #   ###    #\n"))
	w.Write([]byte("   XXXX        ::////XXXXXXXXX#######/ #     #      ####   #  #\n"))
	w.Write([]byte("    XXXX/:     ::////XXXXXXXXXX#####/  #     #########      ##\n"))
	w.Write([]byte("     \"\"XX//::::::////XXXXXXXXXXXXXX/###########     #       #\n"))
	w.Write([]byte("         \"::::::::////XXXXXXXXXXXX/    #     #     #      ##\n"))
	w.Write([]byte("               ::::////XXXXXXXXXX/##################   ###\n"))
	w.Write([]byte("                   ::::://XXXXXX/#    #     #   #######\n"))
	w.Write([]byte("                       ::::::::/################\n"))
	w.Write([]byte("                              /\n"))
	w.Write([]byte("                             /\n"))
	w.Write([]byte("                            /\n"))
	w.Write([]byte("            <font size=\"7\">We're a Teapot!!</font>  "))
	w.Write([]byte("</pre>"))
}
