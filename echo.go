package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

func main() {
	var port *int = flag.Int("Port", 8090, "Listen on this port")
	var help *bool = flag.Bool("Help", false, "Print help")
	flag.Parse()
	if *help {
		flag.CommandLine.Usage()
		os.Exit(1)
	}
	http.HandleFunc("/", handler)
	fmt.Printf("Starting listener on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	fmt.Fprintf(w, "[Headers]\n")
	keys := []string{}
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(w, "%s=%s\n", k, r.Header[k])
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "\n\n-----\n(c) %d Philipp Ritter\n", time.Now().Year())
}
