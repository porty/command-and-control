package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/porty/command-and-control/dongler"
	"math/rand"
	"net/http"
	"strconv"
)

type DongleInfo struct {
	Interface string
	PingMS    uint32
	Raw       dongler.RawInfo
}

func getDongles() *[]DongleInfo {

	dongles := make([]DongleInfo, 2, 2)

	var raw dongler.RawInfo

	raw = dongler.RawInfo{"866948014271756", "505013457518600", "10.1.2.3", "6969::6969", "ppp_connected", "LTE", strconv.Itoa(rand.Intn(6)), "Only_LTE"}
	dongles[0] = DongleInfo{"usb0", uint32(rand.Intn(2000)), raw}

	raw = dongler.RawInfo{"866948014610847", "505013457711705", "10.2.2.3", "7070::7070", "ppp_disconnected", "UMTS", strconv.Itoa(rand.Intn(6)), "NETWORK_auto"}
	dongles[1] = DongleInfo{"usb1", uint32(rand.Intn(2000)), raw}

	return &dongles
}

func donglesHandler(w http.ResponseWriter, r *http.Request) {

	dongles := getDongles()

	b, err := json.Marshal(dongles)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(w)
	fmt.Fprintln(w)
}
