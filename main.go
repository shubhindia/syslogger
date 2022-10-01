package main

import (
	"os"

	"github.com/shubhindia/syslogger/common"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "syslogger"
	app.Usage = "syslogger is a command line tool/daemon for collecting process info"
	app.Version = common.Version()
	app.Authors = []*cli.Author{
		{
			Name:  "Shubham Gopale",
			Email: "shubhindia123@gmail.com",
		},
	}
	app.Commands = common.GetCommands()
	app.CommandNotFound = func(context *cli.Context, command string) {
		logrus.Fatalln("Command", command, "not found.")
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
