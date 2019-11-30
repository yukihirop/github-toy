package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandSettings,
}

var commandSettings = &cli.Command{
	Name:  "settings",
	Usage: "Operate github display name",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "login",
			Aliases: []string{"l"},
		},
	},
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
		login       = c.String("login")
		displayName = c.String("name")
	)

	if login != "" && displayName != "" {
		client := NewClient()
		userBody := &github.User{
			Login: &login,
			Name:  &displayName,
		}
		userRes, _, err := client.Users.Edit(context.Background(), userBody)
		if err == nil {
			fmt.Printf("update github display name: %s\n", *userRes.Name)
		} else {
			log.Fatal("error", err.Error())
		}
	}
	return nil
}
