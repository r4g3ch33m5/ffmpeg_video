package youtube

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

// downloadVideo downloads a YouTube video using yt-dlp or youtube-dl
func downloadVideo(url, outputDir string) error {
	if outputDir == "" {
		outputDir = util.GetTodayFolder()
	}
	cmd := exec.Command("yt-dlp", "-o", fmt.Sprintf("%s/%%(channel)s.%%(id)s.%%(ext)s", outputDir), url)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}

	return nil
}

// DownloadCommand defines the subcommand for downloading YouTube videos
var DownloadCommand = &cli.Command{
	Name:    "download",
	Aliases: []string{"dl"},
	Usage:   "Download a YouTube video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Usage:    "URL of the YouTube video",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Path to the output directory",
			// Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		url := c.String("url")
		output := c.String("output")

		log.Printf("Downloading video from '%s'...\n", url)
		if err := downloadVideo(url, output); err != nil {
			return fmt.Errorf("error downloading video: %v", err)
		}

		log.Printf("Video successfully downloaded to '%s'.\n", output)
		return nil
	},
}
