package cmd

import (
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/ffmpeg"

	"github.com/urfave/cli/v2"
)

// FfmpegCommand defines the main ffmpeg command group
var FfmpegCommand = &cli.Command{
	Name:  "ffmpeg",
	Usage: "Perform various FFmpeg operations",
	Subcommands: []*cli.Command{
		ffmpeg.ConvertCommand,
		ffmpeg.ResizeCommand,
		ffmpeg.ExtractCommand,
	},
}
