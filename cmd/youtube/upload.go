package youtube

import (
	"context"
	"fmt"
	"log"
	"os"

	youtube "google.golang.org/api/youtube/v3"

	"github.com/urfave/cli/v2"
)

// UploadVideo uploads a video to YouTube
func UploadVideo(filePath, title, description, categoryId, privacyStatus string) error {
	ctx := context.Background()

	// Load the client secret from a file
	service, err := youtube.NewService(ctx)
	if err != nil {
		return fmt.Errorf("unable to create YouTube client: %w", err)
	}

	// Open the video file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Define the video details
	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  categoryId,
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: privacyStatus, // "public", "private", or "unlisted"
		},
	}

	// Upload the video
	call := service.Videos.Insert([]string{"snippet", "status"}, video)
	response, err := call.Media(file).Do()
	if err != nil {
		return fmt.Errorf("error uploading video: %w", err)
	}

	log.Printf("Video uploaded successfully: https://www.youtube.com/watch?v=%s\n", response.Id)
	return nil
}

// UploadCommand defines the subcommand for uploading a YouTube video
var UploadCommand = &cli.Command{
	Name:    "upload",
	Aliases: []string{"up"},
	Usage:   "Upload a video to YouTube",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "Path to the video file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "title",
			Aliases:  []string{"t"},
			Usage:    "Title of the video",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Usage:    "Description of the video",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "category_id",
			Aliases: []string{"c"},
			Usage:   "Category ID for the video (default: 22 for 'People & Blogs')",
			Value:   "22",
		},
		&cli.StringFlag{
			Name:    "privacy_status",
			Aliases: []string{"p"},
			Usage:   "Privacy status of the video (public, private, unlisted)",
			Value:   "public",
		},
	},
	Action: func(c *cli.Context) error {
		filePath := c.String("file")
		title := c.String("title")
		description := c.String("description")
		categoryId := c.String("category_id")
		privacyStatus := c.String("privacy_status")

		log.Printf("Uploading video: %s\n", filePath)
		if err := UploadVideo(filePath, title, description, categoryId, privacyStatus); err != nil {
			return fmt.Errorf("error uploading video: %v", err)
		}

		log.Println("Video uploaded successfully.")
		return nil
	},
}
