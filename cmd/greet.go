package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// GreetCommand defines the greet command
var GreetCommand = &cli.Command{
	Name:    "greet",
	Aliases: []string{"g"},
	Usage:   "Greet someone",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name of the person to greet",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		name := c.String("name")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	},
}
