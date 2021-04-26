package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/linde12/gowol"
)

type Configuration struct {
	Port             int
	UseStaticMac     bool
	MACAddress       string
	BroadcastAddress string
}

func WakeOnLANHandler(conf Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mac := conf.MACAddress

		if !conf.UseStaticMac {
			params := mux.Vars(r)
			mac = params["mac"]
		}

		if packet, err := gowol.NewMagicPacket(mac); err == nil {
			packet.Send(conf.BroadcastAddress)
		} else {
			http.Error(w, "Not a valid MAC address.", http.StatusBadRequest)
		}
	}
}

func printConfiguration() {
	fmt.Println("Using configuration:")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name + "\t" + f.Value.String())
	})
}

func main() {
	conf := Configuration{}

	flag.IntVar(&conf.Port, "port", 81, "To specify the http listener port.")
	flag.BoolVar(&conf.UseStaticMac, "use-static-mac", false, "Flag to indicate if a given mac should be used when the http resource is called.")
	flag.StringVar(&conf.MACAddress, "mac-addr", "", "Static MAC address which is used when the UseStaticMac flag is set to true.")
	flag.StringVar(&conf.BroadcastAddress, "broadcast-addr", "255.255.255.255", "Address to which the generated MagicPacket will be send.")

	flag.Parse()

	printConfiguration()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/wake/{mac}", WakeOnLANHandler(conf)).Methods("GET")

	addr := ":" + strconv.Itoa(conf.Port)

	log.Fatal(http.ListenAndServe(addr, router))
}
