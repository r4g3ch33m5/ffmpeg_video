package storage

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
	"github.com/urfave/cli/v3"
)

// createDailyFolder creates a folder named with today's date (video_dd_mm_yy) under the 'source' folder
func createDailyFolder() {

	// Generate the folder name as today's date
	dateFolder := util.GetTodayFolder()
	folderPath := filepath.Join(".", "source", dateFolder)

	// Create the date-named folder
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
			log.Printf("Failed to create date folder: %v\n", err)
			return
		}
		log.Printf("Folder '%s' created successfully!\n", folderPath)
	} else {
		log.Printf("Folder '%s' already exists.\n", folderPath)
	}
	filePath := filepath.Join(folderPath, "source.txt")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			log.Print(err)
		}
	}
}

// CreateLocalFolderCommand defines the CLI command for starting the cron scheduler
var CreateLocalFolderCommand = &cli.Command{
	Name:  "create_local",
	Usage: "create daily folders on local storage",
	Action: func(ctx context.Context, c *cli.Command) error {
		createDailyFolder()
		return nil
	},
}
