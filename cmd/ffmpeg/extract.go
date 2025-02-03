package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

// ExtractAudio extracts the audio from a video file using ffmpeg
func ExtractAudio(videoPath, outputFormat, outputDir string) error {
	// Ensure output directory exists
	if outputDir == "" {
		outputDir = filepath.Join(filepath.Dir(videoPath), "subtile")
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	// Get the filename without extension
	baseName := filepath.Base(videoPath)
	fileName := baseName[:len(baseName)-len(filepath.Ext(baseName))]

	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.%s", fileName, outputFormat))
	opts := []string{
		"-y",
		"-i", videoPath,
		"-map", "0:a",
		"-acodec", "libmp3lame",
		outputPath,
	}
	// Build ffmpeg command
	cmd := exec.Command("ffmpeg", opts...)

	// Set command output to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg command failed: %v", err)
	}

	fmt.Printf("Audio extracted successfully: %s\n", outputPath)
	return nil
}

// ExtractAudioCommand defines the CLI command for audio extraction
var ExtractAudioCommand = &cli.Command{
	Name:    "extract_audio",
	Usage:   "Extracts audio from a video file",
	Aliases: []string{"ea"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "input",
			Aliases: []string{"i"},
			Usage:   "input file or folder",
		},
		&cli.StringFlag{
			Name:     "format",
			Aliases:  []string{"f"},
			Usage:    "Audio format (e.g., mp3, wav, aac)",
			Value:    "mp3",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Output directory for the extracted audio",
			Required: false,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		inputEntry := c.String("input")
		if inputEntry == "" {
			inputEntry = util.GetTodayFolder()
		}
		outputFormat := c.String("format")
		outputDir := c.String("output")

		return util.ApplyAllFileInDir(inputEntry, func(path string) error {
			return ExtractAudio(path, outputFormat, outputDir)
		})
	},
}
