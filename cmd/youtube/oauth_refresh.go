package youtube

import (
	"context"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type RefreshTokenRequest struct {
	ClientId     string
	ClientSecret string
	RefreshToken string
}

func refreshToken(ctx context.Context) error {
	cfg := getGgCredentail()

	tokenFile := filepath.Join(".", "credential", "token.json")
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		return err
	}
	newToken, err := cfg.TokenSource(ctx, token).Token()
	if err != nil {
		return err
	}
	saveToken(tokenFile, newToken)
	return nil
}

var RefreshOauth = &cli.Command{
	Name: "refresh",
	Action: func(ctx context.Context, c *cli.Command) error {
		err := refreshToken(ctx)
		return err
	},
}
