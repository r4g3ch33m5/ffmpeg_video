package main

import (
	"fmt"
	"os"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "ffmpeg_tool",
		Commands: []*cli.Command{
			cmd.CreateCommand,
			cmd.FfmpegCommand,
			cmd.YoutubeCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
