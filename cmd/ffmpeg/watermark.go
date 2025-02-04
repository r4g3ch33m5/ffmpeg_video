package ffmpeg

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

// GetVideoResolution retrieves the width and height of a video using ffprobe
func GetVideoResolution(videoPath string) (width, height int, err error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", videoPath)
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, fmt.Errorf("ffprobe command failed: %v", err)
	}

	resolution := strings.Split(strings.TrimSpace(string(output)), "x")
	if len(resolution) != 2 {
		return 0, 0, fmt.Errorf("could not parse video resolution")
	}

	width, err = strconv.Atoi(resolution[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid width: %v", err)
	}

	height, err = strconv.Atoi(resolution[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid height: %v", err)
	}

	return width, height, nil
}

// addWatermark adds a moving watermark to a video
func addWatermark(inputFile, watermarkFile, outputFile string) error {
	// Define the FFmpeg command
	// ffmpegFilter := fmt.Sprintf(
	// 	"[1:v]scale=100:180,format=rgba,fade=in:st=%v:d=1,fade=out:st=%v:d=1,geq=r='r(X,Y)':a='255*%f'[watermark];[0:v][watermark]overlay=%d:%d",
	// 	0, 10, 0.3, rand.IntN(480), rand.IntN(1080))
	// Calculate watermark size (1/20th of video resolution)
	folderPath := filepath.Dir(outputFile)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
			log.Printf("Failed to create date folder: %v\n", err)
		}
		log.Printf("Folder '%s' created successfully!\n", folderPath)
	} else {
		log.Printf("Folder '%s' already exists.\n", folderPath)
	}
	width, height, err := GetVideoResolution(inputFile)
	if err != nil {
		return err
	}
	watermarkWidth := width / 20
	watermarkHeight := height / 20

	// Generate random appearance time (between 0 and duration - 5s)
	// startTime := rand.Float64() * (30 - 5)
	// endTime := startTime + 5 // Show for 5 seconds

	// Generate random position (x, y)
	randomX := strconv.Itoa(rand.IntN(width - watermarkWidth))   // Random x within the video width
	randomY := strconv.Itoa(rand.IntN(height - watermarkHeight)) // Random y within the video height

	// FFmpeg filter complex for random position and timing
	ffmpegFilter := "[1:v]format=rgba,colorchannelmixer=aa=0.2[watermark];[0:v][watermark]overlay=x=" + randomX + ":y=" + randomY
	opts := []string{
		"-y",
		"-i", inputFile,
		"-i", watermarkFile,
		"-filter_complex",
		// "[1:v]scale=100:180,format=rgba,colorchannelmixer=aa=0.3[watermark];[0:v][watermark]overlay=" + overlays[rand.IntN(len(overlays))],
		ffmpegFilter,
		// "-c:v", "libx264",
		"-crf", "23",
		"-preset", "veryfast",
		"-c:a", "copy",
		outputFile,
	}
	cmd := exec.Command("ffmpeg", opts...)

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
		if watermarkFile == "" {
			watermarkFile = "./source/watermarks/github.png"
		}
		if outputFile == "" {
			fileName := filepath.Base(inputFile)
			outputFile = filepath.Join(".", "watermarked", util.GetTodayFolder(), fileName)
		}
		log.Printf("Adding watermark to video: %s\n", inputFile)
		if err := addWatermark(inputFile, watermarkFile, outputFile); err != nil {
			return fmt.Errorf("error adding watermark: %v", err)
		}

		log.Println("Watermark added successfully.")
		return nil
	},
}
