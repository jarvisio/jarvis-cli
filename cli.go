package main

import (
	"github.com/codegangsta/cli"
	"github.com/jarvisio/commands"
	"os"
	"os/exec"
	"runtime"
)

// Main

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "jarvis"
	app.Usage = "I am a nice butler, who can speak and listen"
	app.Commands = []cli.Command{timeCommand(), dateCommand()}
	app.Run(os.Args)
}

// Commands

func timeCommand() cli.Command {
	command := cli.Command{
		Name:      "clock",
		ShortName: "c",
		Usage:     "Tells you the current time",
		Action: func(c *cli.Context) {
			say("The time is" + commands.Clock(false))
		},
	}
	return command
}

func dateCommand() cli.Command {
	command := cli.Command{
		Name:      "date",
		ShortName: "d",
		Usage:     "Tells you today's date",
		Action: func(c *cli.Context) {
			say("Today is" + commands.Date())
		},
	}
	return command
}

// Helper Methods

func say(text string) {
	os := runtime.GOOS
	switch os {
	case "linux":
		exec.Command("tts", text).Run()
	case "darwin":
		exec.Command("say", text).Run()
	default:
		println(text)
	}
}
