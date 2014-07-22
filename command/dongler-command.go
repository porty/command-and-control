package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/porty/command-and-control/dongler"
	"os"
)

// DonglerCommand is a Command implementation
type DonglerCommand struct {
	Ui   cli.Ui
	host string
}

func (c *DonglerCommand) Help() string {
	return "It gets the status yay"
}

func (c *DonglerCommand) Run(args []string) int {

	if len(args) == 1 {
		c.host = args[0]
	}

	if c.host == "" {
		fmt.Println("No host specified")
		return 1
	}

	info, err := dongler.GetRawInfo(c.host)

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	fmt.Println("===================")
	fmt.Println("Raw Info")
	fmt.Println("===================")

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	out.WriteTo(os.Stdout)
	fmt.Println()

	return 0
}

func (c *DonglerCommand) Synopsis() string {
	return "Prints information about a specified dongle"
}
