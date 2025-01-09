package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
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
	Action: func(ctx context.Context, c *cli.Command) error {
		name := c.String("name")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	},
}
