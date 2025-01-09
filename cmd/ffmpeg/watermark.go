package ffmpeg

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v3"
)

// addWatermark adds a moving watermark to a video
func addWatermark(inputFile, watermarkFile, outputFile string) error {
	// Define the FFmpeg command
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-i", watermarkFile,
		"-filter_complex",
		"[1:v]format=rgba,colorchannelmixer=aa=0.3[watermark];[0:v][watermark]overlay=x=0:y=0",
		"-c:v", "libx264",
		"-crf", "23",
		"-preset", "veryfast",
		"-c:a", "copy",
		outputFile,
	)

	log.Printf("Executing command: ffmpeg -i %s -i %s -filter_complex [1:v]format=rgba,colorchannelmixer=aa=0.3[watermark];[0:v][watermark]overlay=x=W/2*sin(2*PI*t/10):y=H/2*sin(2*PI*t/15) -c:v libx264 -crf 23 -preset veryfast -c:a copy %s\n",
		inputFile, watermarkFile, outputFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error adding watermark: %w", err)
	}

	log.Printf("Watermarked video created: %s\n", outputFile)
	return nil
}

// AddWatermarkCommand defines the command to add a moving watermark to a video
var AddWatermarkCommand = &cli.Command{
	Name:    "add-watermark",
	Aliases: []string{"watermark"},
	Usage:   "Add a moving watermark with opacity 0.3 to a video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Aliases:  []string{"i"},
			Usage:    "Path to the input video file",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "watermark",
			Aliases: []string{"w"},
			Usage:   "Path to the watermark image file (e.g., PNG)",
			// Required: true,
			Value: "./source/watermarks/github.png",
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Path to the output video file",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		inputFile := c.String("input")
		watermarkFile := c.String("watermark")
		outputFile := c.String("output")

		log.Printf("Adding watermark to video: %s\n", inputFile)
		if err := addWatermark(inputFile, watermarkFile, outputFile); err != nil {
			return fmt.Errorf("error adding watermark: %v", err)
		}

		log.Println("Watermark added successfully.")
		return nil
	},
}
