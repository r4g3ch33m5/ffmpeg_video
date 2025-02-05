package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ffmpeg_video "github.com/r4g3ch33m5/ffmpeg_video/api/service"
)

type ffmpegServiceImpl struct {
}

func makeDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}
	return nil
}

// SplitVideoIntoChunks splits the input video into chunks of the specified size using FFmpeg
func SplitVideoIntoChunks(ctx context.Context, req *ffmpeg_video.SplitVideoRequest) error {
	var input, outputDir string = req.InputFile, req.OutputDir
	var chunkSize int = int(req.ChunkSize)
	// Ensure the output directory exists
	makeDirIfNotExists(filepath.Join(outputDir))
	outputDir = filepath.Join(outputDir, filepath.Base(input))
	makeDirIfNotExists(outputDir)
	// Command for splitting video into chunks using FFmpeg
	outputPattern := filepath.Join(outputDir, "chunk_%03d.mp4")
	cmd := exec.Command(
		"ffmpeg",
		"-i", input,
		"-c", "copy", // Use copy codec to avoid re-encoding
		"-map", "0:v", // Map all streams
		"-segment_time", strconv.Itoa(chunkSize),
		"-f", "segment",
		"-reset_timestamps", "1",
		outputPattern,
	)

	// Capture and display command output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	log.Printf("Executing command: %s\n", strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ffmpeg: %w", err)
	}

	return nil
}
