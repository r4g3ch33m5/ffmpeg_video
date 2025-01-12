package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"

	"github.com/urfave/cli/v3"
)

func getToken() (*oauth2.Token, error) {
	tokenFile := filepath.Join(".", "credential", "token.json")
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open token file: %v", err)
	}
	defer file.Close()

	var token oauth2.Token
	if err := json.NewDecoder(file).Decode(&token); err != nil {
		return nil, fmt.Errorf("unable to parse token file: %v", err)
	}
	return &token, nil

}

// UploadVideo uploads a video to YouTube
func UploadVideo(filePath, title, description, categoryId, privacyStatus string) error {
	ctx := context.Background()
	credBytes, _ := os.ReadFile(filepath.Join(".", "credential", "google_client.json"))

	cred, err := google.ConfigFromJSON(credBytes, youtube.YoutubeUploadScope)
	if err != nil {
		fmt.Println("config", err)
		return err
	}
	token, err := getToken()
	if err != nil {
		fmt.Println("token:", err)
		return err
	}
	service, err := youtube.NewService(ctx, option.WithHTTPClient(cred.Client(ctx, token)))
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
			Value:   "private",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
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
