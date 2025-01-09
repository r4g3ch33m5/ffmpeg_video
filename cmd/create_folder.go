package cmd

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v3"
)

// createDailyFolder creates a folder named with today's date (video_dd_mm_yy) under the 'source' folder
func createDailyFolder() {
	sourceDir := "source"

	// Ensure the source folder exists
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		if err := os.Mkdir(sourceDir, os.ModePerm); err != nil {
			log.Printf("Failed to create source folder: %v\n", err)
			return
		}
	}

	// Generate the folder name as today's date
	dateFolder := "video_" + time.Now().Format("02_01_06")
	folderPath := filepath.Join(sourceDir, dateFolder)

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
}

// StartCronJob initializes and starts the cron job for creating daily folders
func StartCronJob() {
	cronScheduler := cron.New()

	// Schedule the job to run every day at 2:00 AM
	_, err := cronScheduler.AddFunc("0 2 * * *", createDailyFolder)
	if err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}
	// trigger when start pod
	createDailyFolder()
	log.Println("Cron job scheduled. Running...")
	cronScheduler.Start()

	// Keep the app running to allow cron to execute tasks
	select {}
}

// CreateCommand defines the CLI command for starting the cron scheduler
var CreateCommand = &cli.Command{
	Name:  "start",
	Usage: "Start the cron job for creating daily folders",
	Action: func(ctx context.Context, c *cli.Command) error {
		StartCronJob()
		return nil
	},
}
