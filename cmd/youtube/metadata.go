package youtube

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v3"
)

// fetchMetadata fetches metadata of a YouTube video
func fetchMetadata(url string) error {
	cmd := exec.Command("yt-dlp", "--print-json", url)
	cmd.Stderr = nil
	cmd.Stdout = nil

	log.Printf("Fetching metadata for video: %s\n", url)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to fetch metadata: %w", err)
	}

	return nil
}

// MetadataCommand defines the subcommand for fetching metadata
var MetadataCommand = &cli.Command{
	Name:    "metadata",
	Aliases: []string{"meta"},
	Usage:   "Fetch metadata of a YouTube video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Usage:    "URL of the YouTube video",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		url := c.String("url")

		log.Printf("Fetching metadata for '%s'...\n", url)
		if err := fetchMetadata(url); err != nil {
			return fmt.Errorf("error fetching metadata: %v", err)
		}

		log.Printf("Metadata successfully fetched.\n")
		return nil
	},
}
