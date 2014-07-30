package multiconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// MultiConnector - Configures the dongles
type MultiConnector struct {
	configuredDevices map[string]*DongleOwnerInfo
	ownerInfo         map[string]DongleOwnerInfo
	tables            map[string]string
}

// DongleOwnerInfo - stored info about a dongle
type DongleOwnerInfo struct {
	Imei               string
	Name               string
	DongleIPAddress    string
	InterfaceIPAddress string
	Network            string
}

// NewMultiConnector - Get a new MultiConnector object
func NewMultiConnector() MultiConnector {
	//cd := make(map[string]bool)
	// ifaceName: owner
	cd := make(map[string]*DongleOwnerInfo)
	// IMEI: owner
	oi := make(map[string]DongleOwnerInfo)
	tables := make(map[string]string)

	// TODO save this in a JSON or text file or something
	oi["866948014271756"] = DongleOwnerInfo{
		"866948014271756",
		"Shorty #1",
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.0/24",
	}

	oi["866948014610847"] = DongleOwnerInfo{
		"866948014610847",
		"Shorty #2",
		"192.168.2.1",
		"192.168.2.2",
		"192.168.2.0/24",
	}

	oi["866948014175684"] = DongleOwnerInfo{
		"866948014175684",
		"Jeff #1",
		"192.168.3.1",
		"192.168.3.2",
		"192.168.3.0/24",
	}

	tables["usb0"] = "uzb0"
	tables["usb1"] = "uzb1"
	tables["usb2"] = "uzb2"

	m := MultiConnector{cd, oi, tables}
	return m
}

// PossibleDongles - What network interfaces would have the dongles?
func (mc MultiConnector) PossibleDongles() []string {
	return []string{"usb0", "usb1", "usb2"}
}

// Exists - Does this network interface exist?
func (mc MultiConnector) Exists(ifaceName string) bool {
	iface, _ := net.InterfaceByName(ifaceName)
	return iface != nil
}

// IsAlreadyConfigured - Have we already done this one?
func (mc MultiConnector) IsAlreadyConfigured(ifaceName string) bool {
	// the "zero" value for a bool is false
	return mc.configuredDevices[ifaceName] != nil
}

// ImeiRequest - What is returned by the IMEI request to the dongle
type ImeiRequest struct {
	Imei string `json:"imei"`
}

// Run - Runs the multi connector
func (mc MultiConnector) Run(args []string) int {
	return run()
}

// Synopsis - Used for command line stuff
func (mc MultiConnector) Synopsis() string {
	return "Connects all the dongles"
}

// Help - Help text
func (mc MultiConnector) Help() string {
	helpText := `
Usage: command-and-control multiconnector

	Connects all the dongles.
`
	return helpText
}

// GetDeviceInfo - For a known dongle, get the owner info
func (mc MultiConnector) GetDeviceInfo() (*DongleOwnerInfo, error) {

	resp, err := http.Get("http://192.168.0.1/goform/goform_get_cmd_process?isTest=false&multi_data=1&cmd=imei")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var imeiRequest ImeiRequest

	err = json.Unmarshal(body, &imeiRequest)
	if err != nil {
		return nil, err
	}

	value, ok := mc.ownerInfo[imeiRequest.Imei]
	if ok {
		return &value, nil
	}
	return nil, errors.New("unknown dongle found - IMEI not recognised")
}

// ChangeIPAddress - Change the IP address of an network interface
func (mc MultiConnector) ChangeIPAddress(ifaceName string, address string) {
	cmd := exec.Command("sudo", "ifconfig", ifaceName, address)
	b, err := cmd.CombinedOutput()

	if err != nil {
		switch v := err.(type) {
		case *exec.ExitError:

			// lol golang
			exitCode := v.Sys().(syscall.WaitStatus)
			exitCode = (exitCode & 0xff00) >> 8

			fmt.Printf("Command exited with exit code %d\n", exitCode)

		default:
			//fmt.Printf("unexpected type %T\n", v)
			fmt.Println("Failed to call ifconfig: " + err.Error())
		}
	}

	fmt.Println(string(b))
}

// TakeDown - Take the interface down!
func (mc MultiConnector) TakeDown(ifaceName string) {
	cmd := exec.Command("sudo", "ifconfig", ifaceName, "down")
	b, err := cmd.CombinedOutput()

	if err != nil {
		switch v := err.(type) {
		case *exec.ExitError:

			// lol golang
			exitCode := v.Sys().(syscall.WaitStatus)
			exitCode = (exitCode & 0xff00) >> 8

			fmt.Printf("Command exited with exit code %d\n", exitCode)

		default:
			//fmt.Printf("unexpected type %T\n", v)
			fmt.Println("Failed to call ifconfig: " + err.Error())
		}
	}

	fmt.Println(string(b))
}

