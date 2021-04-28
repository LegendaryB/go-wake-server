package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linde12/gowol"
)

type BroadcastConfiguration struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Configuration struct {
	Port        string `json:"port"`
	AllowAnyMAC bool   `json:"allow_any_mac"`
	MACAddress  string `json:"mac_address"`

	Broadcast BroadcastConfiguration `json:"broadcast"`
}

func sendWakeOnLAN(mac string, broadcast_addr string, broadcast_port string) error {
	packet, err := gowol.NewMagicPacket(mac)

	if err != nil {
		return err
	}

	err = packet.SendPort(broadcast_addr, broadcast_port)
	return err
}

func StaticWakeOnLANHandler(conf Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := sendWakeOnLAN(conf.MACAddress, conf.Broadcast.Address, conf.Broadcast.Port)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func NotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Disabled by configuration.", http.StatusMethodNotAllowed)
}

func WakeOnLANHandler(broadcast_addr string, broadcast_port string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mac := params["mac"]

		err := sendWakeOnLAN(mac, broadcast_addr, broadcast_port)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func main() {
	conf := getConfiguration("conf.json")

	fmt.Println("Listening on port " + conf.Port)
	fmt.Printf("")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/wake/", StaticWakeOnLANHandler(conf))

	if conf.AllowAnyMAC {
		router.HandleFunc("/wake/{mac}", WakeOnLANHandler(conf.Broadcast.Address, conf.Broadcast.Port))
	} else {
		router.HandleFunc("/wake/{mac}", NotAllowedHandler)
	}

	listenaddr := ":" + conf.Port

	log.Fatal(http.ListenAndServe(listenaddr, router))
}

func getConfiguration(file string) Configuration {
	buffer, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	var conf Configuration

	err = json.Unmarshal(buffer, &conf)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}
