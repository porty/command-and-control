package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Iface struct {
	Name            string
	HardwareAddress string
	IPAddresses     []string
	Throughput      []uint64
	Total           []uint64
	Type            string
}

func getIface(iface string) *Iface {
	if iface == "usb0" {
		return &Iface{"usb0",
			"00:11:22:33:44:55",
			[]string{"192.168.1.1"},
			[]uint64{1024, 2048},
			[]uint64{102400, 204800},
			"dongle",
		}
	}
	return nil
}

func ifaceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("iface")
	if name == "" {
		http.Error(w, "Interface not specified", 404)
		return
	}
	iface := getIface(name)
	if iface == nil {
		http.Error(w, "Interface not found", 404)
		return
	}

	b, err := json.Marshal(iface)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(w)
	fmt.Fprintln(w)
}