// AddGateway - Adds the specified IP address as a default route
func (mc MultiConnector) AddGateway(ifaceName string, d DongleOwnerInfo) {
	var cmd string
	table := mc.tables[ifaceName]
	// http://www.thomas-krenn.com/en/wiki/Two_Default_Gateways_on_One_System
	// ip route add 10.10.0.0/24 dev eth1 src 10.10.0.10 table rt2
	cmd = fmt.Sprintf("ip route add %s dev %s src %s table %s", d.Network, ifaceName, d.InterfaceIPAddress, table)
	runCommand("sudo", strings.Split(cmd, " ")...)
	// ip route add default via 10.10.0.1 dev eth1 table rt2
	cmd = fmt.Sprintf("ip route add default via %s dev %s table %s", d.DongleIPAddress, ifaceName, table)
	runCommand("sudo", strings.Split(cmd, " ")...)
	// ip rule add from 10.10.0.10/32 table rt2
	cmd = fmt.Sprintf("ip rule add from %s/32 table %s", d.InterfaceIPAddress, table)
	runCommand("sudo", strings.Split(cmd, " ")...)
	// ip rule add to 10.10.0.10/32 table rt2
	cmd = fmt.Sprintf("ip rule add to %s/32 table %s", d.InterfaceIPAddress, table)
	runCommand("sudo", strings.Split(cmd, " ")...)
}

func runCommand(command string, args ...string) (ret bool) {
	fmt.Printf("RUN %s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()
	ret = true

	if err != nil {
		ret = false
		switch v := err.(type) {
		case *exec.ExitError:

			// lol golang
			exitCode := v.Sys().(syscall.WaitStatus)
			exitCode = (exitCode & 0xff00) >> 8

			fmt.Printf("Command exited with exit code %d\n", exitCode)

		default:
			//fmt.Printf("unexpected type %T\n", v)
			fmt.Println("Failed to call route: " + err.Error())
		}
	}

	fmt.Println(string(b))
	return
}

// SetConfigured - Set that we have configured this one
func (mc *MultiConnector) SetConfigured(ifaceName string, owner *DongleOwnerInfo) {
	mc.configuredDevices[ifaceName] = owner
}

// ReplaceGateways - Blow away the gateways!
func (mc *MultiConnector) ReplaceGateways() {
	// sudo ip route replace default scope global nexthop via 192.168.1.1 dev usb0 weight 1 nexthop via 192.168.2.1 dev usb1 weight 1
	var hops string
	for iface, hop := range mc.configuredDevices {
		hops = hops + " nexthop via " + hop.DongleIPAddress + " dev " + iface + " weight 1"
	}
	runCommand("sudo", strings.Split("ip route replace default scope global"+hops, " ")...)
}

func run() int {

	m := NewMultiConnector()

	for {

		for _, ifaceName := range m.PossibleDongles() {

			// does usb0 exist?
			if !m.Exists(ifaceName) {
				fmt.Printf("Can't find %s, giving up\n", ifaceName)
				break
			}

			// have we configured that one already?
			if m.IsAlreadyConfigured(ifaceName) {
				fmt.Printf("Interface %s is already configured, trying the next one\n", ifaceName)
				continue
			}

			m.ChangeIPAddress(ifaceName, "192.168.0.182")
			fmt.Printf("Changed the IP address of %s to 192.168.0.182\n", ifaceName)

			// which dongle is it (IMEI -> person) ?
			ownerInfo, err := m.GetDeviceInfo()
			// if ownerInfo == nil {
			// 	continue
			// }
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("The dongle at %s is identified as %s\n", ifaceName, ownerInfo.Name)

			m.TakeDown(ifaceName)
			fmt.Printf("Dongle %s is now down\n", ifaceName)

			// ip route add 10.10.0.0/24 dev eth1 src 10.10.0.10 table rt2
			// ip route add default via 10.10.0.1 dev eth1 table rt2
			// ip rule add from 10.10.0.10/32 table rt2
			// ip rule add to 10.10.0.10/32 table rt2

			// give that interface that particular IP address/range
			m.ChangeIPAddress(ifaceName, ownerInfo.InterfaceIPAddress)
			fmt.Printf("Changed the IP address of %s to %s\n", ifaceName, ownerInfo.InterfaceIPAddress)

			// add that IP as a Internet gateway
			m.AddGateway(ifaceName, *ownerInfo)
			fmt.Printf("Added route for %s to %s\n", ifaceName, ownerInfo.DongleIPAddress)

			// it is now configured
			m.SetConfigured(ifaceName, ownerInfo)

			// make shit work
			m.ReplaceGateways()

			// next
		}
		fmt.Println("Sleeping...")
		time.Sleep(10 * time.Second)
	}
	return 0
}
