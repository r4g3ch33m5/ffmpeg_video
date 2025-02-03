package cmd

import (
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/gpt"
	"github.com/urfave/cli/v3"
)

var GptCommand = &cli.Command{
	Name:  "gpt",
	Usage: "Perform various operations on YouTube videos",
	Commands: []*cli.Command{
		gpt.SummarizeCommand,
	},
}
