package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandSettings,
}

var commandSettings = &cli.Command{
	Name:  "settings",
	Usage: "Operate github display name",
	Subcommands: []*cli.Command{
		{
			Name:   "update",
			Usage:  "Update github display name",
			Action: doSettingsUpdate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
				},
			},
		},
	},
}

func doSettingsUpdate(c *cli.Context) error {
	var (
		argName = c.String("name")
	)

	if argName != "" {
		fmt.Println(argName)
	}
	return nil
}
