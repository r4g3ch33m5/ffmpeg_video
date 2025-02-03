package cmd

import (
	"context"

	cronjob "github.com/r4g3ch33m5/ffmpeg_video/cmd/cron_job"
	"github.com/urfave/cli/v3"
)

var tasks = []cronjob.Task{}

var StartCronCommand = &cli.Command{
	Name: "cron",
	Action: func(ctx context.Context, c *cli.Command) error {
		// cm := cronjob.NewCronManager()

		return nil
	},
}
