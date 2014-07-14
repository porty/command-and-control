package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Dongle struct {
	Interface       string
	IMEI            string
	LocalIp         string
	Connection      string
	SessionTransfer uint64
	PingMS          uint32
}

func getDongle(iface string) *Dongle {
	if iface == "usb0" {
		return &Dongle{"usb0", "123123123", "192.168.1.1", "LTE", 4040400, 250}
	}
	return nil
}

func dongleHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	iface := r.Form.Get("iface")
	if iface == "" {
		http.Error(w, "Dongle not specified", 404)
		return
	}
	dongle := getDongle(iface)
	if dongle == nil {
		http.Error(w, "Dongle not found", 404)
		return
	}

	b, err := json.Marshal(dongle)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(w)
	fmt.Fprintln(w)
}
