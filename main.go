package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/linde12/gowol"
)

type Broadcast struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Configuration struct {
	Port            string    `json:"port"`
	MACRegexPattern string    `json:"mac_regex_pattern"`
	Broadcast       Broadcast `json:"broadcast"`
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

		match, _ := regexp.MatchString(conf.MACRegexPattern, mac)

		if !match {
			http.Error(w, mac+" is not a valid MAC address!", http.StatusBadRequest)
			return
		}

		err := sendWakeOnLAN(mac, conf.Broadcast)

		if err != nil {
			http.Error(w, "Failed to sent magic packet to "+mac, http.StatusBadRequest)
		}

		w.Write([]byte(http.StatusText(200)))
	}
}

func parseConfigurationFile(filename string) (*Configuration, error) {
	buffer, err := ioutil.ReadFile(filename)

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
	var configurationFileName string

	flag.StringVar(&configurationFileName, "c", "conf.json", "Use to specify a custom configuration file.")
	flag.Parse()

	conf, err := parseConfigurationFile(configurationFileName)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/wake/{mac}", WakeOnLANHandler(conf))
	fmt.Println("Added handler for http endpoint '/wake/{mac}'.")

	log.Fatal(listenAndServe(conf.Port, router))
}
