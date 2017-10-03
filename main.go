package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var addr = flag.String("addr", ":8080", "HTTP listening address")
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, http.HandlerFunc(h)))
}

func h(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(string(b))
}
