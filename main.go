package main

import (
	"context"
	"fmt"
	"os"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name: "ffmpeg_tool",
		Commands: []*cli.Command{
			cmd.CreateFolderCommand,
			cmd.FfmpegCommand,
			cmd.YoutubeCommand,
			cmd.GptCommand,
		},
	}
	ctx := context.TODO()
	if err := app.Run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
