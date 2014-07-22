package agent

import (
	"net/http"
)

func run() {

	http.HandleFunc("/dongle", dongleHandler)
	http.HandleFunc("/dongles", donglesHandler)
	http.HandleFunc("/iface", ifaceHandler)
	http.HandleFunc("/routes", routesHandler)
	http.Handle("/", http.StripPrefix("/", AssetsServer{}))
	http.ListenAndServe(":8080", nil)

}
