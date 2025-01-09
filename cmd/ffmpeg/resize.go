package ffmpeg

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ResizeCommand defines the subcommand for resizing a video
var ResizeCommand = &cli.Command{
	Name:    "resize",
	Aliases: []string{"r"},
	Usage:   "Resize a video to specific dimensions",
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
			Usage:    "Path to the output video file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "resolution",
			Aliases:  []string{"res"},
			Usage:    "Target resolution (e.g., 1280x720)",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		input := c.String("input")
		output := c.String("output")
		resolution := c.String("resolution")
		fmt.Printf("Resizing '%s' to resolution '%s' and saving as '%s'...\n", input, resolution, output)
		// Add your FFmpeg logic here
		return nil
	},
}
