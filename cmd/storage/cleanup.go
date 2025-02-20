package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

var actions = []string{
	"source",
	"splitted",
	"watermarked",
}

var LocalStorageCleanupCommand = &cli.Command{
	Name: "cleanup",
	Action: func(ctx context.Context, c *cli.Command) error {
		todayFolder := util.GetTodayFolder()
		for _, action := range actions {
			cleanFolder := filepath.Join(action, todayFolder)
			fmt.Println("cleanup", cleanFolder)
			err := os.RemoveAll(cleanFolder)
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	},
}
