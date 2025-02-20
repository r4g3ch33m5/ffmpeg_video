package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd/ffmpeg"
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/storage"
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/youtube"
	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

var titleReg = regexp.MustCompile(`__(.*?)__`)

var FullFlowCommand = &cli.Command{
	Name: "full",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "source",
			// Required: true,
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
		dateFolder := util.GetTodayFolder()
		err := storage.CreateLocalFolderCommand.Action(ctx, c)
		if err != nil {
			return err
		}
		sourceFolder := filepath.Join("source", dateFolder)
		sourceListFile := filepath.Join(sourceFolder, "source.txt")
		err = youtube.DownloadVideo(sourceListFile, sourceFolder, youtube.DownloadOption{IsBatchDownload: true})
		if err != nil {
			fmt.Println("err download video", err)
			exec.Command("code", sourceListFile).Run()
			return err
		}
		util.ApplyAllFileInDir(sourceFolder, func(path string) error {
			c.Set("input", path)
			if strings.HasSuffix(path, ".txt") {
				return nil
			}
			watermarkedFile := filepath.Join("watermarked", dateFolder, filepath.Base(path))
			fmt.Println(watermarkedFile)
			c.Set("output", watermarkedFile)

			err := ffmpeg.AddWatermarkCommand.Action(ctx, c)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(watermarkedFile)
			err = c.Set("input", watermarkedFile)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// c.Set("output", "")
			// err = ffmpeg.SplitByChunksCommand.Action(ctx, c)
			// if err != nil {
			// 	return err
			// }
			title := titleReg.FindString(watermarkedFile)
			if _, isExists := db[filepath.Base(watermarkedFile)]; isExists {
				fmt.Println("skip video", filepath.Base(watermarkedFile))
				return nil
			}
			err = youtube.UploadVideo(ctx, youtube.UploadVideoRequest{
				FilePath:      watermarkedFile,
				Title:         title,
				Description:   "",
				CategoryID:    "",
				PrivacyStatus: "public",
			})
			fmt.Println(err)
			return err
		})

		return nil
	},
}
