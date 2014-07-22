package dongler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// RawInfo - modem info
type RawInfo struct {
	// 866948014271756 - Rob #1
	// 866948014610847 - Rob #2
	Imei         string `json:"imei"`
	// 505013457518600 - Rob #1
	// 505013457711705 - Rob #2
	Imsi         string `json:"sim_imsi"`
	// The IP address (usually 10.x.x.x)
	ExternalIPv4 string `json:"wan_ipaddr"`
	// The IPv6 address (or blank if not available)
	ExternalIPv6 string `json:"ipv6_wan_ipaddr"`
	// "ppp_connecting" or "ppp_connected" or "ppp_disconnecting" or "ppp_disconnected"
	PPPStatus    string `json:"ppp_status"`
	// "LTE" (4G)
	// "UMTS" (3G (HSDPA?))
	// "EDGE" (2.5G)
	// "LIMITED_SERVICE_GSM" 2G??
	// "NO_SERVICE" (nothing)
	NetworkType  string `json:"network_type"`
	// "0" .. "5"
	SignalBar    string `json:"signalbar"`
	// "NETWORK_auto" (whatever)
	// "Only_LTE" (4G only)
	// "Only_WCDMA" (3G only)
	// "WCDMA_AND_GSM" (3G/2G only)
	// "Only_GSM" (2G only)
	NetSelection string `json:"net_select"`
}

func GetRawInfo(ip string) (*RawInfo, error) {
	resp, err := http.Get("http://" + ip + "/goform/goform_get_cmd_process?isTest=false&multi_data=1&cmd=imei%2Csim_imsi%2Cwan_ipaddr%2Cipv6_wan_ipaddr%2Cppp_status%2Cnetwork_type%2Csignalbar%2Cnet_select")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw RawInfo

	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	return &raw, err
}

func main() {
	fmt.Println(len(os.Args))
}
