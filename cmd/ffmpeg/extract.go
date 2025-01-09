package ffmpeg

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ExtractCommand defines the subcommand for extracting audio from a video
var ExtractCommand = &cli.Command{
	Name:    "extract",
	Aliases: []string{"e"},
	Usage:   "Extract audio from a video file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Aliases:  []string{"i"},
			Usage:    "Path to the input video file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Path to the output audio file",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		input := c.String("input")
		output := c.String("output")
		fmt.Printf("Extracting audio from '%s' and saving as '%s'...\n", input, output)
		// Add your FFmpeg logic here
		return nil
	},
}
