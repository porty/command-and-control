package main

import (
	"os"
	"os/signal"

	"github.com/juju/loggo"

	"github.com/mitchellh/cli"
	"github.com/porty/command-and-control/command"
	"github.com/porty/command-and-control/command/agent"
	"github.com/porty/command-and-control/command/multiconnector"
	"github.com/porty/command-and-control/command/uploader"
)

// Commands is the mapping of all the available Serf commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	log := loggo.GetLogger("myapp")

	Commands = map[string]cli.CommandFactory{

		"agent": func() (cli.Command, error) {
			return &agent.Command{
				Ui:         ui,
				ShutdownCh: make(chan struct{}),
				Log:        log,
			}, nil
		},

		"multiconnector": func() (cli.Command, error) {
			return multiconnector.NewMultiConnector(), nil
		},

		"dongler": func() (cli.Command, error) {
			return &command.DonglerCommand{
				Ui: ui,
			}, nil
		},

		"upload": func() (cli.Command, error) {
			return &uploader.UploadCommand{}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
