package cmd

import (
	"context"
	"fmt"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd/ffmpeg"
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/youtube"
	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

var FullFlowCommand = &cli.Command{
	Name: "full",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "source",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		},
		&cli.StringFlag{
			Name: "input",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		source := c.String("source")
		output := c.String("output")
		if output == "" {
			output = util.GetTodayFolder()
		}
		err := youtube.DownloadVideo(source, output)
		if err != nil {
			return err
		}
		util.ApplyAllFileInDir(output, func(path string) error {
			c.Set("input", path)
			// c.Set("output", filepath.Join("watermarked", output+".mp4"))
			err := ffmpeg.AddWatermarkCommand.Action(ctx, c)
			if err != nil {
				fmt.Println(err)
			}
			err = c.Set("input", output)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = ffmpeg.SplitByChunksCommand.Action(ctx, c)
			if err != nil {
				return err
			}
			return err
		})

		return nil
	},
}
