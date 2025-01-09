package youtube

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v3"
)

// extractAudioFromYouTube downloads and extracts audio from a YouTube video
func extractAudioFromYouTube(url, outputDir string) error {
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", fmt.Sprintf("%s/%%(title)s.%%(ext)s", outputDir), url)
	cmd.Stderr = nil
	cmd.Stdout = nil

	log.Printf("Executing command: yt-dlp -x --audio-format mp3 -o '%s/%%(title)s.%%(ext)s' %s\n", outputDir, url)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract audio: %w", err)
	}

	return nil
}

// ExtractAudioCommand defines the subcommand for extracting audio
var ExtractAudioCommand = &cli.Command{
	Name:    "extract-audio",
	Aliases: []string{"ea"},
	Usage:   "Extract audio from a YouTube video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Usage:    "URL of the YouTube video",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Path to the output directory",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		url := c.String("url")
		output := c.String("output")

		log.Printf("Extracting audio from '%s'...\n", url)
		if err := extractAudioFromYouTube(url, output); err != nil {
			return fmt.Errorf("error extracting audio: %v", err)
		}

		log.Printf("Audio successfully extracted to '%s'.\n", output)
		return nil
	},
}
