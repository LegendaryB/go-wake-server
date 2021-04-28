package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/linde12/gowol"
)

const ConfigurationFileName = "conf.json"

type Broadcast struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Configuration struct {
	Port      string    `json:"port"`
	Broadcast Broadcast `json:"broadcast"`
}

func sendWakeOnLAN(mac string, broadcast Broadcast) error {
	packet, err := gowol.NewMagicPacket(mac)

	if err != nil {
		return err
	}

	return packet.SendPort(broadcast.Address, broadcast.Port)
}

func WakeOnLANHandler(conf *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mac := params["mac"]

		fmt.Printf("Http endpoint called with parameter: %s.\n", mac)

		match, _ := regexp.MatchString("^([0-9A-F]{2}[:-]){5}([0-9A-F]{2})$", mac)

		if !match {
			http.Error(w, fmt.Sprintf("'%s' is not a valid MAC address.", mac), http.StatusBadRequest)
			return
		}

		err := sendWakeOnLAN(mac, conf.Broadcast)

		if err != nil {
			fmt.Printf("Failed to send magic packet to '%s'.\n", mac)
			http.Error(w, "", http.StatusBadRequest)
		}
	}
}

func parseConfigurationFile() (*Configuration, error) {
	buffer, err := ioutil.ReadFile(ConfigurationFileName)

	if err != nil {
		return nil, err
	}

	var conf Configuration

	err = json.Unmarshal(buffer, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func listenAndServe(port string, router *mux.Router) error {
	addr := ":" + port

	fmt.Printf("Starting server on port %s...\n", port)
	return http.ListenAndServe(addr, router)
}

func main() {
	conf, err := parseConfigurationFile()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/wake/{mac}", WakeOnLANHandler(conf))
	fmt.Println("Added handler for http endpoint '/wake/{mac}'.")

	log.Fatal(listenAndServe(conf.Port, router))
}
