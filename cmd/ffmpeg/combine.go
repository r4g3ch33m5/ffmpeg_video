package ffmpeg

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/urfave/cli/v3"
)

// combineVideos combines two videos horizontally or vertically
func combineVideos(video1, video2, output, mode string) error {
	var filter string
	if mode == "horizontal" {
		filter = "hstack=inputs=2"
	} else if mode == "vertical" {
		filter = "vstack=inputs=2"
	} else {
		return fmt.Errorf("invalid mode: %s. Use 'horizontal' or 'vertical'", mode)
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i", video1,
		"-i", video2,
		"-filter_complex", filter,
		"-c:v", "libx264",
		"-crf", "23",
		"-preset", "veryfast",
		output,
	)

	log.Printf("Executing command: ffmpeg -i %s -i %s -filter_complex %s -c:v libx264 -crf 23 -preset veryfast %s\n",
		video1, video2, filter, output)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error combining videos: %w", err)
	}

	log.Printf("Combined video created: %s\n", output)
	return nil
}

// CombineVideosCommand defines the command to combine videos from two folders
var CombineVideosCommand = &cli.Command{
	Name:    "combine-videos",
	Aliases: []string{"combine"},
	Usage:   "Combine videos from two folders horizontally or vertically",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "folder1",
			Aliases:  []string{"f1"},
			Usage:    "Path to the first folder of videos",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "folder2",
			Aliases:  []string{"f2"},
			Usage:    "Path to the second folder of videos",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Path to the output folder",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "mode",
			Aliases: []string{"m"},
			Usage:   "Mode for combining videos ('horizontal' or 'vertical')",
			// Required: true,

		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		folder1 := c.String("folder1")
		folder2 := c.String("folder2")
		outputDir := c.String("output")
		mode := c.String("mode")

		if mode != "horizontal" && mode != "vertical" {
			return fmt.Errorf("invalid mode: %s. Use 'horizontal' or 'vertical'", mode)
		}

		// Ensure the output directory exists
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}

		// Get sorted video files from both folders
		files1, err := getSortedFiles(folder1)
		if err != nil {
			return fmt.Errorf("error reading files from folder1: %w", err)
		}
		files2, err := getSortedFiles(folder2)
		if err != nil {
			return fmt.Errorf("error reading files from folder2: %w", err)
		}

		if len(files1) != len(files2) {
			return fmt.Errorf("mismatched number of videos in folder1 (%d) and folder2 (%d)", len(files1), len(files2))
		}

		// Combine each pair of videos
		for i := range files1 {
			video1 := filepath.Join(folder1, files1[i])
			video2 := filepath.Join(folder2, files2[i])
			outputFile := filepath.Join(outputDir, fmt.Sprintf("combined_%d.mp4", i+1))

			log.Printf("Combining: %s and %s\n", video1, video2)
			if err := combineVideos(video1, video2, outputFile, mode); err != nil {
				return fmt.Errorf("error combining videos %s and %s: %v", video1, video2, err)
			}
		}

		log.Println("All videos combined successfully.")
		return nil
	},
}

// getSortedFiles retrieves a sorted list of files from a directory
func getSortedFiles(folder string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", folder, err)
	}

	var fileList []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp4") {
			fileList = append(fileList, file.Name())
		}
	}

	sort.Strings(fileList)
	return fileList, nil
}
