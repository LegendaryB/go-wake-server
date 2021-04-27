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

func StaticWakeOnLANHandler(conf Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		packet, err := gowol.NewMagicPacket(conf.MACAddress)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = packet.SendPort(conf.Broadcast.Address, conf.Broadcast.Port)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func WakeOnLANHandler(conf Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !conf.AllowAnyMAC {
			http.Error(w, "Disabled by configuration.", http.StatusMethodNotAllowed)
			return
		}

		params := mux.Vars(r)
		mac := params["mac"]

		packet, err := gowol.NewMagicPacket(mac)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = packet.SendPort(conf.Broadcast.Address, conf.Broadcast.Port)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func main() {
	conf := getConfiguration("conf.json")
	printConfiguration(conf)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/wake/", StaticWakeOnLANHandler(conf))
	router.HandleFunc("/wake/{mac}", WakeOnLANHandler(conf))

	listenaddr := ":" + conf.Port

	log.Fatal(http.ListenAndServe(listenaddr, router))
}

func printConfiguration(v interface{}) {
	json, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Using configuration:\n%s\n", string(json))
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
