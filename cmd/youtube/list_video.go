package youtube

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd/storage"
	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

var defaultSources []string

func init() {
	file, err := os.ReadFile("sources.txt")
	if err != nil {
		panic(err)
	}
	defaultSources = strings.Split(string(file), "\n")
	fmt.Println("default sources", defaultSources)
	fmt.Println("init sources", string(file))
}

var ListVideoCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{},
	Usage:   "",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name: "source",
		}},
	Action: func(ctx context.Context, c *cli.Command) error {
		sources := c.StringSlice("source")
		if len(sources) == 0 {
			sources = defaultSources
		}
		videoSources := make([]string, 0, 10)
		for _, source := range sources {
			args := []string{
				"-I1",
				"--flat-playlist",
				"--print",
				"webpage_url",
				source,
			}

			cmd := exec.Command("yt-dlp", args...)
			value := make([]byte, 0, 2048)
			w := bytes.NewBuffer(value)
			cmd.Stdout = w
			cmd.Stderr = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("err list video", err)
				return err
			}
			fmt.Println(w.String())
			videoSources = append(videoSources, w.String())
		}
		fmt.Println(storage.CreateLocalFolderCommand.Action(ctx, c))
		dateFolder := util.GetTodayFolder()
		file, err := os.OpenFile(filepath.Join("source", dateFolder, "source.txt"), os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(file.Write([]byte(strings.Join(videoSources, "\n"))))
		return nil
	},
}
