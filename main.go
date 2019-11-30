package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

var revision = "HEAD"

func main() {
	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		log.Fatal("error", err.Error())
		os.Exit(exitCode)
	}
}

func newApp() *cli.App {
	app := &cli.App{}
	app.Name = "github-toy"
	app.Usage = "Toy to operate github"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, revision)
	app.Commands = commands
	return app
}
