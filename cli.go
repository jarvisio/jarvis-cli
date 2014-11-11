package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "jarvis"
	app.Usage = "I am a nice butler, who can speak and listen"
	app.Run(os.Args)
}
