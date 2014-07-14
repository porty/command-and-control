package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Route struct {
	Destination string
	Gateway     string
	Genmask     string
	Flags       string
	Metric      uint32
	Ref         uint32
	Use         uint32
	Iface       string
}

func getRoutes() []Route {
	routes := make([]Route, 3, 3)
	routes[0] = Route{"0.0.0.0", "192.168.0.1", "0.0.0.0", "UG", 0, 0, 0, "usb0"}
	routes[1] = Route{"192.168.0.0", "0.0.0.0", "255.255.255.0", "U", 0, 0, 0, "usb0"}
	routes[2] = Route{"192.168.1.0", "0.0.0.0", "255.255.255.0", "U", 0, 0, 0, "eth0"}
	return routes
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	routes := getRoutes()

	b, err := json.Marshal(routes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(w)
	fmt.Fprintln(w)
}
