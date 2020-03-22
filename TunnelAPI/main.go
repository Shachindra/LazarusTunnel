package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./apiutil"
)

func tunnelConfigHandler(w http.ResponseWriter, r *http.Request) {
	apiutil.TunnelConfigHandler(w, r)
}

func setupRoutes() {
	// Handle VPN Related Requests
	http.HandleFunc("/tunnel/api/v1/config", tunnelConfigHandler)

	// Handle Any Request
	port := flag.String("p", "80", "Port to Serve UI on")
	directory := flag.String("d", "./httpdocs", "Directory of Static HTML files to host")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*directory))
	http.Handle("/", fileServer)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}
	log.Fatal(err)
}

func main() {
	fmt.Println("Starting TunnelAPI Server")
}
