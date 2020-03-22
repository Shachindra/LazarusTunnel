package apiutil

import (
	"io/ioutil"
	"log"
	"net/http"
)

// TunnelConfigHandler Handle Tunnel related requests for creation, statistics and deletion
func TunnelConfigHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		jsonMsg, err := ioutil.ReadAll(r.Body)
		log.Printf("Got a GET Request with request: %v" + string(jsonMsg))
		defer r.Body.Close()
		if err != nil {
			log.Printf("Error Reading Message Body: %v", err)
			http.Error(w, "Can't Read Message Body. "+err.Error(), http.StatusBadRequest) //HTTP 400
			return
		}
	}
	if r.Method == http.MethodPost {
		jsonMsg, err := ioutil.ReadAll(r.Body)
		log.Printf("Got a GET Request with request: %v" + string(jsonMsg))
		defer r.Body.Close()
		if err != nil {
			log.Printf("Error Reading Message Body: %v", err)
			http.Error(w, "Can't Read Message Body. "+err.Error(), http.StatusBadRequest) //HTTP 400
			return
		}
	}
	return
}
