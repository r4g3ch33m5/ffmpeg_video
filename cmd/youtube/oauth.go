package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// Retrieve OAuth2 token from credentials.json
func getClient(config *oauth2.Config) *http.Client {
	// Token file to store user's token
	const tokenFile = "credential.json"

	// Check if token file exists
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}

// Get token from file
func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

// Get token from the web
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Open the following link in your browser and authorize the application:\n%v\n", authURL)

	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&authCode)

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

// Save token to a file
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to save token to file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

var Oauth2Command = &cli.Command{
	Name:  "oauth2",
	Usage: "Upload a video to YouTube",
	Action: func(ctx context.Context, c *cli.Command) error {
		credentialsFile := filepath.Join(".", "credential", "google_client.json")
		tokenFile := filepath.Join(".", "credential", "token.json")

		// Load credentials.json
		b, err := os.ReadFile(credentialsFile)
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// Parse the credentials file
		config, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		// Get token from the web
		token := getTokenFromWeb(config)

		// Save token to file
		saveToken(tokenFile, token)
		return nil
	},
}
