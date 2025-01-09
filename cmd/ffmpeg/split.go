package ffmpeg

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	api "github.com/r4g3ch33m5/ffmpeg_video/api/service"
	"github.com/r4g3ch33m5/ffmpeg_video/service"
	"github.com/urfave/cli/v3"
)

// SplitByChunksCommand defines the subcommand for converting video formats
var SplitByChunksCommand = &cli.Command{
	Name:  "split",
	Usage: "Split a video file into chunks of a specified size",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "input",
			Aliases: []string{"i"},
			Usage:   "Path to the input video file",
			// Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Path to the output directory for chunks",
			// Required: true,
		},
		&cli.IntFlag{
			Name:    "chunk_size",
			Aliases: []string{"cs"},
			Usage:   "Size of each chunk in seconds (default: 30)",
			Value:   30, // Default chunk size
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		input := c.String("input")
		dateFolder := "video_" + time.Now().Format("02_01_06")
		if input == "" {
			input = filepath.Join(".", "source", dateFolder)
			fmt.Println(input)
		}
		output := c.String("output")
		if output == "" {
			output = filepath.Join(".", "output", dateFolder)
		}
		chunkSize := c.Int("chunk_size")

		err := filepath.WalkDir(input, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// Check if it's a file (not a directory)
			if !d.IsDir() {
				log.Printf("Splitting video '%s' into chunks of %d seconds...\n", input, chunkSize)
				if err := service.SplitVideoIntoChunks(ctx, &api.SplitVideoRequest{
					InputFile:    path,
					OutputDir:    output,
					CutTimeStamp: []*api.VideoTimestamp{},
					ChunkSize:    int32(chunkSize),
				}); err != nil {
					return fmt.Errorf("error splitting video: %v", err)
				}
			}

			return nil
		})

		fmt.Println(err)

		log.Printf("Video successfully split into chunks in '%s'.\n", output)
		return nil
	},
}

// splitVideo splits the video based on a list of timestamps
func splitVideo(inputFile string, timestamps []string, outputDir string) error {
	for i, timestamp := range timestamps {
		parts := strings.Split(timestamp, "-")
		if len(parts) != 2 {
			return fmt.Errorf("invalid timestamp format: %s. Expected <start> - <end>", timestamp)
		}
		start := strings.TrimSpace(parts[0])
		end := strings.TrimSpace(parts[1])

		outputFile := fmt.Sprintf("%s/output_%d.mp4", outputDir, i+1)
		cmd := exec.Command(
			"ffmpeg",
			"-i", inputFile,
			"-ss", start,
			"-to", end,
			"-c", "copy",
			outputFile,
		)

		log.Printf("Executing command: ffmpeg -i %s -ss %s -to %s -c copy %s\n", inputFile, start, end, outputFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error splitting video at %s-%s: %w", start, end, err)
		}

		log.Printf("Created clip: %s\n", outputFile)
	}
	return nil
}

// SplitByTimestampsCommand defines the command to split a video based on timestamps
var SplitByTimestampsCommand = &cli.Command{
	Name:    "split-by-timestamps",
	Aliases: []string{"split-ts"},
	Usage:   "Split a video into clips based on a list of timestamps",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Aliases:  []string{"i"},
			Usage:    "Path to the input video file",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "timestamps",
			Aliases:  []string{"t"},
			Usage:    "List of timestamps in the format <start> - <end>",
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
		inputFile := c.String("input")
		timestamps := c.StringSlice("timestamps")
		outputDir := c.String("output")

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}

		log.Printf("Splitting video: %s\n", inputFile)
		if err := splitVideo(inputFile, timestamps, outputDir); err != nil {
			return fmt.Errorf("error splitting video: %v", err)
		}

		log.Println("Video successfully split into clips.")
		return nil
	},
}
