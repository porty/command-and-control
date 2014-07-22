package multiconnector

import (
	"fmt"
	"github.com/porty/command-and-control/config"
	"net/http"
)

func donglesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, "{}")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, "This is for the dongle multi-connector status information")
	fmt.Fprint(w, "Try looking at /dongles.json")
	fmt.Fprint(w, "Either that, or try a different port")
}

func getListenString() string {
	return fmt.Sprintf(":%d", config.DongleMultiConnectorPort)
}

func listenForHTTPRequests() {
	http.HandleFunc("/dongles.json", donglesHandler)
	http.HandleFunc("/", infoHandler)
	http.ListenAndServe(getListenString(), nil)
}
