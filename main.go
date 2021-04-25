package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linde12/gowol"
)

func sendWakeOnLan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mac := params["mac"]

	if packet, err := gowol.NewMagicPacket(mac); err == nil {
		packet.Send("255.255.255.255")
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/wake/{mac}", sendWakeOnLan).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}
