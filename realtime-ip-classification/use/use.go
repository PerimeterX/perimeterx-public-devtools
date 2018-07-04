package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	maxminddb "github.com/oschwald/maxminddb-golang"
)

// ClassificationRecord IP classification record
type ClassificationRecord struct {
	Service string `maxminddb:"service"`
	Region  string `maxminddb:"region"`
}

func getRealIP(r *http.Request) string {
	rip := r.Header.Get("X-REAL-IP")
	if rip == "" {
		rip = r.RemoteAddr
	}
	return rip
}

// enrichClassification enrich request with classification information on headers
func enrichClassification(reader *maxminddb.Reader, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log request real ip
		rip := getRealIP(r)
		log.Printf("Request from %s", rip)
		// parse ip and enrich classification information
		addr := net.ParseIP(rip)
		if addr != nil {
			var record ClassificationRecord
			_ = reader.Lookup(addr, &record)
			if record.Service != "" {
				r.Header.Set("X-IP-SERVICE", record.Service)
			}
			if record.Region != "" {
				r.Header.Set("X-IP-REGION", record.Region)
			}
		}
		next.ServeHTTP(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "What would life be if we had no courage to attempt anything?\n")
	// use enriched information if found on request
	if service := r.Header.Get("X-IP-SERVICE"); service != "" {
		fmt.Fprintf(w, "Known service %s\n", service)
	}
	if region := r.Header.Get("X-IP-REGION"); region != "" {
		fmt.Fprintf(w, "Known region %s\n", region)
	}
}

func main() {
	// load mmdb with classifications
	mmdb, err := maxminddb.Open("feed.mmdb")
	if err != nil {
		log.Fatalln(err)
	}
	defer mmdb.Close()
	// register routes and serve requests
	http.Handle("/", enrichClassification(mmdb, index))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
