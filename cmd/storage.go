package cmd

import (
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/storage"
	"github.com/urfave/cli/v3"
)

// CreateFolderCommand defines the CLI command for starting the cron scheduler
var CreateFolderCommand = &cli.Command{
	Name:  "storage",
	Usage: "Start the cron job for creating daily folders",
	Commands: []*cli.Command{
		storage.CreateLocalFolderCommand,
	},
}
