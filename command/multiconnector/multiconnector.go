package multiconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"syscall"
	"time"
)

// MultiConnector - Configures the dongles
type MultiConnector struct {
	configuredDevices map[string]bool
	ownerInfo         map[string]DongleOwnerInfo
}

// DongleOwnerInfo - stored info about a dongle
type DongleOwnerInfo struct {
	Imei               string
	Name               string
	DongleIPAddress    string
	InterfaceIPAddress string
}

// NewMultiConnector - Get a new MultiConnector object
func NewMultiConnector() MultiConnector {
	cd := make(map[string]bool)
	oi := make(map[string]DongleOwnerInfo)

	// TODO save this in a JSON or text file or something
	oi["866948014271756"] = DongleOwnerInfo{
		"866948014271756",
		"Shorty #1",
		"192.168.1.1",
		"192.168.1.2",
	}

	oi["866948014610847"] = DongleOwnerInfo{
		"866948014610847",
		"Shorty #2",
		"192.168.2.1",
		"192.168.2.2",
	}

	oi["866948014175684"] = DongleOwnerInfo{
		"866948014175684",
		"Jeff #1",
		"192.168.3.1",
		"192.168.3.2",
	}

	m := MultiConnector{cd, oi}
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
	return mc.configuredDevices[ifaceName]
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

// AddGateway - Adds the specified IP address as a default route
func (mc MultiConnector) AddGateway(gateway string) {
	// sudo route add default gw 192.168.3.1 metric 10
	cmd := exec.Command("sudo", "route", "add", "default", "gw", gateway, "metric", "10")
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
			fmt.Println("Failed to call route: " + err.Error())
		}
	}

	fmt.Println(string(b))
}

// SetConfigured - Set that we have configured this one
func (mc *MultiConnector) SetConfigured(ifaceName string) {
	mc.configuredDevices[ifaceName] = true
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

			// give that interface that particular IP address/range
			m.ChangeIPAddress(ifaceName, ownerInfo.InterfaceIPAddress)
			fmt.Printf("Changed the IP address of %s to %s\n", ifaceName, ownerInfo.InterfaceIPAddress)

			// add that IP as a Internet gateway
			m.AddGateway(ownerInfo.DongleIPAddress)
			fmt.Printf("Added route for %s to %s\n", ifaceName, ownerInfo.DongleIPAddress)

			// it is now configured
			m.SetConfigured(ifaceName)

			// next
		}
		fmt.Println("Sleeping...")
		time.Sleep(10 * time.Second)
	}
	return 0
}
